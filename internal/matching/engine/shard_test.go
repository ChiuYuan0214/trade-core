package engine

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"exchange-demo/internal/domain/order"
)

func TestShardIncrementsSequenceOnApplyAndCancel(t *testing.T) {
	t.Parallel()

	shard, err := NewShard("shard-1", order.SymbolBTCUSDT)
	if err != nil {
		t.Fatalf("new shard: %v", err)
	}

	now := time.Date(2026, 4, 9, 4, 0, 0, 0, time.UTC)
	resting, err := order.NewPending(order.CreateInput{
		OrderID:   uuid.New(),
		UserID:    uuid.New(),
		Symbol:    order.SymbolBTCUSDT,
		Side:      order.SideBuy,
		Type:      order.TypeLimit,
		Price:     decimal.RequireFromString("59000"),
		Quantity:  decimal.RequireFromString("1"),
		CreatedAt: now,
	})
	if err != nil {
		t.Fatalf("new pending order: %v", err)
	}

	applyResult, sequence, err := shard.Apply(resting, now)
	if err != nil {
		t.Fatalf("apply order: %v", err)
	}
	if applyResult.RestingOrder == nil {
		t.Fatalf("expected resting order after apply")
	}
	if sequence != 1 {
		t.Fatalf("expected sequence 1, got %d", sequence)
	}

	_, sequence, ok := shard.Cancel(resting.ID, now.Add(time.Second))
	if !ok {
		t.Fatalf("expected cancel to succeed")
	}
	if sequence != 2 {
		t.Fatalf("expected sequence 2, got %d", sequence)
	}
}
