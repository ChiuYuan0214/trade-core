package inmemory

import (
	"sync"

	"local.exchange-demo/exchange-core-go/app"
)

type OrderState struct {
	OrderStore   *app.InMemoryOrderStore
	BalanceStore *app.InMemoryBalanceStore
	LedgerStore  *app.InMemoryLedgerStore
}

var (
	orderStateOnce sync.Once
	orderState     *OrderState
	orderStateErr  error
)

func SharedOrderState() (*OrderState, error) {
	orderStateOnce.Do(func() {
		state := &OrderState{
			OrderStore:   &app.InMemoryOrderStore{},
			BalanceStore: &app.InMemoryBalanceStore{},
			LedgerStore:  &app.InMemoryLedgerStore{},
		}
		if err := state.OrderStore.Run(); err != nil {
			orderStateErr = err
			return
		}
		if err := state.BalanceStore.Run(); err != nil {
			orderStateErr = err
			return
		}
		if err := seedDemoBalances(state.BalanceStore); err != nil {
			orderStateErr = err
			return
		}
		orderState = state
	})
	return orderState, orderStateErr
}

func seedDemoBalances(balanceStore *app.InMemoryBalanceStore) error {
	for _, userID := range []string{
		"11111111-1111-1111-1111-111111111111",
		"22222222-2222-2222-2222-222222222222",
	} {
		if err := balanceStore.SeedString(userID, "USDT", "1000000", "0"); err != nil {
			return err
		}
		if err := balanceStore.SeedString(userID, "BTC", "100", "0"); err != nil {
			return err
		}
		if err := balanceStore.SeedString(userID, "ETH", "1000", "0"); err != nil {
			return err
		}
		if err := balanceStore.SeedString(userID, "SOL", "10000", "0"); err != nil {
			return err
		}
	}
	return nil
}
