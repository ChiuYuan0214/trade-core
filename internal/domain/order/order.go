package order

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Order struct {
	ID              uuid.UUID
	ClientOrderID   string
	UserID          uuid.UUID
	Symbol          Symbol
	Side            Side
	Type            Type
	Price           decimal.Decimal
	Quantity        decimal.Decimal
	FilledQuantity  decimal.Decimal
	Status          Status
	RejectionReason string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type CreateInput struct {
	OrderID       uuid.UUID
	ClientOrderID string
	UserID        uuid.UUID
	Symbol        Symbol
	Side          Side
	Type          Type
	Price         decimal.Decimal
	Quantity      decimal.Decimal
	CreatedAt     time.Time
}

func NewPending(input CreateInput) (Order, error) {
	if input.OrderID == uuid.Nil {
		input.OrderID = uuid.New()
	}
	if input.CreatedAt.IsZero() {
		input.CreatedAt = time.Now().UTC()
	}
	if err := ValidateCreateInput(input); err != nil {
		return Order{}, err
	}

	return Order{
		ID:             input.OrderID,
		ClientOrderID:  input.ClientOrderID,
		UserID:         input.UserID,
		Symbol:         input.Symbol,
		Side:           input.Side,
		Type:           input.Type,
		Price:          input.Price,
		Quantity:       input.Quantity,
		FilledQuantity: decimal.Zero,
		Status:         StatusPendingAccept,
		CreatedAt:      input.CreatedAt,
		UpdatedAt:      input.CreatedAt,
	}, nil
}

func ValidateCreateInput(input CreateInput) error {
	if input.UserID == uuid.Nil {
		return fmt.Errorf("user_id is required")
	}
	if err := input.Symbol.Validate(); err != nil {
		return err
	}
	if err := input.Side.Validate(); err != nil {
		return err
	}
	if err := input.Type.Validate(); err != nil {
		return err
	}
	if !input.Quantity.IsPositive() {
		return fmt.Errorf("quantity must be positive")
	}
	if input.Type == TypeLimit && !input.Price.IsPositive() {
		return fmt.Errorf("limit order price must be positive")
	}
	if input.Type == TypeMarket && input.Price.IsNegative() {
		return fmt.Errorf("market order price cannot be negative")
	}
	return nil
}

func (o Order) RemainingQuantity() decimal.Decimal {
	return o.Quantity.Sub(o.FilledQuantity)
}

func (o Order) IsFilled() bool {
	return !o.RemainingQuantity().IsPositive()
}
