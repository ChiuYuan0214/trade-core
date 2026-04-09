package events

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type FundsReserved struct {
	OrderID       uuid.UUID
	UserID        uuid.UUID
	Asset         string
	Amount        decimal.Decimal
	ReservedAt    time.Time
	ReservationID uuid.UUID
}

type FundsReservationFailed struct {
	OrderID    uuid.UUID
	UserID     uuid.UUID
	Asset      string
	Amount     decimal.Decimal
	Reason     string
	FailedAt   time.Time
}

type FundsReleased struct {
	OrderID     uuid.UUID
	UserID      uuid.UUID
	Asset       string
	Amount      decimal.Decimal
	ReleasedAt  time.Time
}

type LedgerPosted struct {
	EntryIDs  []uuid.UUID
	PostedAt  time.Time
}
