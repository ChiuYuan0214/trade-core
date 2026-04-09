package app

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"local.exchange-demo/exchange-core-go/domain/account"
	"local.exchange-demo/exchange-core-go/domain/ledger"
	"local.exchange-demo/exchange-core-go/domain/order"
	"local.exchange-demo/exchange-core-go/domain/trade"
	"local.exchange-demo/exchange-core-go/events"
)

type BalanceStore interface {
	GetBalance(ctx context.Context, userID uuid.UUID, asset string) (account.Balance, error)
	SaveBalance(ctx context.Context, balance account.Balance) error
	Run() error
	Stop()
}

type LedgerStore interface {
	AppendEntries(ctx context.Context, entries ...ledger.Entry) error
	Run() error
	Stop()
}

type AccountApplication interface {
	Reserve(ctx context.Context, input ReserveFundsInput) error
	Release(ctx context.Context, input ReleaseFundsInput) error
	GetBalance(ctx context.Context, userID uuid.UUID, asset string) (account.Balance, error)
	ApplyTrade(ctx context.Context, execution trade.Trade, maker order.Order, taker order.Order) error
}

type ReserveFundsInput struct {
	UserID        uuid.UUID
	Asset         string
	Amount        decimal.Decimal
	ReferenceID   uuid.UUID
	ReferenceType ledger.ReferenceType
}

type ReleaseFundsInput struct {
	UserID        uuid.UUID
	Asset         string
	Amount        decimal.Decimal
	ReferenceID   uuid.UUID
	ReferenceType ledger.ReferenceType
}

type AccountAppService struct {
	BalanceStore BalanceStore
	LedgerStore  LedgerStore
	PrivateEventPublisher PrivateEventPublisher
}

func (s *AccountAppService) Run() error { return nil }
func (s *AccountAppService) Stop()      {}

func (s *AccountAppService) Reserve(ctx context.Context, input ReserveFundsInput) error {
	if s.BalanceStore == nil || s.LedgerStore == nil {
		return fmt.Errorf("account application dependencies not configured")
	}

	current, err := s.BalanceStore.GetBalance(ctx, input.UserID, input.Asset)
	if err != nil {
		return err
	}

	now := time.Now().UTC()
	updated, err := current.Reserve(input.Amount, now)
	if err != nil {
		return err
	}

	entry := ledger.Entry{
		ID:             uuid.New(),
		UserID:         input.UserID,
		Asset:          input.Asset,
		DeltaAvailable: input.Amount.Neg(),
		DeltaFrozen:    input.Amount,
		ReferenceType:  input.ReferenceType,
		ReferenceID:    input.ReferenceID,
		EventID:        uuid.New(),
		CreatedAt:      now,
	}
	if err := s.LedgerStore.AppendEntries(ctx, entry); err != nil {
		return err
	}
	if err := s.BalanceStore.SaveBalance(ctx, updated); err != nil {
		return err
	}
	if s.PrivateEventPublisher != nil {
		_ = s.PrivateEventPublisher.Publish(ctx, PrivateEvent{
			EventID:       uuid.New(),
			EventType:     events.TypeFundsReserved,
			OccurredAt:    now,
			CorrelationID: input.ReferenceID,
			CausationID:   entry.EventID,
			OrderID:       input.ReferenceID,
			UserID:        input.UserID,
			Version:       1,
			Payload: map[string]any{
				"asset":  input.Asset,
				"amount": input.Amount.String(),
			},
		})
	}
	return nil
}

