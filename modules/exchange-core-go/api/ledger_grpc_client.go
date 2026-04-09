package api

import (
	"context"
	"fmt"
	"time"

	"local.exchange-demo/exchange-core-go/app"
	"local.exchange-demo/exchange-core-go/config"
	"local.exchange-demo/exchange-core-go/domain/account"
	"local.exchange-demo/exchange-core-go/domain/order"
	"local.exchange-demo/exchange-core-go/domain/trade"
	ledgerv1 "local.exchange-demo/exchange-core-go/gen/proto/exchange/ledger/v1"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type LedgerServiceClient struct {
	Config config.Provider

	Address string
	Timeout time.Duration

	conn         *grpc.ClientConn
	reserveFn    func(context.Context, *ledgerv1.ReserveFundsRequest, ...grpc.CallOption) (*ledgerv1.ReserveFundsResponse, error)
	releaseFn    func(context.Context, *ledgerv1.ReleaseFundsRequest, ...grpc.CallOption) (*ledgerv1.ReleaseFundsResponse, error)
	applyTradeFn func(context.Context, *ledgerv1.ApplyTradeRequest, ...grpc.CallOption) (*ledgerv1.ApplyTradeResponse, error)
	getBalanceFn func(context.Context, *ledgerv1.GetBalanceRequest, ...grpc.CallOption) (*ledgerv1.GetBalanceResponse, error)
}

func (c *LedgerServiceClient) Run() error {
	if c.Config != nil {
		settings := c.Config.Snapshot()
		if c.Address == "" {
			c.Address = settings.LedgerServiceGRPCAddress
		}
	}
	if c.Address == "" {
		return fmt.Errorf("ledger service grpc address is required")
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
		return fmt.Errorf("dial ledger service grpc: %w", err)
	}

	c.conn = conn
	client := ledgerv1.NewLedgerServiceClient(conn)
	c.reserveFn = client.ReserveFunds
	c.releaseFn = client.ReleaseFunds
	c.applyTradeFn = client.ApplyTrade
	c.getBalanceFn = client.GetBalance
	return nil
}

func (c *LedgerServiceClient) Stop() {
	if c.conn != nil {
		_ = c.conn.Close()
	}
}

func (c *LedgerServiceClient) Reserve(ctx context.Context, input app.ReserveFundsInput) error {
	if c.reserveFn == nil {
		return fmt.Errorf("ledger service client is not initialized")
	}
	_, err := c.reserveFn(ctx, &ledgerv1.ReserveFundsRequest{
		UserId:        input.UserID.String(),
		Asset:         input.Asset,
		Amount:        input.Amount.String(),
		ReferenceId:   input.ReferenceID.String(),
		ReferenceType: string(input.ReferenceType),
	})
	return err
}

func (c *LedgerServiceClient) Release(ctx context.Context, input app.ReleaseFundsInput) error {
	if c.releaseFn == nil {
		return fmt.Errorf("ledger service client is not initialized")
	}
	_, err := c.releaseFn(ctx, &ledgerv1.ReleaseFundsRequest{
		UserId:        input.UserID.String(),
		Asset:         input.Asset,
		Amount:        input.Amount.String(),
		ReferenceId:   input.ReferenceID.String(),
		ReferenceType: string(input.ReferenceType),
	})
	return err
}

func (c *LedgerServiceClient) GetBalance(ctx context.Context, userID uuid.UUID, asset string) (account.Balance, error) {
	if c.getBalanceFn == nil {
		return account.Balance{}, fmt.Errorf("ledger service client is not initialized")
	}
	response, err := c.getBalanceFn(ctx, &ledgerv1.GetBalanceRequest{
		UserId: userID.String(),
		Asset:  asset,
	})
	if err != nil {
		return account.Balance{}, err
	}
	return fromLedgerBalanceProto(response.GetBalance())
}

func (c *LedgerServiceClient) ApplyTrade(ctx context.Context, execution trade.Trade, maker order.Order, taker order.Order) error {
	if c.applyTradeFn == nil {
		return fmt.Errorf("ledger service client is not initialized")
	}
	_, err := c.applyTradeFn(ctx, &ledgerv1.ApplyTradeRequest{
		Trade:      toLedgerTradeProto(execution),
		MakerOrder: toLedgerOrderProto(maker),
		TakerOrder: toLedgerOrderProto(taker),
	})
	return err
}

var _ app.AccountApplication = (*LedgerServiceClient)(nil)
