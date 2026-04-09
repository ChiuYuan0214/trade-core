package bootstrap

import (
	"fmt"

	"github.com/ChiuYuan0214/depin"

	"local.exchange-demo/exchange-core-go/api"
	"local.exchange-demo/exchange-core-go/app"
	"local.exchange-demo/exchange-core-go/postgres"
	sharedinmemory "local.exchange-demo/exchange-core-go/shared/inmemory"
)

func RunLedgerService() error {
	return run("ledger-service", func(rt *runtime) error {
		settings := rt.config.Snapshot()

		privateEvents := app.PrivateEventPublisher(&app.NoopPrivateEventPublisher{})

		var balanceStore app.BalanceStore
		var ledgerStore app.LedgerStore

		switch settings.LedgerStoreBackend {
		case "", "memory":
			state, err := sharedinmemory.SharedLedgerState()
			if err != nil {
				return err
			}
			balanceStore = state.BalanceStore
			ledgerStore = state.LedgerStore
		case "postgres":
			if settings.PostgresDSN == "" {
				return fmt.Errorf("POSTGRES_DSN is required when LEDGER_STORE_BACKEND=postgres")
			}
			db := &postgres.DB{Config: postgres.Config{DSN: settings.PostgresDSN}}
			depin.RunAndSet[postgres.ConnectionProvider](db)
			depin.Set[postgres.Migrator](&postgres.MigrationRunner{MigrationsDir: settings.MigrationsDir})
			depin.Set[*postgres.DemoLedgerSeed](&postgres.DemoLedgerSeed{})
			balanceStore = &postgres.BalanceStore{}
			ledgerStore = &postgres.LedgerStore{}
		default:
			return fmt.Errorf("unsupported LEDGER_STORE_BACKEND: %s", settings.LedgerStoreBackend)
		}

		accountApp := &app.AccountAppService{
			BalanceStore:          balanceStore,
			LedgerStore:           ledgerStore,
			PrivateEventPublisher: privateEvents,
		}
		service := &api.LedgerServiceServer{
			Config:             rt.config,
			Logger:             rt.logger,
			AccountApplication: accountApp,
		}

		depin.Set[app.BalanceStore](balanceStore)
		depin.Set[app.LedgerStore](ledgerStore)
		depin.Set[app.PrivateEventPublisher](privateEvents)
		depin.Set[app.AccountApplication](accountApp)
		depin.Set[app.Service](service)
		return nil
	})
}

func RunWSGateway() error {
	return runGeneric("ws-gateway")
}

func RunReplayTool() error {
	return runGeneric("replay-tool")
}

func runGeneric(processName string) error {
	return run(processName, func(rt *runtime) error {
		depin.Set[app.Service](app.NewProcessService(processName))
		return nil
	})
}
