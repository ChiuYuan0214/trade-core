package api

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"local.exchange-demo/exchange-core-go/domain/account"
	"local.exchange-demo/exchange-core-go/domain/order"
	"local.exchange-demo/exchange-core-go/domain/trade"
	ledgerv1 "local.exchange-demo/exchange-core-go/gen/proto/exchange/ledger/v1"
	matchingv1 "local.exchange-demo/exchange-core-go/gen/proto/exchange/matching/v1"
	orderv1 "local.exchange-demo/exchange-core-go/gen/proto/exchange/order/v1"
)

func toOrderProto(current order.Order) *orderv1.OrderView {
	return &orderv1.OrderView{
		OrderId:         current.ID.String(),
		ClientOrderId:   current.ClientOrderID,
		UserId:          current.UserID.String(),
		Symbol:          string(current.Symbol),
		Side:            string(current.Side),
		Type:            string(current.Type),
		Price:           current.Price.String(),
		Quantity:        current.Quantity.String(),
		FilledQuantity:  current.FilledQuantity.String(),
		Status:          string(current.Status),
		RejectionReason: current.RejectionReason,
	}
}

func toBalanceProto(current account.Balance) *orderv1.BalanceView {
	return &orderv1.BalanceView{
		UserId:    current.UserID.String(),
		Asset:     current.Asset,
		Available: current.AvailableAmount.String(),
		Frozen:    current.FrozenAmount.String(),
	}
}

func fromOrderProto(view *orderv1.OrderView) (order.Order, error) {
	if view == nil {
		return order.Order{}, fmt.Errorf("order response is empty")
	}
	orderID, err := uuid.Parse(view.GetOrderId())
	if err != nil {
		return order.Order{}, err
	}
	userID, err := uuid.Parse(view.GetUserId())
	if err != nil {
		return order.Order{}, err
	}
	price, err := decimal.NewFromString(view.GetPrice())
	if err != nil {
		return order.Order{}, err
	}
	quantity, err := decimal.NewFromString(view.GetQuantity())
	if err != nil {
		return order.Order{}, err
	}
	filled, err := decimal.NewFromString(view.GetFilledQuantity())
	if err != nil {
		return order.Order{}, err
	}
	return order.Order{
		ID:              orderID,
		ClientOrderID:   view.GetClientOrderId(),
		UserID:          userID,
		Symbol:          order.Symbol(view.GetSymbol()),
		Side:            order.Side(view.GetSide()),
		Type:            order.Type(view.GetType()),
		Price:           price,
		Quantity:        quantity,
		FilledQuantity:  filled,
		Status:          order.Status(view.GetStatus()),
		RejectionReason: view.GetRejectionReason(),
	}, nil
}

func fromBalanceProto(view *orderv1.BalanceView) (account.Balance, error) {
	if view == nil {
		return account.Balance{}, fmt.Errorf("balance response is empty")
	}
	userID, err := uuid.Parse(view.GetUserId())
	if err != nil {
		return account.Balance{}, err
	}
	available, err := decimal.NewFromString(view.GetAvailable())
	if err != nil {
		return account.Balance{}, err
	}
	frozen, err := decimal.NewFromString(view.GetFrozen())
	if err != nil {
		return account.Balance{}, err
	}
	return account.Balance{
		UserID:          userID,
		Asset:           view.GetAsset(),
		AvailableAmount: available,
		FrozenAmount:    frozen,
	}, nil
}

func toMatchingOrderProto(current order.Order) *matchingv1.MatchingOrder {
	return &matchingv1.MatchingOrder{
		OrderId:         current.ID.String(),
		ClientOrderId:   current.ClientOrderID,
		UserId:          current.UserID.String(),
		Symbol:          string(current.Symbol),
		Side:            string(current.Side),
		Type:            string(current.Type),
		Price:           current.Price.String(),
		Quantity:        current.Quantity.String(),
		FilledQuantity:  current.FilledQuantity.String(),
		Status:          string(current.Status),
		RejectionReason: current.RejectionReason,
		CreatedAt:       current.CreatedAt.UTC().Format(time.RFC3339Nano),
		UpdatedAt:       current.UpdatedAt.UTC().Format(time.RFC3339Nano),
	}
}

