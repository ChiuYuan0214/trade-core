package bootstrap

import (
	"fmt"

	"github.com/ChiuYuan0214/depin"

	"local.exchange-demo/exchange-core-go/api"
	"local.exchange-demo/exchange-core-go/app"
	"local.exchange-demo/exchange-core-go/postgres"
	sharedinmemory "local.exchange-demo/exchange-core-go/shared/inmemory"
)

func RunOrderService() error {
	return run("order-service", func(rt *runtime) error {
		settings := rt.config.Snapshot()

		state, err := sharedinmemory.SharedOrderState()
		if err != nil {
			return err
		}

		privateEvents := app.PrivateEventPublisher(&app.NoopPrivateEventPublisher{})
		if settings.NotificationGRPCAddress != "" {
			privateEvents = &app.GRPCPrivateEventPublisher{Address: settings.NotificationGRPCAddress}
		}

		orderStore := app.OrderStore(state.OrderStore)
		if settings.OrderStoreBackend == "postgres" {
			if settings.PostgresDSN == "" {
				return fmt.Errorf("POSTGRES_DSN is required when ORDER_STORE_BACKEND=postgres")
			}
			db := &postgres.DB{Config: postgres.Config{DSN: settings.PostgresDSN}}
			depin.RunAndSet[postgres.ConnectionProvider](db)
			depin.Set[postgres.Migrator](&postgres.MigrationRunner{MigrationsDir: settings.MigrationsDir})
			orderStore = &postgres.OrderStore{}
		}

		shardRouter := &api.MatchingEngineRouter{Config: rt.config}
		accountApp := &api.LedgerServiceClient{Config: rt.config}
		orderApp := &app.OrderAppService{
			OrderStore:            orderStore,
			ShardRouter:           shardRouter,
			AccountApplication:    accountApp,
			PrivateEventPublisher: privateEvents,
		}
		service := &api.OrderServiceServer{
			Config:             rt.config,
			Logger:             rt.logger,
			OrderApplication:   orderApp,
			AccountApplication: accountApp,
		}

		depin.Set[app.OrderStore](orderStore)
		depin.Set[app.PrivateEventPublisher](privateEvents)
		depin.Set[app.ShardRouter](shardRouter)
		depin.Set[app.AccountApplication](accountApp)
		depin.Set[app.OrderApplication](orderApp)
		depin.Set[app.Service](service)
		return nil
	})
}
