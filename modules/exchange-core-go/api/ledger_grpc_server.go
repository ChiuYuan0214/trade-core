package api

import (
	"context"
	"fmt"
	"net"

	"local.exchange-demo/exchange-core-go/app"
	"local.exchange-demo/exchange-core-go/config"
	"local.exchange-demo/exchange-core-go/domain/ledger"
	ledgerv1 "local.exchange-demo/exchange-core-go/gen/proto/exchange/ledger/v1"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LedgerServiceServer struct {
	ledgerv1.UnimplementedLedgerServiceServer

	Config             config.Provider
	Logger             app.Logger
	AccountApplication app.AccountApplication

	closeListener func() error
	grpcServer    *grpc.Server
	errCh         chan error
}

func (s *LedgerServiceServer) Summary() string {
	return "ledger service gRPC listening on " + s.Config.Snapshot().GRPCAddress
}

func (s *LedgerServiceServer) Run() error {
	if s.Config == nil || s.Logger == nil || s.AccountApplication == nil {
		return fmt.Errorf("ledger service dependencies not configured")
	}

	listener, err := net.Listen("tcp", s.Config.Snapshot().GRPCAddress)
	if err != nil {
		return fmt.Errorf("listen ledger service grpc: %w", err)
	}

	s.closeListener = listener.Close
	s.grpcServer = grpc.NewServer()
	s.errCh = make(chan error, 1)
	ledgerv1.RegisterLedgerServiceServer(s.grpcServer, s)

	go func() {
		s.Logger.Printf("ledger service grpc listening addr=%s", s.Config.Snapshot().GRPCAddress)
		if serveErr := s.grpcServer.Serve(listener); serveErr != nil {
			s.errCh <- serveErr
			return
		}
		s.errCh <- nil
	}()

	return nil
}

func (s *LedgerServiceServer) Stop() {
	if s.grpcServer != nil {
		s.grpcServer.GracefulStop()
	}
	if s.closeListener != nil {
		_ = s.closeListener()
	}
}

func (s *LedgerServiceServer) Wait() error {
	if s.errCh == nil {
		return nil
	}
	return <-s.errCh
}

func (s *LedgerServiceServer) ReserveFunds(ctx context.Context, request *ledgerv1.ReserveFundsRequest) (*ledgerv1.ReserveFundsResponse, error) {
	userID, err := uuid.Parse(request.GetUserId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	amount, err := decimal.NewFromString(request.GetAmount())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	referenceID, err := uuid.Parse(request.GetReferenceId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if err := s.AccountApplication.Reserve(ctx, app.ReserveFundsInput{
		UserID:        userID,
		Asset:         request.GetAsset(),
		Amount:        amount,
		ReferenceID:   referenceID,
		ReferenceType: ledger.ReferenceType(request.GetReferenceType()),
	}); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &ledgerv1.ReserveFundsResponse{}, nil
}

func (s *LedgerServiceServer) ReleaseFunds(ctx context.Context, request *ledgerv1.ReleaseFundsRequest) (*ledgerv1.ReleaseFundsResponse, error) {
	userID, err := uuid.Parse(request.GetUserId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	amount, err := decimal.NewFromString(request.GetAmount())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	referenceID, err := uuid.Parse(request.GetReferenceId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if err := s.AccountApplication.Release(ctx, app.ReleaseFundsInput{
		UserID:        userID,
		Asset:         request.GetAsset(),
		Amount:        amount,
		ReferenceID:   referenceID,
		ReferenceType: ledger.ReferenceType(request.GetReferenceType()),
	}); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &ledgerv1.ReleaseFundsResponse{}, nil
}

func (s *LedgerServiceServer) ApplyTrade(ctx context.Context, request *ledgerv1.ApplyTradeRequest) (*ledgerv1.ApplyTradeResponse, error) {
	execution, err := fromLedgerTradeProto(request.GetTrade())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	maker, err := fromLedgerOrderProto(request.GetMakerOrder())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	taker, err := fromLedgerOrderProto(request.GetTakerOrder())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if err := s.AccountApplication.ApplyTrade(ctx, execution, maker, taker); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &ledgerv1.ApplyTradeResponse{}, nil
}

func (s *LedgerServiceServer) GetBalance(ctx context.Context, request *ledgerv1.GetBalanceRequest) (*ledgerv1.GetBalanceResponse, error) {
	userID, err := uuid.Parse(request.GetUserId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	current, err := s.AccountApplication.GetBalance(ctx, userID, request.GetAsset())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &ledgerv1.GetBalanceResponse{Balance: toLedgerBalanceProto(current)}, nil
}
