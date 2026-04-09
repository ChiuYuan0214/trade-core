package bootstrap

import (
	"github.com/ChiuYuan0214/depin"

	"local.exchange-demo/exchange-core-go/api"
	"local.exchange-demo/exchange-core-go/app"
	sharedinmemory "local.exchange-demo/exchange-core-go/shared/inmemory"
)

func RunMatchingEngine() error {
	return run("matching-engine", func(rt *runtime) error {
		state, err := sharedinmemory.SharedMatchingState()
		if err != nil {
			return err
		}

		service := &api.MatchingEngineService{
			Config:      rt.config,
			Logger:      rt.logger,
			ShardRouter: state.Router,
		}
		depin.Set[app.ShardRouter](state.Router)
		depin.Set[app.Service](service)
		return nil
	})
}
