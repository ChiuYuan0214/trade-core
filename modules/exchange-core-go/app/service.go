package app

import (
	"fmt"

	"local.exchange-demo/exchange-core-go/config"
)

type Service interface {
	Summary() string
	Run() error
	Stop()
}

type ProcessService struct {
	Config config.Provider
	Logger Logger

	processName string
}

func NewProcessService(processName string) *ProcessService {
	return &ProcessService{processName: processName}
}

func (s *ProcessService) Run() error {
	cfg := s.Config.Snapshot()
	s.Logger.Printf("service bootstrapped env=%s http_addr=%s grpc_addr=%s", cfg.Environment, cfg.HTTPAddress, cfg.GRPCAddress)
	return nil
}

func (s *ProcessService) Stop() {
	s.Logger.Printf("service stopped")
}

func (s *ProcessService) Summary() string {
	cfg := s.Config.Snapshot()
	return fmt.Sprintf(
		"bootstrapped process=%s env=%s http_addr=%s grpc_addr=%s",
		s.processName,
		cfg.Environment,
		cfg.HTTPAddress,
		cfg.GRPCAddress,
	)
}
