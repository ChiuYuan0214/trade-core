package bootstrap

import (
	"fmt"

	"github.com/ChiuYuan0214/depin"

	"exchange-demo/internal/app"
	"exchange-demo/internal/api"
	"exchange-demo/internal/config"
	"exchange-demo/internal/postgres"
)

// RunProcess wires the minimal process graph with depin so each binary starts
// from the same bootstrap convention before business services are added.
func RunProcess(processName string) error {
	depin.Reset(depin.Global)

	cfg := config.NewStaticProvider(config.Load(processName))
	logger := app.NewStdLogger(processName)

	depin.Set[config.Provider](cfg)
	depin.Set[app.Logger](logger)

	switch processName {
	case "rest-gateway":
		if err := registerRESTGateway(cfg); err != nil {
			return err
		}
	default:
		service := app.NewProcessService(processName)
		depin.Set[app.Service](service)
	}

	depin.Run()
	defer depin.Stop()

	booted, ok := depin.GetGlobal[app.Service]()
	if !ok {
		return fmt.Errorf("service %q not found in depin container", processName)
	}

	fmt.Println(booted.Summary())
	if blocking, ok := booted.(interface{ Wait() error }); ok {
		return blocking.Wait()
	}
	return nil
}

func registerRESTGateway(cfg *config.StaticProvider) error {
	settings := cfg.Snapshot()
	balanceStore := &app.InMemoryBalanceStore{}
	if err := balanceStore.Run(); err != nil {
		return err
	}
	if err := seedDemoBalances(balanceStore); err != nil {
		return err
	}

	if settings.OrderStoreBackend == "postgres" {
		if settings.PostgresDSN == "" {
			return fmt.Errorf("POSTGRES_DSN is required when ORDER_STORE_BACKEND=postgres")
		}
		db := &postgres.DB{Config: postgres.Config{DSN: settings.PostgresDSN}}
		depin.RunAndSet[postgres.ConnectionProvider](db)
		depin.Set[postgres.Migrator](&postgres.MigrationRunner{MigrationsDir: settings.MigrationsDir})
		depin.Set[app.OrderStore](&postgres.OrderStore{})
	} else {
		depin.Set[app.OrderStore](&app.InMemoryOrderStore{})
	}

	depin.Set[app.BalanceStore](balanceStore)
	depin.Set[app.LedgerStore](&app.InMemoryLedgerStore{})
	depin.Set[app.AccountApplication](&app.AccountAppService{})
	depin.Set[app.ShardRouter](&app.InMemoryShardRouter{})
	depin.Set[app.OrderApplication](&app.OrderAppService{})
	depin.Set[api.HandlerProvider](&api.HTTPServer{})
	depin.Set[app.Service](&api.GatewayService{})
	return nil
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
