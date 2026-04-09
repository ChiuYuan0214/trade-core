package events

import (
	"time"

	"github.com/google/uuid"

	"exchange-demo/internal/domain/order"
)

type Type string

const (
	TypePlaceOrderCommand        Type = "PlaceOrderCommand"
	TypeCancelOrderCommand       Type = "CancelOrderCommand"
	TypeFundsReserved            Type = "FundsReserved"
	TypeFundsReservationFailed   Type = "FundsReservationFailed"
	TypeOrderAccepted            Type = "OrderAccepted"
	TypeOrderRejected            Type = "OrderRejected"
	TypeOrderCanceled            Type = "OrderCanceled"
	TypeTradeExecuted            Type = "TradeExecuted"
	TypeOrderBookUpdated         Type = "OrderBookUpdated"
	TypeFundsReleased            Type = "FundsReleased"
	TypeLedgerPosted             Type = "LedgerPosted"
)

type Envelope[T any] struct {
	EventID       uuid.UUID
	EventType     Type
	OccurredAt    time.Time
	CorrelationID uuid.UUID
	CausationID   uuid.UUID
	Symbol        order.Symbol
	OrderID       uuid.UUID
	UserID        uuid.UUID
	ShardID       string
	Version       int
	Payload       T
}
