package api

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/google/uuid"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"local.exchange-demo/exchange-core-go/app"
	"local.exchange-demo/exchange-core-go/config"
	"local.exchange-demo/exchange-core-go/domain/order"
	matchingv1 "local.exchange-demo/exchange-core-go/gen/proto/exchange/matching/v1"
)

type MatchingEngineService struct {
	matchingv1.UnimplementedMatchingEngineServiceServer

	Config      config.Provider
	Logger      app.Logger
	ShardRouter app.ShardRouter

	closeListener func() error
	grpcServer    *grpc.Server
	errCh         chan error
}

func (s *MatchingEngineService) Summary() string {
	return "matching engine gRPC listening on " + s.Config.Snapshot().GRPCAddress
}

func (s *MatchingEngineService) Run() error {
	if s.Config == nil || s.Logger == nil || s.ShardRouter == nil {
		return fmt.Errorf("matching engine dependencies not configured")
	}

	listener, err := net.Listen("tcp", s.Config.Snapshot().GRPCAddress)
	if err != nil {
		return fmt.Errorf("listen matching engine grpc: %w", err)
	}

	s.closeListener = listener.Close
	s.grpcServer = grpc.NewServer()
	s.errCh = make(chan error, 1)
	matchingv1.RegisterMatchingEngineServiceServer(s.grpcServer, s)

	go func() {
		s.Logger.Printf("matching engine grpc listening addr=%s", s.Config.Snapshot().GRPCAddress)
		if serveErr := s.grpcServer.Serve(listener); serveErr != nil {
			s.errCh <- serveErr
			return
		}
		s.errCh <- nil
	}()

	return nil
}

func (s *MatchingEngineService) Stop() {
	if s.grpcServer != nil {
		s.grpcServer.GracefulStop()
	}
	if s.closeListener != nil {
		_ = s.closeListener()
	}
}

func (s *MatchingEngineService) Wait() error {
	if s.errCh == nil {
		return nil
	}
	return <-s.errCh
}

func (s *MatchingEngineService) ApplyOrder(ctx context.Context, request *matchingv1.ApplyOrderRequest) (*matchingv1.ApplyOrderResponse, error) {
	incoming, err := fromMatchingOrderProto(request.GetOrder())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	shard, err := s.ShardRouter.ForSymbol(incoming.Symbol)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	result, sequence, err := shard.Apply(incoming, incoming.UpdatedAt)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	response := &matchingv1.ApplyOrderResponse{
		IncomingOrder: toMatchingOrderProto(result.IncomingOrder),
		Sequence:      sequence,
	}
	for _, updated := range result.OrderUpdates {
		response.OrderUpdates = append(response.OrderUpdates, toMatchingOrderProto(updated))
	}
	for _, execution := range result.Trades {
		response.Trades = append(response.Trades, toTradeProto(execution))
	}
	return response, nil
}

func (s *MatchingEngineService) CancelOrder(ctx context.Context, request *matchingv1.CancelOrderRequest) (*matchingv1.CancelOrderResponse, error) {
	symbol := order.Symbol(request.GetSymbol())
	if err := symbol.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	orderID, err := uuid.Parse(request.GetOrderId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	shard, err := s.ShardRouter.ForSymbol(symbol)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	current, sequence, found := shard.Cancel(orderID, time.Now().UTC())
	if !found {
		return &matchingv1.CancelOrderResponse{Found: false, Sequence: sequence}, nil
	}
	return &matchingv1.CancelOrderResponse{
		Found:    true,
		Order:    toMatchingOrderProto(*current),
		Sequence: sequence,
	}, nil
}
