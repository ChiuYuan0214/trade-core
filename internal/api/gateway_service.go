package api

import (
	"context"
	"errors"
	"net/http"
	"time"

	"exchange-demo/internal/app"
	"exchange-demo/internal/config"
)

type GatewayService struct {
	Config      config.Provider
	Logger      app.Logger
	HTTPServer  HandlerProvider

	server *http.Server
	errCh  chan error
}

func (s *GatewayService) Summary() string {
	cfg := s.Config.Snapshot()
	return "rest gateway listening on " + cfg.HTTPAddress
}

func (s *GatewayService) Run() error {
	if s.Config == nil || s.Logger == nil || s.HTTPServer == nil {
		return errors.New("gateway dependencies not configured")
	}

	cfg := s.Config.Snapshot()
	s.server = &http.Server{
		Addr:              cfg.HTTPAddress,
		Handler:           s.HTTPServer.Handler(),
		ReadHeaderTimeout: 5 * time.Second,
	}
	s.errCh = make(chan error, 1)

	go func() {
		s.Logger.Printf("http server listening addr=%s", cfg.HTTPAddress)
		err := s.server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.errCh <- err
			return
		}
		s.errCh <- nil
	}()

	return nil
}

func (s *GatewayService) Stop() {
	if s.server == nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = s.server.Shutdown(ctx)
}

func (s *GatewayService) Wait() error {
	if s.errCh == nil {
		return nil
	}
	return <-s.errCh
}
