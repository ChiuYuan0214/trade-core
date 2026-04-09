package inmemory

import (
	"sync"

	"local.exchange-demo/exchange-core-go/app"
)

type LedgerState struct {
	BalanceStore *app.InMemoryBalanceStore
	LedgerStore  *app.InMemoryLedgerStore
}

var (
	ledgerStateOnce sync.Once
	ledgerState     *LedgerState
	ledgerStateErr  error
)

func SharedLedgerState() (*LedgerState, error) {
	ledgerStateOnce.Do(func() {
		state := &LedgerState{
			BalanceStore: &app.InMemoryBalanceStore{},
			LedgerStore:  &app.InMemoryLedgerStore{},
		}
		if err := state.BalanceStore.Run(); err != nil {
			ledgerStateErr = err
			return
		}
		if err := seedDemoBalances(state.BalanceStore); err != nil {
			ledgerStateErr = err
			return
		}
		ledgerState = state
	})
	return ledgerState, ledgerStateErr
}
