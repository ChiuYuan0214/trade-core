package book

import (
	"fmt"
	"slices"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"exchange-demo/internal/domain/order"
	"exchange-demo/internal/domain/trade"
)

type Book struct {
	symbol order.Symbol
	buys   []order.Order
	sells  []order.Order
	lookup map[uuid.UUID]bookSide
}

type bookSide string

const (
	sideBuy  bookSide = "BUY"
	sideSell bookSide = "SELL"
)

type ApplyResult struct {
	IncomingOrder order.Order
	RestingOrder *order.Order
	OrderUpdates []order.Order
	Trades       []trade.Trade
}

func New(symbol order.Symbol) (*Book, error) {
	if err := symbol.Validate(); err != nil {
		return nil, err
	}
	return &Book{
		symbol: symbol,
		lookup: make(map[uuid.UUID]bookSide),
	}, nil
}

func (b *Book) Apply(incoming order.Order, now time.Time) (ApplyResult, error) {
	if incoming.Symbol != b.symbol {
		return ApplyResult{}, fmt.Errorf("book symbol %s does not match incoming symbol %s", b.symbol, incoming.Symbol)
	}
	if !incoming.RemainingQuantity().IsPositive() {
		return ApplyResult{}, fmt.Errorf("incoming order has no remaining quantity")
	}

	working := incoming
	working.UpdatedAt = now.UTC()
	working.Status = order.StatusOpen

	trades, orderUpdates, err := b.match(&working, now.UTC())
	if err != nil {
		return ApplyResult{}, err
	}

	if !working.RemainingQuantity().IsPositive() {
		working.Status = order.StatusFilled
		orderUpdates = append(orderUpdates, working)
		return ApplyResult{
			IncomingOrder: working,
			OrderUpdates:  orderUpdates,
			Trades:        trades,
		}, nil
	}

	if working.Type == order.TypeMarket {
		if working.FilledQuantity.IsPositive() {
			working.Status = order.StatusPartiallyFilled
		} else {
			working.Status = order.StatusRejected
			working.RejectionReason = "insufficient liquidity"
		}
		orderUpdates = append(orderUpdates, working)
		return ApplyResult{
			IncomingOrder: working,
			OrderUpdates:  orderUpdates,
			Trades:        trades,
		}, nil
	}

	if working.FilledQuantity.IsPositive() {
		working.Status = order.StatusPartiallyFilled
	}

	b.addResting(working)
	resting := working
	orderUpdates = append(orderUpdates, working)
	return ApplyResult{
		IncomingOrder: working,
		RestingOrder: &resting,
		OrderUpdates: orderUpdates,
		Trades:       trades,
	}, nil
}

func (b *Book) Cancel(orderID uuid.UUID, now time.Time) (*order.Order, bool) {
	side, ok := b.lookup[orderID]
	if !ok {
		return nil, false
	}

	var canceled order.Order
	switch side {
	case sideBuy:
		index := slices.IndexFunc(b.buys, func(candidate order.Order) bool {
			return candidate.ID == orderID
		})
		if index < 0 {
			return nil, false
		}
		canceled = b.buys[index]
		b.buys = append(b.buys[:index], b.buys[index+1:]...)
	case sideSell:
		index := slices.IndexFunc(b.sells, func(candidate order.Order) bool {
			return candidate.ID == orderID
		})
		if index < 0 {
			return nil, false
		}
		canceled = b.sells[index]
		b.sells = append(b.sells[:index], b.sells[index+1:]...)
	default:
		return nil, false
	}

	delete(b.lookup, orderID)
	canceled.Status = order.StatusCanceled
	canceled.UpdatedAt = now.UTC()
	return &canceled, true
}

func (b *Book) Snapshot(depth int) (bids []Level, asks []Level) {
	if depth <= 0 {
		return nil, nil
	}

	return aggregateLevels(b.buys, depth), aggregateLevels(b.sells, depth)
}

type Level struct {
	Price    decimal.Decimal
	Quantity decimal.Decimal
}

func (b *Book) match(incoming *order.Order, now time.Time) ([]trade.Trade, []order.Order, error) {
	var (
		trades       []trade.Trade
		orderUpdates []order.Order
		err          error
	)

	switch incoming.Side {
	case order.SideBuy:
		trades, orderUpdates, err = b.consume(&b.sells, incoming, now, crossesBuy)
	case order.SideSell:
		trades, orderUpdates, err = b.consume(&b.buys, incoming, now, crossesSell)
	default:
		err = fmt.Errorf("unsupported incoming side %s", incoming.Side)
	}

	return trades, orderUpdates, err
}

