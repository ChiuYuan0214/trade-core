package app

import (
	"fmt"

	"local.exchange-demo/exchange-core-go/domain/order"
	"local.exchange-demo/exchange-core-go/matching/engine"
)

type InMemoryShardRouter struct {
	Shards map[order.Symbol]*engine.Shard
}

func (r *InMemoryShardRouter) Run() error {
	if r.Shards != nil {
		return nil
	}

	r.Shards = make(map[order.Symbol]*engine.Shard, len(order.SupportedSymbols()))
	for index, symbol := range order.SupportedSymbols() {
		shard, err := engine.NewShard(fmt.Sprintf("shard-%d", index+1), symbol)
		if err != nil {
			return err
		}
		r.Shards[symbol] = shard
	}
	return nil
}

func (r *InMemoryShardRouter) Stop() {}

func (r *InMemoryShardRouter) ForSymbol(symbol order.Symbol) (ShardMatcher, error) {
	if r.Shards == nil {
		if err := r.Run(); err != nil {
			return nil, err
		}
	}

	shard, ok := r.Shards[symbol]
	if !ok {
		return nil, fmt.Errorf("no shard configured for symbol %s", symbol)
	}
	return shard, nil
}
