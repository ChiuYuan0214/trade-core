package account

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func TestReserveAndRelease(t *testing.T) {
	t.Parallel()

	now := time.Date(2026, 4, 9, 2, 0, 0, 0, time.UTC)
	balance := Balance{
		UserID:          uuid.New(),
		Asset:           "USDT",
		AvailableAmount: decimal.RequireFromString("100"),
		FrozenAmount:    decimal.Zero,
	}

	reserved, err := balance.Reserve(decimal.RequireFromString("30"), now)
	if err != nil {
		t.Fatalf("reserve returned error: %v", err)
	}
	if !reserved.AvailableAmount.Equal(decimal.RequireFromString("70")) {
		t.Fatalf("unexpected available balance: %s", reserved.AvailableAmount)
	}
	if !reserved.FrozenAmount.Equal(decimal.RequireFromString("30")) {
		t.Fatalf("unexpected frozen balance: %s", reserved.FrozenAmount)
	}

	released, err := reserved.Release(decimal.RequireFromString("10"), now.Add(time.Minute))
	if err != nil {
		t.Fatalf("release returned error: %v", err)
	}
	if !released.AvailableAmount.Equal(decimal.RequireFromString("80")) {
		t.Fatalf("unexpected available balance after release: %s", released.AvailableAmount)
	}
	if !released.FrozenAmount.Equal(decimal.RequireFromString("20")) {
		t.Fatalf("unexpected frozen balance after release: %s", released.FrozenAmount)
	}
}
