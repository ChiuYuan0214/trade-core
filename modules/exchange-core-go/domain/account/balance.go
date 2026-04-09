package account

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Balance struct {
	UserID          uuid.UUID
	Asset           string
	AvailableAmount decimal.Decimal
	FrozenAmount    decimal.Decimal
	UpdatedAt       time.Time
}

func (b Balance) CanReserve(amount decimal.Decimal) bool {
	return b.AvailableAmount.GreaterThanOrEqual(amount)
}

func (b Balance) Reserve(amount decimal.Decimal, now time.Time) (Balance, error) {
	if !amount.IsPositive() {
		return Balance{}, fmt.Errorf("reserve amount must be positive")
	}
	if !b.CanReserve(amount) {
		return Balance{}, fmt.Errorf("insufficient available balance")
	}

	b.AvailableAmount = b.AvailableAmount.Sub(amount)
	b.FrozenAmount = b.FrozenAmount.Add(amount)
	b.UpdatedAt = now.UTC()
	return b, nil
}

func (b Balance) Release(amount decimal.Decimal, now time.Time) (Balance, error) {
	if !amount.IsPositive() {
		return Balance{}, fmt.Errorf("release amount must be positive")
	}
	if b.FrozenAmount.LessThan(amount) {
		return Balance{}, fmt.Errorf("insufficient frozen balance")
	}

	b.AvailableAmount = b.AvailableAmount.Add(amount)
	b.FrozenAmount = b.FrozenAmount.Sub(amount)
	b.UpdatedAt = now.UTC()
	return b, nil
}
