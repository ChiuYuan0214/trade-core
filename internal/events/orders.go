package events

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"exchange-demo/internal/domain/order"
)

type PlaceOrderCommand struct {
	OrderID       uuid.UUID
	ClientOrderID string
	UserID        uuid.UUID
	Symbol        order.Symbol
	Side          order.Side
	Type          order.Type
	Price         decimal.Decimal
	Quantity      decimal.Decimal
	SubmittedAt   time.Time
}

type CancelOrderCommand struct {
	OrderID     uuid.UUID
	UserID      uuid.UUID
	Symbol      order.Symbol
	RequestedAt time.Time
}

type OrderAccepted struct {
	OrderID    uuid.UUID
	UserID     uuid.UUID
	Symbol     order.Symbol
	AcceptedAt time.Time
}

type OrderRejected struct {
	OrderID         uuid.UUID
	UserID          uuid.UUID
	Symbol          order.Symbol
	Reason          string
	RejectedAt      time.Time
	ReleaseRequired bool
}

type OrderCanceled struct {
	OrderID             uuid.UUID
	UserID              uuid.UUID
	Symbol              order.Symbol
	CanceledAt          time.Time
	ReleasedBaseAmount  decimal.Decimal
	ReleasedQuoteAmount decimal.Decimal
}
