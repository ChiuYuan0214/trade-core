package api

import (
	"context"
	"fmt"
	"net"

	"local.exchange-demo/exchange-core-go/app"
	"local.exchange-demo/exchange-core-go/config"
	"local.exchange-demo/exchange-core-go/domain/order"
	orderv1 "local.exchange-demo/exchange-core-go/gen/proto/exchange/order/v1"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type OrderServiceServer struct {
	orderv1.UnimplementedOrderServiceServer

	Config             config.Provider
	Logger             app.Logger
	OrderApplication   app.OrderApplication
	AccountApplication app.AccountApplication

	closeListener func() error
	grpcServer    *grpc.Server
	errCh         chan error
}

func (s *OrderServiceServer) Summary() string {
	return "order service gRPC listening on " + s.Config.Snapshot().GRPCAddress
}

func (s *OrderServiceServer) Run() error {
	if s.Config == nil || s.Logger == nil || s.OrderApplication == nil || s.AccountApplication == nil {
		return fmt.Errorf("order service dependencies not configured")
	}

	listener, err := net.Listen("tcp", s.Config.Snapshot().GRPCAddress)
	if err != nil {
		return fmt.Errorf("listen order service grpc: %w", err)
	}

	s.closeListener = listener.Close
	s.grpcServer = grpc.NewServer()
	s.errCh = make(chan error, 1)
	orderv1.RegisterOrderServiceServer(s.grpcServer, s)

	go func() {
		s.Logger.Printf("order service grpc listening addr=%s", s.Config.Snapshot().GRPCAddress)
		if serveErr := s.grpcServer.Serve(listener); serveErr != nil {
			s.errCh <- serveErr
			return
		}
		s.errCh <- nil
	}()

	return nil
}

func (s *OrderServiceServer) Stop() {
	if s.grpcServer != nil {
		s.grpcServer.GracefulStop()
	}
	if s.closeListener != nil {
		_ = s.closeListener()
	}
}

func (s *OrderServiceServer) Wait() error {
	if s.errCh == nil {
		return nil
	}
	return <-s.errCh
}

func (s *OrderServiceServer) PlaceOrder(ctx context.Context, request *orderv1.PlaceOrderRequest) (*orderv1.PlaceOrderResponse, error) {
	userID, err := uuid.Parse(request.GetUserId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	price, err := decimal.NewFromString(request.GetPrice())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	quantity, err := decimal.NewFromString(request.GetQuantity())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	placed, err := s.OrderApplication.PlaceOrder(ctx, app.PlaceOrderInput{
		ClientOrderID: request.GetClientOrderId(),
		UserID:        userID,
		Symbol:        order.Symbol(request.GetSymbol()),
		Side:          order.Side(request.GetSide()),
		Type:          order.Type(request.GetType()),
		Price:         price,
		Quantity:      quantity,
	})
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &orderv1.PlaceOrderResponse{Order: toOrderProto(placed.Order)}, nil
}

func (s *OrderServiceServer) CancelOrder(ctx context.Context, request *orderv1.CancelOrderRequest) (*orderv1.CancelOrderResponse, error) {
	orderID, err := uuid.Parse(request.GetOrderId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	current, err := s.OrderApplication.CancelOrder(ctx, orderID)
	if err != nil {
		if err == app.ErrOrderNotFound {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &orderv1.CancelOrderResponse{Order: toOrderProto(current)}, nil
}

func (s *OrderServiceServer) GetOrder(ctx context.Context, request *orderv1.GetOrderRequest) (*orderv1.GetOrderResponse, error) {
	orderID, err := uuid.Parse(request.GetOrderId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	current, err := s.OrderApplication.GetOrder(ctx, orderID)
	if err != nil {
		if err == app.ErrOrderNotFound {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &orderv1.GetOrderResponse{Order: toOrderProto(current)}, nil
}

func (s *OrderServiceServer) GetBalance(ctx context.Context, request *orderv1.GetBalanceRequest) (*orderv1.GetBalanceResponse, error) {
	userID, err := uuid.Parse(request.GetUserId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	current, err := s.AccountApplication.GetBalance(ctx, userID, request.GetAsset())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &orderv1.GetBalanceResponse{Balance: toBalanceProto(current)}, nil
}
