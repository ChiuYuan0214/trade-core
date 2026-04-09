package api

import (
	"context"

	"local.exchange-demo/exchange-core-go/app"
	"local.exchange-demo/exchange-core-go/domain/account"
	"local.exchange-demo/exchange-core-go/domain/order"
)

type GatewayApplication interface {
	PlaceOrder(ctx context.Context, input app.PlaceOrderInput) (app.PlaceOrderResult, error)
	CancelOrder(ctx context.Context, orderID string) (order.Order, error)
	GetOrder(ctx context.Context, orderID string) (order.Order, error)
	GetBalance(ctx context.Context, userID string, asset string) (account.Balance, error)
	Run() error
	Stop()
}
