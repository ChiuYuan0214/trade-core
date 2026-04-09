package app

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"exchange-demo/internal/domain/order"
)

func TestPlaceOrderPersistsMatchedState(t *testing.T) {
	t.Parallel()

	store := &InMemoryOrderStore{}
	router := &InMemoryShardRouter{}
	balanceStore := &InMemoryBalanceStore{}
	ledgerStore := &InMemoryLedgerStore{}
	_ = store.Run()
	_ = router.Run()
	_ = balanceStore.Run()
	_ = ledgerStore.Run()

	makerUserID := uuid.New()
	takerUserID := uuid.New()
	if err := balanceStore.Seed(makerUserID, "BTC", "10", "0"); err != nil {
		t.Fatalf("seed maker balance: %v", err)
	}
	if err := balanceStore.Seed(makerUserID, "USDT", "0", "0"); err != nil {
		t.Fatalf("seed maker quote balance: %v", err)
	}
	if err := balanceStore.Seed(takerUserID, "USDT", "100000", "0"); err != nil {
		t.Fatalf("seed taker balance: %v", err)
	}
	if err := balanceStore.Seed(takerUserID, "BTC", "0", "0"); err != nil {
		t.Fatalf("seed taker base balance: %v", err)
	}

	service := &OrderAppService{
		OrderStore:  store,
		ShardRouter: router,
		AccountApplication: &AccountAppService{
			BalanceStore: balanceStore,
			LedgerStore:  ledgerStore,
		},
	}

	maker, err := service.PlaceOrder(context.Background(), PlaceOrderInput{
		UserID:   makerUserID,
		Symbol:   order.SymbolBTCUSDT,
		Side:     order.SideSell,
		Type:     order.TypeLimit,
		Price:    decimal.RequireFromString("60000"),
		Quantity: decimal.RequireFromString("1"),
	})
	if err != nil {
		t.Fatalf("place maker: %v", err)
	}
	if maker.Order.Status != order.StatusOpen {
		t.Fatalf("expected maker to be open, got %s", maker.Order.Status)
	}

	taker, err := service.PlaceOrder(context.Background(), PlaceOrderInput{
		UserID:   takerUserID,
		Symbol:   order.SymbolBTCUSDT,
		Side:     order.SideBuy,
		Type:     order.TypeLimit,
		Price:    decimal.RequireFromString("60000"),
		Quantity: decimal.RequireFromString("1"),
	})
	if err != nil {
		t.Fatalf("place taker: %v", err)
	}
	if taker.Order.Status != order.StatusFilled {
		t.Fatalf("expected taker to be filled, got %s", taker.Order.Status)
	}

	updatedMaker, err := store.Get(context.Background(), maker.Order.ID)
	if err != nil {
		t.Fatalf("get maker: %v", err)
	}
	if updatedMaker.Status != order.StatusFilled {
		t.Fatalf("expected maker to be filled after match, got %s", updatedMaker.Status)
	}
}

func TestPlaceOrderRejectsInsufficientFunds(t *testing.T) {
	t.Parallel()

	store := &InMemoryOrderStore{}
	router := &InMemoryShardRouter{}
	balanceStore := &InMemoryBalanceStore{}
	ledgerStore := &InMemoryLedgerStore{}
	_ = store.Run()
	_ = router.Run()
	_ = balanceStore.Run()
	_ = ledgerStore.Run()

	userID := uuid.New()
	if err := balanceStore.Seed(userID, "USDT", "10", "0"); err != nil {
		t.Fatalf("seed user balance: %v", err)
	}

	service := &OrderAppService{
		OrderStore:  store,
		ShardRouter: router,
		AccountApplication: &AccountAppService{
			BalanceStore: balanceStore,
			LedgerStore:  ledgerStore,
		},
	}

	_, err := service.PlaceOrder(context.Background(), PlaceOrderInput{
		UserID:   userID,
		Symbol:   order.SymbolBTCUSDT,
		Side:     order.SideBuy,
		Type:     order.TypeLimit,
		Price:    decimal.RequireFromString("60000"),
		Quantity: decimal.RequireFromString("1"),
	})
	if err == nil {
		t.Fatalf("expected insufficient funds error")
	}
}

