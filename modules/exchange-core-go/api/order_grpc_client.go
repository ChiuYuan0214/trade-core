package api

import (
	"context"
	"fmt"
	"time"

	"local.exchange-demo/exchange-core-go/app"
	"local.exchange-demo/exchange-core-go/config"
	"local.exchange-demo/exchange-core-go/domain/account"
	"local.exchange-demo/exchange-core-go/domain/order"
	orderv1 "local.exchange-demo/exchange-core-go/gen/proto/exchange/order/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type OrderServiceClient struct {
	Config config.Provider

	Address string
	Timeout time.Duration

	conn          *grpc.ClientConn
	placeOrderFn  func(context.Context, *orderv1.PlaceOrderRequest, ...grpc.CallOption) (*orderv1.PlaceOrderResponse, error)
	cancelOrderFn func(context.Context, *orderv1.CancelOrderRequest, ...grpc.CallOption) (*orderv1.CancelOrderResponse, error)
	getOrderFn    func(context.Context, *orderv1.GetOrderRequest, ...grpc.CallOption) (*orderv1.GetOrderResponse, error)
	getBalanceFn  func(context.Context, *orderv1.GetBalanceRequest, ...grpc.CallOption) (*orderv1.GetBalanceResponse, error)
}

func (c *OrderServiceClient) Run() error {
	if c.Config != nil {
		settings := c.Config.Snapshot()
		if c.Address == "" {
			c.Address = settings.OrderServiceGRPCAddress
		}
	}
	if c.Address == "" {
		return fmt.Errorf("order service grpc address is required")
	}
	if c.Timeout <= 0 {
		c.Timeout = 3 * time.Second
	}

	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		c.Address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return fmt.Errorf("dial order service grpc: %w", err)
	}

	c.conn = conn
	client := orderv1.NewOrderServiceClient(conn)
	c.placeOrderFn = client.PlaceOrder
	c.cancelOrderFn = client.CancelOrder
	c.getOrderFn = client.GetOrder
	c.getBalanceFn = client.GetBalance
	return nil
}

func (c *OrderServiceClient) Stop() {
	if c.conn != nil {
		_ = c.conn.Close()
	}
}

func (c *OrderServiceClient) PlaceOrder(ctx context.Context, input app.PlaceOrderInput) (app.PlaceOrderResult, error) {
	if c.placeOrderFn == nil {
		return app.PlaceOrderResult{}, fmt.Errorf("order service client is not initialized")
	}
	response, err := c.placeOrderFn(ctx, &orderv1.PlaceOrderRequest{
		ClientOrderId: input.ClientOrderID,
		UserId:        input.UserID.String(),
		Symbol:        string(input.Symbol),
		Side:          string(input.Side),
		Type:          string(input.Type),
		Price:         input.Price.String(),
		Quantity:      input.Quantity.String(),
	})
	if err != nil {
		return app.PlaceOrderResult{}, err
	}
	mapped, err := fromOrderProto(response.GetOrder())
	if err != nil {
		return app.PlaceOrderResult{}, err
	}
	return app.PlaceOrderResult{Order: mapped}, nil
}

func (c *OrderServiceClient) CancelOrder(ctx context.Context, orderID string) (order.Order, error) {
	if c.cancelOrderFn == nil {
		return order.Order{}, fmt.Errorf("order service client is not initialized")
	}
	response, err := c.cancelOrderFn(ctx, &orderv1.CancelOrderRequest{OrderId: orderID})
	if err != nil {
		return order.Order{}, err
	}
	return fromOrderProto(response.GetOrder())
}

func (c *OrderServiceClient) GetOrder(ctx context.Context, orderID string) (order.Order, error) {
	if c.getOrderFn == nil {
		return order.Order{}, fmt.Errorf("order service client is not initialized")
	}
	response, err := c.getOrderFn(ctx, &orderv1.GetOrderRequest{OrderId: orderID})
	if err != nil {
		return order.Order{}, err
	}
	return fromOrderProto(response.GetOrder())
}

func (c *OrderServiceClient) GetBalance(ctx context.Context, userID string, asset string) (account.Balance, error) {
	if c.getBalanceFn == nil {
		return account.Balance{}, fmt.Errorf("order service client is not initialized")
	}
	response, err := c.getBalanceFn(ctx, &orderv1.GetBalanceRequest{
		UserId: userID,
		Asset:  asset,
	})
	if err != nil {
		return account.Balance{}, err
	}
	return fromBalanceProto(response.GetBalance())
}