func fromMatchingOrderProto(view *matchingv1.MatchingOrder) (order.Order, error) {
	if view == nil {
		return order.Order{}, fmt.Errorf("matching order is empty")
	}
	orderID, err := uuid.Parse(view.GetOrderId())
	if err != nil {
		return order.Order{}, err
	}
	userID, err := uuid.Parse(view.GetUserId())
	if err != nil {
		return order.Order{}, err
	}
	price, err := decimal.NewFromString(view.GetPrice())
	if err != nil {
		return order.Order{}, err
	}
	quantity, err := decimal.NewFromString(view.GetQuantity())
	if err != nil {
		return order.Order{}, err
	}
	filled, err := decimal.NewFromString(view.GetFilledQuantity())
	if err != nil {
		return order.Order{}, err
	}
	createdAt, err := time.Parse(time.RFC3339Nano, view.GetCreatedAt())
	if err != nil {
		return order.Order{}, err
	}
	updatedAt, err := time.Parse(time.RFC3339Nano, view.GetUpdatedAt())
	if err != nil {
		return order.Order{}, err
	}
	return order.Order{
		ID:              orderID,
		ClientOrderID:   view.GetClientOrderId(),
		UserID:          userID,
		Symbol:          order.Symbol(view.GetSymbol()),
		Side:            order.Side(view.GetSide()),
		Type:            order.Type(view.GetType()),
		Price:           price,
		Quantity:        quantity,
		FilledQuantity:  filled,
		Status:          order.Status(view.GetStatus()),
		RejectionReason: view.GetRejectionReason(),
		CreatedAt:       createdAt,
		UpdatedAt:       updatedAt,
	}, nil
}

func toTradeProto(current trade.Trade) *matchingv1.TradeView {
	return &matchingv1.TradeView{
		TradeId:      current.ID.String(),
		Symbol:       string(current.Symbol),
		MakerOrderId: current.MakerOrderID.String(),
		TakerOrderId: current.TakerOrderID.String(),
		MakerUserId:  current.MakerUserID.String(),
		TakerUserId:  current.TakerUserID.String(),
		Price:        current.Price.String(),
		Quantity:     current.Quantity.String(),
		MakerFee:     current.MakerFee.String(),
		TakerFee:     current.TakerFee.String(),
		ExecutedAt:   current.ExecutedAt.UTC().Format(time.RFC3339Nano),
	}
}

func fromTradeProto(view *matchingv1.TradeView) (trade.Trade, error) {
	if view == nil {
		return trade.Trade{}, fmt.Errorf("trade response is empty")
	}
	tradeID, err := uuid.Parse(view.GetTradeId())
	if err != nil {
		return trade.Trade{}, err
	}
	makerOrderID, err := uuid.Parse(view.GetMakerOrderId())
	if err != nil {
		return trade.Trade{}, err
	}
	takerOrderID, err := uuid.Parse(view.GetTakerOrderId())
	if err != nil {
		return trade.Trade{}, err
	}
	makerUserID, err := uuid.Parse(view.GetMakerUserId())
	if err != nil {
		return trade.Trade{}, err
	}
	takerUserID, err := uuid.Parse(view.GetTakerUserId())
	if err != nil {
		return trade.Trade{}, err
	}
	price, err := decimal.NewFromString(view.GetPrice())
	if err != nil {
		return trade.Trade{}, err
	}
	quantity, err := decimal.NewFromString(view.GetQuantity())
	if err != nil {
		return trade.Trade{}, err
	}
	makerFee, err := decimal.NewFromString(view.GetMakerFee())
	if err != nil {
		return trade.Trade{}, err
	}
	takerFee, err := decimal.NewFromString(view.GetTakerFee())
	if err != nil {
		return trade.Trade{}, err
	}
	executedAt, err := time.Parse(time.RFC3339Nano, view.GetExecutedAt())
	if err != nil {
		return trade.Trade{}, err
	}
	return trade.Trade{
		ID:           tradeID,
		Symbol:       order.Symbol(view.GetSymbol()),
		MakerOrderID: makerOrderID,
		TakerOrderID: takerOrderID,
		MakerUserID:  makerUserID,
		TakerUserID:  takerUserID,
		Price:        price,
		Quantity:     quantity,
		MakerFee:     makerFee,
		TakerFee:     takerFee,
		ExecutedAt:   executedAt,
	}, nil
}

func toLedgerOrderProto(current order.Order) *ledgerv1.LedgerOrder {
	return &ledgerv1.LedgerOrder{
		OrderId:         current.ID.String(),
		ClientOrderId:   current.ClientOrderID,
		UserId:          current.UserID.String(),
		Symbol:          string(current.Symbol),
		Side:            string(current.Side),
		Type:            string(current.Type),
		Price:           current.Price.String(),
		Quantity:        current.Quantity.String(),
		FilledQuantity:  current.FilledQuantity.String(),
		Status:          string(current.Status),
		RejectionReason: current.RejectionReason,
		CreatedAt:       current.CreatedAt.UTC().Format(time.RFC3339Nano),
		UpdatedAt:       current.UpdatedAt.UTC().Format(time.RFC3339Nano),
	}
}