func (s *AccountAppService) Release(ctx context.Context, input ReleaseFundsInput) error {
	if s.BalanceStore == nil || s.LedgerStore == nil {
		return fmt.Errorf("account application dependencies not configured")
	}

	current, err := s.BalanceStore.GetBalance(ctx, input.UserID, input.Asset)
	if err != nil {
		return err
	}

	now := time.Now().UTC()
	updated, err := current.Release(input.Amount, now)
	if err != nil {
		return err
	}

	entry := ledger.Entry{
		ID:             uuid.New(),
		UserID:         input.UserID,
		Asset:          input.Asset,
		DeltaAvailable: input.Amount,
		DeltaFrozen:    input.Amount.Neg(),
		ReferenceType:  input.ReferenceType,
		ReferenceID:    input.ReferenceID,
		EventID:        uuid.New(),
		CreatedAt:      now,
	}
	if err := s.LedgerStore.AppendEntries(ctx, entry); err != nil {
		return err
	}
	if err := s.BalanceStore.SaveBalance(ctx, updated); err != nil {
		return err
	}
	if s.PrivateEventPublisher != nil {
		_ = s.PrivateEventPublisher.Publish(ctx, PrivateEvent{
			EventID:       uuid.New(),
			EventType:     events.TypeFundsReleased,
			OccurredAt:    now,
			CorrelationID: input.ReferenceID,
			CausationID:   entry.EventID,
			OrderID:       input.ReferenceID,
			UserID:        input.UserID,
			Version:       1,
			Payload: map[string]any{
				"asset":  input.Asset,
				"amount": input.Amount.String(),
			},
		})
	}
	return nil
}

func (s *AccountAppService) GetBalance(ctx context.Context, userID uuid.UUID, asset string) (account.Balance, error) {
	if s.BalanceStore == nil {
		return account.Balance{}, fmt.Errorf("balance store not configured")
	}
	return s.BalanceStore.GetBalance(ctx, userID, asset)
}

