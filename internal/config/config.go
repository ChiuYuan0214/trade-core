package config

import (
	"fmt"
	"os"
)

type Settings struct {
	ProcessName       string
	Environment       string
	HTTPAddress       string
	GRPCAddress       string
	PostgresDSN       string
	OrderStoreBackend string
	MigrationsDir     string
}

type Provider interface {
	Snapshot() Settings
	Run() error
	Stop()
}

type StaticProvider struct {
	settings Settings
}

func NewStaticProvider(settings Settings) *StaticProvider {
	return &StaticProvider{settings: settings}
}

func Load(processName string) Settings {
	if envProcessName := os.Getenv("PROCESS_NAME"); envProcessName != "" {
		processName = envProcessName
	}
	return Settings{
		ProcessName:       processName,
		Environment:       envOrDefault("APP_ENV", "local"),
		HTTPAddress:       envOrDefault("HTTP_ADDR", defaultHTTPAddress(processName)),
		GRPCAddress:       envOrDefault("GRPC_ADDR", defaultGRPCAddress(processName)),
		PostgresDSN:       os.Getenv("POSTGRES_DSN"),
		OrderStoreBackend: envOrDefault("ORDER_STORE_BACKEND", defaultOrderStoreBackend(processName)),
		MigrationsDir:     envOrDefault("MIGRATIONS_DIR", "migrations"),
	}
}

func (p *StaticProvider) Snapshot() Settings {
	return p.settings
}

func (p *StaticProvider) Run() error {
	return nil
}

func (p *StaticProvider) Stop() {}

func envOrDefault(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func defaultHTTPAddress(processName string) string {
	return fmt.Sprintf("127.0.0.1:%d", defaultPort(processName))
}

func defaultGRPCAddress(processName string) string {
	return fmt.Sprintf("127.0.0.1:%d", defaultPort(processName)+1000)
}

func defaultPort(processName string) int {
	switch processName {
	case "rest-gateway":
		return 8080
	case "ws-gateway":
		return 8081
	case "order-service":
		return 9080
	case "matching-engine":
		return 9081
	case "ledger-service":
		return 9082
	case "market-data-service":
		return 9083
	case "notification-service":
		return 9084
	case "replay-tool":
		return 9085
	default:
		return 9099
	}
}

func defaultOrderStoreBackend(processName string) string {
	switch processName {
	case "rest-gateway", "order-service":
		return "memory"
	default:
		return ""
	}
}
