package trade

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"local.exchange-demo/exchange-core-go/domain/order"
)

type Trade struct {
	ID           uuid.UUID
	Symbol       order.Symbol
	MakerOrderID uuid.UUID
	TakerOrderID uuid.UUID
	MakerUserID  uuid.UUID
	TakerUserID  uuid.UUID
	Price        decimal.Decimal
	Quantity     decimal.Decimal
	MakerFee     decimal.Decimal
	TakerFee     decimal.Decimal
	ExecutedAt   time.Time
}