func TestCancelOrderReleasesReservedFunds(t *testing.T) {
	t.Parallel()

	store := &InMemoryOrderStore{}
	router := &InMemoryShardRouter{}
	balanceStore := &InMemoryBalanceStore{}
	ledgerStore := &InMemoryLedgerStore{}
	_ = store.Run()
	_ = router.Run()
	_ = balanceStore.Run()
	_ = ledgerStore.Run()

	userID := uuid.New()
	if err := balanceStore.Seed(userID, "USDT", "100000", "0"); err != nil {
		t.Fatalf("seed user balance: %v", err)
	}

	service := &OrderAppService{
		OrderStore:  store,
		ShardRouter: router,
		AccountApplication: &AccountAppService{
			BalanceStore: balanceStore,
			LedgerStore:  ledgerStore,
		},
	}

	placed, err := service.PlaceOrder(context.Background(), PlaceOrderInput{
		UserID:   userID,
		Symbol:   order.SymbolBTCUSDT,
		Side:     order.SideBuy,
		Type:     order.TypeLimit,
		Price:    decimal.RequireFromString("60000"),
		Quantity: decimal.RequireFromString("1"),
	})
	if err != nil {
		t.Fatalf("place order: %v", err)
	}

	reserved, err := balanceStore.GetBalance(context.Background(), userID, "USDT")
	if err != nil {
		t.Fatalf("get reserved balance: %v", err)
	}
	if !reserved.AvailableAmount.Equal(decimal.RequireFromString("40000")) {
		t.Fatalf("unexpected available after reserve: %s", reserved.AvailableAmount)
	}
	if !reserved.FrozenAmount.Equal(decimal.RequireFromString("60000")) {
		t.Fatalf("unexpected frozen after reserve: %s", reserved.FrozenAmount)
	}

	if _, err := service.CancelOrder(context.Background(), placed.Order.ID); err != nil {
		t.Fatalf("cancel order: %v", err)
	}

	released, err := balanceStore.GetBalance(context.Background(), userID, "USDT")
	if err != nil {
		t.Fatalf("get released balance: %v", err)
	}
	if !released.AvailableAmount.Equal(decimal.RequireFromString("100000")) {
		t.Fatalf("unexpected available after release: %s", released.AvailableAmount)
	}
	if !released.FrozenAmount.Equal(decimal.Zero) {
		t.Fatalf("unexpected frozen after release: %s", released.FrozenAmount)
	}
}

func TestTradeSettlementMovesBalances(t *testing.T) {
	t.Parallel()

	store := &InMemoryOrderStore{}
	router := &InMemoryShardRouter{}
	balanceStore := &InMemoryBalanceStore{}
	ledgerStore := &InMemoryLedgerStore{}
	_ = store.Run()
	_ = router.Run()
	_ = balanceStore.Run()
	_ = ledgerStore.Run()

	sellerUserID := uuid.New()
	buyerUserID := uuid.New()
	if err := balanceStore.Seed(sellerUserID, "BTC", "2", "0"); err != nil {
		t.Fatalf("seed seller BTC: %v", err)
	}
	if err := balanceStore.Seed(sellerUserID, "USDT", "0", "0"); err != nil {
		t.Fatalf("seed seller USDT: %v", err)
	}
	if err := balanceStore.Seed(buyerUserID, "USDT", "100000", "0"); err != nil {
		t.Fatalf("seed buyer USDT: %v", err)
	}
	if err := balanceStore.Seed(buyerUserID, "BTC", "0", "0"); err != nil {
		t.Fatalf("seed buyer BTC: %v", err)
	}

	service := &OrderAppService{
		OrderStore:  store,
		ShardRouter: router,
		AccountApplication: &AccountAppService{
			BalanceStore: balanceStore,
			LedgerStore:  ledgerStore,
		},
	}

	_, err := service.PlaceOrder(context.Background(), PlaceOrderInput{
		UserID:   sellerUserID,
		Symbol:   order.SymbolBTCUSDT,
		Side:     order.SideSell,
		Type:     order.TypeLimit,
		Price:    decimal.RequireFromString("59000"),
		Quantity: decimal.RequireFromString("1"),
	})
	if err != nil {
		t.Fatalf("place seller order: %v", err)
	}

	_, err = service.PlaceOrder(context.Background(), PlaceOrderInput{
		UserID:   buyerUserID,
		Symbol:   order.SymbolBTCUSDT,
		Side:     order.SideBuy,
		Type:     order.TypeLimit,
		Price:    decimal.RequireFromString("60000"),
		Quantity: decimal.RequireFromString("1"),
	})
	if err != nil {
		t.Fatalf("place buyer order: %v", err)
	}

	buyerUSDT, _ := balanceStore.GetBalance(context.Background(), buyerUserID, "USDT")
	buyerBTC, _ := balanceStore.GetBalance(context.Background(), buyerUserID, "BTC")
	sellerBTC, _ := balanceStore.GetBalance(context.Background(), sellerUserID, "BTC")
	sellerUSDT, _ := balanceStore.GetBalance(context.Background(), sellerUserID, "USDT")

	if !buyerUSDT.AvailableAmount.Equal(decimal.RequireFromString("41000")) {
		t.Fatalf("unexpected buyer available USDT: %s", buyerUSDT.AvailableAmount)
	}
	if !buyerUSDT.FrozenAmount.Equal(decimal.Zero) {
		t.Fatalf("unexpected buyer frozen USDT: %s", buyerUSDT.FrozenAmount)
	}
	if !buyerBTC.AvailableAmount.Equal(decimal.RequireFromString("1")) {
		t.Fatalf("unexpected buyer BTC: %s", buyerBTC.AvailableAmount)
	}
	if !sellerBTC.AvailableAmount.Equal(decimal.RequireFromString("1")) {
		t.Fatalf("unexpected seller available BTC: %s", sellerBTC.AvailableAmount)
	}
	if !sellerBTC.FrozenAmount.Equal(decimal.Zero) {
		t.Fatalf("unexpected seller frozen BTC: %s", sellerBTC.FrozenAmount)
	}
	if !sellerUSDT.AvailableAmount.Equal(decimal.RequireFromString("59000")) {
		t.Fatalf("unexpected seller USDT: %s", sellerUSDT.AvailableAmount)
	}
}