func (s *AccountAppService) ApplyTrade(ctx context.Context, execution trade.Trade, maker order.Order, taker order.Order) error {
	if s.BalanceStore == nil || s.LedgerStore == nil {
		return fmt.Errorf("account application dependencies not configured")
	}

	buyerOrder, sellerOrder, err := buyerAndSellerOrders(maker, taker)
	if err != nil {
		return err
	}

	quoteAsset, err := account.QuoteAsset(execution.Symbol)
	if err != nil {
		return err
	}
	baseAsset, err := account.BaseAsset(execution.Symbol)
	if err != nil {
		return err
	}

	now := execution.ExecutedAt
	if now.IsZero() {
		now = time.Now().UTC()
	}

	buyerQuote, err := s.BalanceStore.GetBalance(ctx, buyerOrder.UserID, quoteAsset)
	if err != nil {
		return err
	}
	sellerBase, err := s.BalanceStore.GetBalance(ctx, sellerOrder.UserID, baseAsset)
	if err != nil {
		return err
	}

	buyerReservedPerUnit := buyerOrder.Price
	if buyerOrder.Type == order.TypeMarket {
		buyerReservedPerUnit = execution.Price
	}
	frozenQuoteToConsume := buyerReservedPerUnit.Mul(execution.Quantity)
	actualQuoteCost := execution.Price.Mul(execution.Quantity)
	priceImprovementRefund := frozenQuoteToConsume.Sub(actualQuoteCost)
	if priceImprovementRefund.IsNegative() {
		priceImprovementRefund = decimal.Zero
	}

	updatedBuyerQuote := buyerQuote
	updatedBuyerQuote.FrozenAmount = updatedBuyerQuote.FrozenAmount.Sub(frozenQuoteToConsume)
	updatedBuyerQuote.AvailableAmount = updatedBuyerQuote.AvailableAmount.Add(priceImprovementRefund)
	updatedBuyerQuote.UpdatedAt = now
	if updatedBuyerQuote.FrozenAmount.IsNegative() {
		return fmt.Errorf("buyer frozen quote balance became negative")
	}

	updatedSellerBase := sellerBase
	updatedSellerBase.FrozenAmount = updatedSellerBase.FrozenAmount.Sub(execution.Quantity)
	updatedSellerBase.UpdatedAt = now
	if updatedSellerBase.FrozenAmount.IsNegative() {
		return fmt.Errorf("seller frozen base balance became negative")
	}

	buyerBase, err := s.BalanceStore.GetBalance(ctx, buyerOrder.UserID, baseAsset)
	if err != nil {
		return err
	}
	updatedBuyerBase := buyerBase
	updatedBuyerBase.AvailableAmount = updatedBuyerBase.AvailableAmount.Add(execution.Quantity.Sub(execution.TakerFee))
	updatedBuyerBase.UpdatedAt = now

	sellerQuote, err := s.BalanceStore.GetBalance(ctx, sellerOrder.UserID, quoteAsset)
	if err != nil {
		return err
	}
	updatedSellerQuote := sellerQuote
	updatedSellerQuote.AvailableAmount = updatedSellerQuote.AvailableAmount.Add(actualQuoteCost.Sub(execution.MakerFee))
	updatedSellerQuote.UpdatedAt = now

	entries := []ledger.Entry{
		{
			ID:             uuid.New(),
			UserID:         buyerOrder.UserID,
			Asset:          quoteAsset,
			DeltaAvailable: priceImprovementRefund,
			DeltaFrozen:    frozenQuoteToConsume.Neg(),
			ReferenceType:  ledger.ReferenceTypeTradeSettlement,
			ReferenceID:    execution.ID,
			EventID:        uuid.New(),
			CreatedAt:      now,
		},
		{
			ID:             uuid.New(),
			UserID:         buyerOrder.UserID,
			Asset:          baseAsset,
			DeltaAvailable: execution.Quantity.Sub(execution.TakerFee),
			DeltaFrozen:    decimal.Zero,
			ReferenceType:  ledger.ReferenceTypeTradeSettlement,
			ReferenceID:    execution.ID,
			EventID:        uuid.New(),
			CreatedAt:      now,
		},
		{
			ID:             uuid.New(),
			UserID:         sellerOrder.UserID,
			Asset:          baseAsset,
			DeltaAvailable: decimal.Zero,
			DeltaFrozen:    execution.Quantity.Neg(),
			ReferenceType:  ledger.ReferenceTypeTradeSettlement,
			ReferenceID:    execution.ID,
			EventID:        uuid.New(),
			CreatedAt:      now,
		},
		{
			ID:             uuid.New(),
			UserID:         sellerOrder.UserID,
			Asset:          quoteAsset,
			DeltaAvailable: actualQuoteCost.Sub(execution.MakerFee),
			DeltaFrozen:    decimal.Zero,
			ReferenceType:  ledger.ReferenceTypeTradeSettlement,
			ReferenceID:    execution.ID,
			EventID:        uuid.New(),
			CreatedAt:      now,
		},
	}

	if err := s.LedgerStore.AppendEntries(ctx, entries...); err != nil {
		return err
	}
	if err := s.BalanceStore.SaveBalance(ctx, updatedBuyerQuote); err != nil {
		return err
	}
	if err := s.BalanceStore.SaveBalance(ctx, updatedBuyerBase); err != nil {
		return err
	}
	if err := s.BalanceStore.SaveBalance(ctx, updatedSellerBase); err != nil {
		return err
	}
	if err := s.BalanceStore.SaveBalance(ctx, updatedSellerQuote); err != nil {
		return err
	}
	if s.PrivateEventPublisher != nil {
		tradePayload := map[string]any{
			"tradeId":  execution.ID.String(),
			"price":    execution.Price.String(),
			"quantity": execution.Quantity.String(),
		}
		_ = s.PrivateEventPublisher.Publish(ctx, PrivateEvent{
			EventID:       uuid.New(),
			EventType:     events.TypeTradeExecuted,
			OccurredAt:    now,
			CorrelationID: execution.ID,
			CausationID:   execution.ID,
			Symbol:        string(execution.Symbol),
			OrderID:       execution.TakerOrderID,
			UserID:        buyerOrder.UserID,
			Version:       1,
			Payload:       tradePayload,
		})
		_ = s.PrivateEventPublisher.Publish(ctx, PrivateEvent{
			EventID:       uuid.New(),
			EventType:     events.TypeTradeExecuted,
			OccurredAt:    now,
			CorrelationID: execution.ID,
			CausationID:   execution.ID,
			Symbol:        string(execution.Symbol),
			OrderID:       execution.MakerOrderID,
			UserID:        sellerOrder.UserID,
			Version:       1,
			Payload:       tradePayload,
		})
	}
	return nil
}

func buyerAndSellerOrders(maker order.Order, taker order.Order) (buyer order.Order, seller order.Order, err error) {
	switch {
	case maker.Side == order.SideSell && taker.Side == order.SideBuy:
		return taker, maker, nil
	case maker.Side == order.SideBuy && taker.Side == order.SideSell:
		return maker, taker, nil
	default:
		return order.Order{}, order.Order{}, fmt.Errorf("unable to determine buyer/seller for maker=%s taker=%s", maker.Side, taker.Side)
	}
}
