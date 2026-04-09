package events

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"exchange-demo/internal/domain/order"
)

type TradeExecuted struct {
	TradeID       uuid.UUID
	Symbol        order.Symbol
	MakerOrderID  uuid.UUID
	TakerOrderID  uuid.UUID
	MakerUserID   uuid.UUID
	TakerUserID   uuid.UUID
	Price         decimal.Decimal
	Quantity      decimal.Decimal
	MakerFee      decimal.Decimal
	TakerFee      decimal.Decimal
	ExecutedAt    time.Time
}

type PriceLevel struct {
	Price    decimal.Decimal
	Quantity decimal.Decimal
}

type OrderBookUpdated struct {
	Symbol      order.Symbol
	Sequence    uint64
	GeneratedAt time.Time
	Bids        []PriceLevel
	Asks        []PriceLevel
}
