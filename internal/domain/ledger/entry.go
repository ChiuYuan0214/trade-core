package ledger

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type ReferenceType string

const (
	ReferenceTypeOrderReservation ReferenceType = "ORDER_RESERVATION"
	ReferenceTypeOrderRelease     ReferenceType = "ORDER_RELEASE"
	ReferenceTypeTradeSettlement  ReferenceType = "TRADE_SETTLEMENT"
)

type Entry struct {
	ID             uuid.UUID
	UserID         uuid.UUID
	Asset          string
	DeltaAvailable decimal.Decimal
	DeltaFrozen    decimal.Decimal
	ReferenceType  ReferenceType
	ReferenceID    uuid.UUID
	EventID        uuid.UUID
	CreatedAt      time.Time
}