func (b *Book) consume(bookSideOrders *[]order.Order, taker *order.Order, now time.Time, crosses func(order.Order, order.Order) bool) ([]trade.Trade, []order.Order, error) {
	var (
		trades       []trade.Trade
		orderUpdates []order.Order
	)

	for len(*bookSideOrders) > 0 && taker.RemainingQuantity().IsPositive() {
		maker := &(*bookSideOrders)[0]
		if !crosses(*maker, *taker) {
			break
		}

		fillQty := decimal.Min(maker.RemainingQuantity(), taker.RemainingQuantity())
		if !fillQty.IsPositive() {
			return nil, nil, fmt.Errorf("invalid non-positive fill quantity")
		}

		maker.FilledQuantity = maker.FilledQuantity.Add(fillQty)
		taker.FilledQuantity = taker.FilledQuantity.Add(fillQty)
		maker.UpdatedAt = now
		taker.UpdatedAt = now

		if maker.IsFilled() {
			maker.Status = order.StatusFilled
		} else {
			maker.Status = order.StatusPartiallyFilled
		}
		if taker.IsFilled() {
			taker.Status = order.StatusFilled
		} else {
			taker.Status = order.StatusPartiallyFilled
		}

		trades = append(trades, trade.Trade{
			ID:           uuid.New(),
			Symbol:       b.symbol,
			MakerOrderID: maker.ID,
			TakerOrderID: taker.ID,
			MakerUserID:  maker.UserID,
			TakerUserID:  taker.UserID,
			Price:        maker.Price,
			Quantity:     fillQty,
			MakerFee:     decimal.Zero,
			TakerFee:     decimal.Zero,
			ExecutedAt:   now,
		})
		orderUpdates = append(orderUpdates, *maker)

		if maker.IsFilled() {
			delete(b.lookup, maker.ID)
			*bookSideOrders = (*bookSideOrders)[1:]
		}
	}

	return trades, orderUpdates, nil
}

func (b *Book) addResting(incoming order.Order) {
	switch incoming.Side {
	case order.SideBuy:
		b.buys = append(b.buys, incoming)
		sortBuys(b.buys)
		b.lookup[incoming.ID] = sideBuy
	case order.SideSell:
		b.sells = append(b.sells, incoming)
		sortSells(b.sells)
		b.lookup[incoming.ID] = sideSell
	}
}

func sortBuys(orders []order.Order) {
	slices.SortStableFunc(orders, func(left order.Order, right order.Order) int {
		switch {
		case left.Price.GreaterThan(right.Price):
			return -1
		case left.Price.LessThan(right.Price):
			return 1
		case left.CreatedAt.Before(right.CreatedAt):
			return -1
		case left.CreatedAt.After(right.CreatedAt):
			return 1
		default:
			switch {
			case left.ID.String() < right.ID.String():
				return -1
			case left.ID.String() > right.ID.String():
				return 1
			default:
				return 0
			}
		}
	})
}

func sortSells(orders []order.Order) {
	slices.SortStableFunc(orders, func(left order.Order, right order.Order) int {
		switch {
		case left.Price.LessThan(right.Price):
			return -1
		case left.Price.GreaterThan(right.Price):
			return 1
		case left.CreatedAt.Before(right.CreatedAt):
			return -1
		case left.CreatedAt.After(right.CreatedAt):
			return 1
		default:
			switch {
			case left.ID.String() < right.ID.String():
				return -1
			case left.ID.String() > right.ID.String():
				return 1
			default:
				return 0
			}
		}
	})
}

func crossesBuy(bestAsk order.Order, incoming order.Order) bool {
	if incoming.Type == order.TypeMarket {
		return true
	}
	return incoming.Price.GreaterThanOrEqual(bestAsk.Price)
}

func crossesSell(bestBid order.Order, incoming order.Order) bool {
	if incoming.Type == order.TypeMarket {
		return true
	}
	return incoming.Price.LessThanOrEqual(bestBid.Price)
}

func aggregateLevels(orders []order.Order, depth int) []Level {
	levels := make([]Level, 0, depth)
	for _, resting := range orders {
		remaining := resting.RemainingQuantity()
		if !remaining.IsPositive() {
			continue
		}
		if len(levels) > 0 && levels[len(levels)-1].Price.Equal(resting.Price) {
			levels[len(levels)-1].Quantity = levels[len(levels)-1].Quantity.Add(remaining)
			continue
		}
		levels = append(levels, Level{Price: resting.Price, Quantity: remaining})
		if len(levels) == depth {
			break
		}
	}
	return levels
}