func fromLedgerOrderProto(view *ledgerv1.LedgerOrder) (order.Order, error) {
	if view == nil {
		return order.Order{}, fmt.Errorf("ledger order is empty")
	}
	orderID, err := uuid.Parse(view.GetOrderId())
	if err != nil {
		return order.Order{}, err
	}
	userID, err := uuid.Parse(view.GetUserId())
	if err != nil {
		return order.Order{}, err
	}
	price, err := decimal.NewFromString(view.GetPrice())
	if err != nil {
		return order.Order{}, err
	}
	quantity, err := decimal.NewFromString(view.GetQuantity())
	if err != nil {
		return order.Order{}, err
	}
	filled, err := decimal.NewFromString(view.GetFilledQuantity())
	if err != nil {
		return order.Order{}, err
	}
	createdAt, err := time.Parse(time.RFC3339Nano, view.GetCreatedAt())
	if err != nil {
		return order.Order{}, err
	}
	updatedAt, err := time.Parse(time.RFC3339Nano, view.GetUpdatedAt())
	if err != nil {
		return order.Order{}, err
	}
	return order.Order{
		ID:              orderID,
		ClientOrderID:   view.GetClientOrderId(),
		UserID:          userID,
		Symbol:          order.Symbol(view.GetSymbol()),
		Side:            order.Side(view.GetSide()),
		Type:            order.Type(view.GetType()),
		Price:           price,
		Quantity:        quantity,
		FilledQuantity:  filled,
		Status:          order.Status(view.GetStatus()),
		RejectionReason: view.GetRejectionReason(),
		CreatedAt:       createdAt,
		UpdatedAt:       updatedAt,
	}, nil
}

func toLedgerTradeProto(current trade.Trade) *ledgerv1.TradeView {
	return &ledgerv1.TradeView{
		TradeId:      current.ID.String(),
		Symbol:       string(current.Symbol),
		MakerOrderId: current.MakerOrderID.String(),
		TakerOrderId: current.TakerOrderID.String(),
		MakerUserId:  current.MakerUserID.String(),
		TakerUserId:  current.TakerUserID.String(),
		Price:        current.Price.String(),
		Quantity:     current.Quantity.String(),
		MakerFee:     current.MakerFee.String(),
		TakerFee:     current.TakerFee.String(),
		ExecutedAt:   current.ExecutedAt.UTC().Format(time.RFC3339Nano),
	}
}

func fromLedgerTradeProto(view *ledgerv1.TradeView) (trade.Trade, error) {
	if view == nil {
		return trade.Trade{}, fmt.Errorf("ledger trade is empty")
	}
	tradeID, err := uuid.Parse(view.GetTradeId())
	if err != nil {
		return trade.Trade{}, err
	}
	makerOrderID, err := uuid.Parse(view.GetMakerOrderId())
	if err != nil {
		return trade.Trade{}, err
	}
	takerOrderID, err := uuid.Parse(view.GetTakerOrderId())
	if err != nil {
		return trade.Trade{}, err
	}
	makerUserID, err := uuid.Parse(view.GetMakerUserId())
	if err != nil {
		return trade.Trade{}, err
	}
	takerUserID, err := uuid.Parse(view.GetTakerUserId())
	if err != nil {
		return trade.Trade{}, err
	}
	price, err := decimal.NewFromString(view.GetPrice())
	if err != nil {
		return trade.Trade{}, err
	}
	quantity, err := decimal.NewFromString(view.GetQuantity())
	if err != nil {
		return trade.Trade{}, err
	}
	makerFee, err := decimal.NewFromString(view.GetMakerFee())
	if err != nil {
		return trade.Trade{}, err
	}
	takerFee, err := decimal.NewFromString(view.GetTakerFee())
	if err != nil {
		return trade.Trade{}, err
	}
	executedAt, err := time.Parse(time.RFC3339Nano, view.GetExecutedAt())
	if err != nil {
		return trade.Trade{}, err
	}
	return trade.Trade{
		ID:           tradeID,
		Symbol:       order.Symbol(view.GetSymbol()),
		MakerOrderID: makerOrderID,
		TakerOrderID: takerOrderID,
		MakerUserID:  makerUserID,
		TakerUserID:  takerUserID,
		Price:        price,
		Quantity:     quantity,
		MakerFee:     makerFee,
		TakerFee:     takerFee,
		ExecutedAt:   executedAt,
	}, nil
}

func toLedgerBalanceProto(current account.Balance) *ledgerv1.BalanceView {
	return &ledgerv1.BalanceView{
		UserId:    current.UserID.String(),
		Asset:     current.Asset,
		Available: current.AvailableAmount.String(),
		Frozen:    current.FrozenAmount.String(),
	}
}

func fromLedgerBalanceProto(view *ledgerv1.BalanceView) (account.Balance, error) {
	if view == nil {
		return account.Balance{}, fmt.Errorf("ledger balance is empty")
	}
	userID, err := uuid.Parse(view.GetUserId())
	if err != nil {
		return account.Balance{}, err
	}
	available, err := decimal.NewFromString(view.GetAvailable())
	if err != nil {
		return account.Balance{}, err
	}
	frozen, err := decimal.NewFromString(view.GetFrozen())
	if err != nil {
		return account.Balance{}, err
	}
	return account.Balance{
		UserID:          userID,
		Asset:           view.GetAsset(),
		AvailableAmount: available,
		FrozenAmount:    frozen,
	}, nil
}
