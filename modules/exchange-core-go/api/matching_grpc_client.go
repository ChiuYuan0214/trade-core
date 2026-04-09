package api

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"local.exchange-demo/exchange-core-go/app"
	"local.exchange-demo/exchange-core-go/config"
	"local.exchange-demo/exchange-core-go/domain/order"
	"local.exchange-demo/exchange-core-go/domain/trade"
	matchingv1 "local.exchange-demo/exchange-core-go/gen/proto/exchange/matching/v1"
	"local.exchange-demo/exchange-core-go/matching/book"
)

type MatchingEngineRouter struct {
	Config config.Provider

	Address string
	Timeout time.Duration

	conn          *grpc.ClientConn
	applyOrderFn  func(context.Context, *matchingv1.ApplyOrderRequest, ...grpc.CallOption) (*matchingv1.ApplyOrderResponse, error)
	cancelOrderFn func(context.Context, *matchingv1.CancelOrderRequest, ...grpc.CallOption) (*matchingv1.CancelOrderResponse, error)
}

func (r *MatchingEngineRouter) Run() error {
	if r.Config != nil {
		settings := r.Config.Snapshot()
		if r.Address == "" {
			r.Address = settings.MatchingEngineGRPCAddress
		}
	}
	if r.Address == "" {
		return fmt.Errorf("matching engine grpc address is required")
	}
	if r.Timeout <= 0 {
		r.Timeout = 3 * time.Second
	}

	ctx, cancel := context.WithTimeout(context.Background(), r.Timeout)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		r.Address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return fmt.Errorf("dial matching engine grpc: %w", err)
	}

	r.conn = conn
	client := matchingv1.NewMatchingEngineServiceClient(conn)
	r.applyOrderFn = client.ApplyOrder
	r.cancelOrderFn = client.CancelOrder
	return nil
}

func (r *MatchingEngineRouter) Stop() {
	if r.conn != nil {
		_ = r.conn.Close()
	}
}

func (r *MatchingEngineRouter) ForSymbol(symbol order.Symbol) (app.ShardMatcher, error) {
	if r.applyOrderFn == nil || r.cancelOrderFn == nil {
		return nil, fmt.Errorf("matching engine client is not initialized")
	}
	return &matchingEngineShardClient{
		symbol:        symbol,
		applyOrderFn:  r.applyOrderFn,
		cancelOrderFn: r.cancelOrderFn,
	}, nil
}

type matchingEngineShardClient struct {
	symbol order.Symbol

	applyOrderFn  func(context.Context, *matchingv1.ApplyOrderRequest, ...grpc.CallOption) (*matchingv1.ApplyOrderResponse, error)
	cancelOrderFn func(context.Context, *matchingv1.CancelOrderRequest, ...grpc.CallOption) (*matchingv1.CancelOrderResponse, error)
}

func (c *matchingEngineShardClient) Apply(incoming order.Order, _ time.Time) (book.ApplyResult, uint64, error) {
	response, err := c.applyOrderFn(context.Background(), &matchingv1.ApplyOrderRequest{
		Order: toMatchingOrderProto(incoming),
	})
	if err != nil {
		return book.ApplyResult{}, 0, err
	}
	incomingOrder, err := fromMatchingOrderProto(response.GetIncomingOrder())
	if err != nil {
		return book.ApplyResult{}, 0, err
	}
	orderUpdates := make([]order.Order, 0, len(response.GetOrderUpdates()))
	for _, view := range response.GetOrderUpdates() {
		mapped, mapErr := fromMatchingOrderProto(view)
		if mapErr != nil {
			return book.ApplyResult{}, 0, mapErr
		}
		orderUpdates = append(orderUpdates, mapped)
	}
	trades := make([]trade.Trade, 0, len(response.GetTrades()))
	for _, view := range response.GetTrades() {
		mapped, mapErr := fromTradeProto(view)
		if mapErr != nil {
			return book.ApplyResult{}, 0, mapErr
		}
		trades = append(trades, mapped)
	}
	return book.ApplyResult{
		IncomingOrder: incomingOrder,
		OrderUpdates:  orderUpdates,
		Trades:        trades,
	}, response.GetSequence(), nil
}

func (c *matchingEngineShardClient) Cancel(orderID uuid.UUID, _ time.Time) (*order.Order, uint64, bool) {
	response, err := c.cancelOrderFn(context.Background(), &matchingv1.CancelOrderRequest{
		Symbol:  string(c.symbol),
		OrderId: orderID.String(),
	})
	if err != nil || !response.GetFound() {
		return nil, 0, false
	}
	current, mapErr := fromMatchingOrderProto(response.GetOrder())
	if mapErr != nil {
		return nil, 0, false
	}
	return &current, response.GetSequence(), true
}
