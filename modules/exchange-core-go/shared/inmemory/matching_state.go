package inmemory

import (
	"sync"

	"local.exchange-demo/exchange-core-go/app"
)

type MatchingState struct {
	Router *app.InMemoryShardRouter
}

var (
	matchingStateOnce sync.Once
	matchingState     *MatchingState
	matchingStateErr  error
)

func SharedMatchingState() (*MatchingState, error) {
	matchingStateOnce.Do(func() {
		state := &MatchingState{
			Router: &app.InMemoryShardRouter{},
		}
		if err := state.Router.Run(); err != nil {
			matchingStateErr = err
			return
		}
		matchingState = state
	})
	return matchingState, matchingStateErr
}
