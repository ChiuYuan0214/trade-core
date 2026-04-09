package bootstrap

import (
	"fmt"

	"github.com/ChiuYuan0214/depin"

	"local.exchange-demo/exchange-core-go/app"
	"local.exchange-demo/exchange-core-go/config"
)

type runtime struct {
	processName string
	config      *config.StaticProvider
	logger      app.Logger
}

func newRuntime(processName string) *runtime {
	cfg := config.NewStaticProvider(config.Load(processName))
	return &runtime{
		processName: processName,
		config:      cfg,
		logger:      app.NewStdLogger(processName),
	}
}

func run(processName string, register func(*runtime) error) error {
	depin.Reset(depin.Global)

	rt := newRuntime(processName)
	depin.Set[config.Provider](rt.config)
	depin.Set[app.Logger](rt.logger)

	if register != nil {
		if err := register(rt); err != nil {
			return err
		}
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
