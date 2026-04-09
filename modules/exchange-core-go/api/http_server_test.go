package api

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"local.exchange-demo/exchange-core-go/app"
	"local.exchange-demo/exchange-core-go/domain/account"
	"local.exchange-demo/exchange-core-go/domain/order"
)

func TestPlaceAndGetOrder(t *testing.T) {
	t.Parallel()

	store := &app.InMemoryOrderStore{}
	router := &app.InMemoryShardRouter{}
	balanceStore := &app.InMemoryBalanceStore{}
	ledgerStore := &app.InMemoryLedgerStore{}
	if err := store.Run(); err != nil {
		t.Fatalf("store run: %v", err)
	}
	if err := router.Run(); err != nil {
		t.Fatalf("router run: %v", err)
	}
	if err := balanceStore.Run(); err != nil {
		t.Fatalf("balance store run: %v", err)
	}

	userID := uuid.New()
	if err := balanceStore.Seed(userID, "USDT", "100000", "0"); err != nil {
		t.Fatalf("seed balance: %v", err)
	}

	server := &HTTPServer{
		GatewayApplication: &testGatewayApplication{
			OrderApplication: &app.OrderAppService{
				OrderStore:  store,
				ShardRouter: router,
				AccountApplication: &app.AccountAppService{
					BalanceStore: balanceStore,
					LedgerStore:  ledgerStore,
				},
			},
			AccountApplication: &app.AccountAppService{
				BalanceStore: balanceStore,
				LedgerStore:  ledgerStore,
			},
		},
	}

	body := []byte(`{
		"user_id":"` + userID.String() + `",
		"client_order_id":"demo-1",
		"symbol":"BTC/USDT",
		"side":"BUY",
		"type":"LIMIT",
		"price":"60000",
		"quantity":"0.5"
	}`)

	placeReq := httptest.NewRequest(http.MethodPost, "/api/v1/orders", bytes.NewReader(body))
	placeResp := httptest.NewRecorder()
	server.Handler().ServeHTTP(placeResp, placeReq)

	if placeResp.Code != http.StatusAccepted {
		t.Fatalf("expected 202, got %d: %s", placeResp.Code, placeResp.Body.String())
	}

	var placed orderResponse
	if err := json.Unmarshal(placeResp.Body.Bytes(), &placed); err != nil {
		t.Fatalf("unmarshal place response: %v", err)
	}

	getReq := httptest.NewRequest(http.MethodGet, "/api/v1/orders/"+placed.OrderID, nil)
	getResp := httptest.NewRecorder()
	server.Handler().ServeHTTP(getResp, getReq)

	if getResp.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", getResp.Code, getResp.Body.String())
	}
}

func TestCancelOrder(t *testing.T) {
	t.Parallel()

	store := &app.InMemoryOrderStore{}
	router := &app.InMemoryShardRouter{}
	balanceStore := &app.InMemoryBalanceStore{}
	ledgerStore := &app.InMemoryLedgerStore{}
	_ = store.Run()
	_ = router.Run()
	_ = balanceStore.Run()
	_ = ledgerStore.Run()

	userID := uuid.New()
	if err := balanceStore.Seed(userID, "USDT", "100000", "0"); err != nil {
		t.Fatalf("seed balance: %v", err)
	}

	service := &app.OrderAppService{
		OrderStore:  store,
		ShardRouter: router,
		AccountApplication: &app.AccountAppService{
			BalanceStore: balanceStore,
			LedgerStore:  ledgerStore,
		},
	}
	server := &HTTPServer{GatewayApplication: &testGatewayApplication{
		OrderApplication: service,
		AccountApplication: &app.AccountAppService{
			BalanceStore: balanceStore,
			LedgerStore:  ledgerStore,
		},
	}}

	placed, err := service.PlaceOrder(context.Background(), app.PlaceOrderInput{
		ClientOrderID: "demo-2",
		UserID:        userID,
		Symbol:        "BTC/USDT",
		Side:          "BUY",
		Type:          "LIMIT",
		Price:         mustDecimal(t, "59000"),
		Quantity:      mustDecimal(t, "1"),
	})
	if err != nil {
		t.Fatalf("place order: %v", err)
	}

	cancelReq := httptest.NewRequest(http.MethodDelete, "/api/v1/orders/"+placed.Order.ID.String(), nil)
	cancelResp := httptest.NewRecorder()
	server.Handler().ServeHTTP(cancelResp, cancelReq)

	if cancelResp.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", cancelResp.Code, cancelResp.Body.String())
	}
}

func (a *testGatewayApplication) Run() error { return nil }
func (a *testGatewayApplication) Stop()      {}

type testGatewayApplication struct {
	OrderApplication   app.OrderApplication
	AccountApplication app.AccountApplication
}

func (a *testGatewayApplication) PlaceOrder(ctx context.Context, input app.PlaceOrderInput) (app.PlaceOrderResult, error) {
	return a.OrderApplication.PlaceOrder(ctx, input)
}

func (a *testGatewayApplication) CancelOrder(ctx context.Context, orderID string) (order.Order, error) {
	parsed, err := uuid.Parse(orderID)
	if err != nil {
		return order.Order{}, err
	}
	return a.OrderApplication.CancelOrder(ctx, parsed)
}

func (a *testGatewayApplication) GetOrder(ctx context.Context, orderID string) (order.Order, error) {
	parsed, err := uuid.Parse(orderID)
	if err != nil {
		return order.Order{}, err
	}
	return a.OrderApplication.GetOrder(ctx, parsed)
}

func (a *testGatewayApplication) GetBalance(ctx context.Context, userID string, asset string) (account.Balance, error) {
	parsed, err := uuid.Parse(userID)
	if err != nil {
		return account.Balance{}, err
	}
	return a.AccountApplication.GetBalance(ctx, parsed, asset)
}

func mustDecimal(t *testing.T, raw string) (value decimal.Decimal) {
	t.Helper()

	value, err := decimal.NewFromString(raw)
	if err != nil {
		t.Fatalf("parse decimal %q: %v", raw, err)
	}
	return value
}
