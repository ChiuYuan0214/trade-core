package book

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"exchange-demo/internal/domain/order"
)

func TestApplyUsesPriceTimePriority(t *testing.T) {
	t.Parallel()

	now := time.Date(2026, 4, 9, 3, 0, 0, 0, time.UTC)
	book, err := New(order.SymbolBTCUSDT)
	if err != nil {
		t.Fatalf("new book: %v", err)
	}

	firstAsk := mustOrder(t, order.CreateInput{
		OrderID:   uuid.New(),
		UserID:    uuid.New(),
		Symbol:    order.SymbolBTCUSDT,
		Side:      order.SideSell,
		Type:      order.TypeLimit,
		Price:     decimal.RequireFromString("60000"),
		Quantity:  decimal.RequireFromString("1"),
		CreatedAt: now,
	})
	secondAsk := mustOrder(t, order.CreateInput{
		OrderID:   uuid.New(),
		UserID:    uuid.New(),
		Symbol:    order.SymbolBTCUSDT,
		Side:      order.SideSell,
		Type:      order.TypeLimit,
		Price:     decimal.RequireFromString("60000"),
		Quantity:  decimal.RequireFromString("1"),
		CreatedAt: now.Add(time.Second),
	})

	if _, err := book.Apply(firstAsk, now); err != nil {
		t.Fatalf("apply first ask: %v", err)
	}
	if _, err := book.Apply(secondAsk, now.Add(time.Second)); err != nil {
		t.Fatalf("apply second ask: %v", err)
	}

	taker := mustOrder(t, order.CreateInput{
		OrderID:   uuid.New(),
		UserID:    uuid.New(),
		Symbol:    order.SymbolBTCUSDT,
		Side:      order.SideBuy,
		Type:      order.TypeLimit,
		Price:     decimal.RequireFromString("60000"),
		Quantity:  decimal.RequireFromString("1.5"),
		CreatedAt: now.Add(2 * time.Second),
	})

	result, err := book.Apply(taker, now.Add(2*time.Second))
	if err != nil {
		t.Fatalf("apply taker: %v", err)
	}
	if len(result.Trades) != 2 {
		t.Fatalf("expected 2 trades, got %d", len(result.Trades))
	}
	if result.Trades[0].MakerOrderID != firstAsk.ID {
		t.Fatalf("expected first maker to be oldest order")
	}
	if !result.Trades[0].Quantity.Equal(decimal.RequireFromString("1")) {
		t.Fatalf("unexpected first fill quantity: %s", result.Trades[0].Quantity)
	}
	if !result.Trades[1].Quantity.Equal(decimal.RequireFromString("0.5")) {
		t.Fatalf("unexpected second fill quantity: %s", result.Trades[1].Quantity)
	}
}

func TestCancelRemovesRestingOrder(t *testing.T) {
	t.Parallel()

	now := time.Date(2026, 4, 9, 3, 30, 0, 0, time.UTC)
	book, err := New(order.SymbolBTCUSDT)
	if err != nil {
		t.Fatalf("new book: %v", err)
	}

	resting := mustOrder(t, order.CreateInput{
		OrderID:   uuid.New(),
		UserID:    uuid.New(),
		Symbol:    order.SymbolBTCUSDT,
		Side:      order.SideBuy,
		Type:      order.TypeLimit,
		Price:     decimal.RequireFromString("59000"),
		Quantity:  decimal.RequireFromString("1"),
		CreatedAt: now,
	})

	result, err := book.Apply(resting, now)
	if err != nil {
		t.Fatalf("apply resting order: %v", err)
	}
	if result.RestingOrder == nil {
		t.Fatalf("expected resting order to remain in book")
	}

	canceled, ok := book.Cancel(resting.ID, now.Add(time.Second))
	if !ok {
		t.Fatalf("expected cancel to succeed")
	}
	if canceled.Status != order.StatusCanceled {
		t.Fatalf("expected canceled status, got %s", canceled.Status)
	}

	bids, _ := book.Snapshot(5)
	if len(bids) != 0 {
		t.Fatalf("expected empty bid side after cancel")
	}
}

func mustOrder(t *testing.T, input order.CreateInput) order.Order {
	t.Helper()

	created, err := order.NewPending(input)
	if err != nil {
		t.Fatalf("new pending order: %v", err)
	}
	return created
}
