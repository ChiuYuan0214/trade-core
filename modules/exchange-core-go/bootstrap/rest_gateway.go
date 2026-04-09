package bootstrap

import (
	"github.com/ChiuYuan0214/depin"

	"local.exchange-demo/exchange-core-go/api"
	"local.exchange-demo/exchange-core-go/app"
)

func RunRESTGateway() error {
	return run("rest-gateway", func(rt *runtime) error {
		client := &api.OrderServiceClient{Config: rt.config}
		httpServer := &api.HTTPServer{GatewayApplication: client}
		service := &api.GatewayService{
			Config:     rt.config,
			Logger:     rt.logger,
			HTTPServer: httpServer,
		}

		depin.Set[api.GatewayApplication](client)
		depin.Set[api.HandlerProvider](httpServer)
		depin.Set[app.Service](service)
		return nil
	})
}
