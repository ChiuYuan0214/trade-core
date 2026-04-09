package app

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"exchange-demo/internal/domain/account"
	"exchange-demo/internal/domain/ledger"
	"exchange-demo/internal/domain/order"
	"exchange-demo/internal/domain/trade"
	"exchange-demo/internal/matching/book"
)

var ErrOrderNotFound = errors.New("order not found")

type OrderStore interface {
	Save(ctx context.Context, current order.Order) error
	Get(ctx context.Context, orderID uuid.UUID) (order.Order, error)
	Run() error
	Stop()
}

type ShardMatcher interface {
	Apply(incoming order.Order, now time.Time) (book.ApplyResult, uint64, error)
	Cancel(orderID uuid.UUID, now time.Time) (*order.Order, uint64, bool)
}

type ShardRouter interface {
	ForSymbol(symbol order.Symbol) (ShardMatcher, error)
	Run() error
	Stop()
}

type OrderApplication interface {
	PlaceOrder(ctx context.Context, input PlaceOrderInput) (PlaceOrderResult, error)
	CancelOrder(ctx context.Context, orderID uuid.UUID) (order.Order, error)
	GetOrder(ctx context.Context, orderID uuid.UUID) (order.Order, error)
}

type PlaceOrderInput struct {
	ClientOrderID string
	UserID        uuid.UUID
	Symbol        order.Symbol
	Side          order.Side
	Type          order.Type
	Price         decimal.Decimal
	Quantity      decimal.Decimal
}

type PlaceOrderResult struct {
	Order  order.Order
	Trades []trade.Trade
}

type OrderAppService struct {
	OrderStore         OrderStore
	ShardRouter        ShardRouter
	AccountApplication AccountApplication
}

func (s *OrderAppService) Run() error { return nil }
func (s *OrderAppService) Stop()      {}

func (s *OrderAppService) PlaceOrder(ctx context.Context, input PlaceOrderInput) (PlaceOrderResult, error) {
	if s.OrderStore == nil || s.ShardRouter == nil || s.AccountApplication == nil {
		return PlaceOrderResult{}, fmt.Errorf("order application dependencies not configured")
	}

	pending, err := order.NewPending(order.CreateInput{
		ClientOrderID: input.ClientOrderID,
		UserID:        input.UserID,
		Symbol:        input.Symbol,
		Side:          input.Side,
		Type:          input.Type,
		Price:         input.Price,
		Quantity:      input.Quantity,
	})
	if err != nil {
		return PlaceOrderResult{}, err
	}

	reserveAsset, err := account.ReservationAsset(pending)
	if err != nil {
		return PlaceOrderResult{}, err
	}
	reserveAmount, err := account.ReservationAmount(pending)
	if err != nil {
		return PlaceOrderResult{}, err
	}
	if err := s.AccountApplication.Reserve(ctx, ReserveFundsInput{
		UserID:        pending.UserID,
		Asset:         reserveAsset,
		Amount:        reserveAmount,
		ReferenceID:   pending.ID,
		ReferenceType: ledger.ReferenceTypeOrderReservation,
	}); err != nil {
		return PlaceOrderResult{}, err
	}

	if err := s.OrderStore.Save(ctx, pending); err != nil {
		return PlaceOrderResult{}, err
	}

	shard, err := s.ShardRouter.ForSymbol(input.Symbol)
	if err != nil {
		return PlaceOrderResult{}, err
	}

	result, _, err := shard.Apply(pending, time.Now().UTC())
	if err != nil {
		return PlaceOrderResult{}, err
	}

	orderByID := map[uuid.UUID]order.Order{
		result.IncomingOrder.ID: result.IncomingOrder,
	}
	for _, updated := range result.OrderUpdates {
		orderByID[updated.ID] = updated
		if err := s.OrderStore.Save(ctx, updated); err != nil {
			return PlaceOrderResult{}, err
		}
	}
	for _, executedTrade := range result.Trades {
		makerOrder, ok := orderByID[executedTrade.MakerOrderID]
		if !ok {
			return PlaceOrderResult{}, fmt.Errorf("maker order %s missing from apply result", executedTrade.MakerOrderID)
		}
		takerOrder, ok := orderByID[executedTrade.TakerOrderID]
		if !ok {
			return PlaceOrderResult{}, fmt.Errorf("taker order %s missing from apply result", executedTrade.TakerOrderID)
		}
		if err := s.AccountApplication.ApplyTrade(ctx, executedTrade, makerOrder, takerOrder); err != nil {
			return PlaceOrderResult{}, err
		}
	}

	return PlaceOrderResult{
		Order:  result.IncomingOrder,
		Trades: result.Trades,
	}, nil
}

func (s *OrderAppService) CancelOrder(ctx context.Context, orderID uuid.UUID) (order.Order, error) {
	if s.OrderStore == nil || s.ShardRouter == nil || s.AccountApplication == nil {
		return order.Order{}, fmt.Errorf("order application dependencies not configured")
	}

	current, err := s.OrderStore.Get(ctx, orderID)
	if err != nil {
		return order.Order{}, err
	}

	shard, err := s.ShardRouter.ForSymbol(current.Symbol)
	if err != nil {
		return order.Order{}, err
	}

	canceled, _, ok := shard.Cancel(orderID, time.Now().UTC())
	if !ok {
		return order.Order{}, ErrOrderNotFound
	}

	releaseAsset, err := account.ReservationAsset(current)
	if err != nil {
		return order.Order{}, err
	}
	remainingAmount, err := account.ReservationAmount(order.Order{
		Symbol:   current.Symbol,
		Side:     current.Side,
		Type:     current.Type,
		Price:    current.Price,
		Quantity: current.RemainingQuantity(),
	})
	if err != nil {
		return order.Order{}, err
	}
	if remainingAmount.IsPositive() {
		if err := s.AccountApplication.Release(ctx, ReleaseFundsInput{
			UserID:        current.UserID,
			Asset:         releaseAsset,
			Amount:        remainingAmount,
			ReferenceID:   current.ID,
			ReferenceType: ledger.ReferenceTypeOrderRelease,
		}); err != nil {
			return order.Order{}, err
		}
	}

	if err := s.OrderStore.Save(ctx, *canceled); err != nil {
		return order.Order{}, err
	}
	return *canceled, nil
}

func (s *OrderAppService) GetOrder(ctx context.Context, orderID uuid.UUID) (order.Order, error) {
	if s.OrderStore == nil {
		return order.Order{}, fmt.Errorf("order store not configured")
	}
	return s.OrderStore.Get(ctx, orderID)
}
