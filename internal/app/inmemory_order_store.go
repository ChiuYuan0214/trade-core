package app

import (
	"context"
	"sync"

	"github.com/google/uuid"

	"exchange-demo/internal/domain/order"
)

type InMemoryOrderStore struct {
	mu     sync.RWMutex
	orders map[uuid.UUID]order.Order
}

func (s *InMemoryOrderStore) Run() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.orders == nil {
		s.orders = make(map[uuid.UUID]order.Order)
	}
	return nil
}

func (s *InMemoryOrderStore) Stop() {}

func (s *InMemoryOrderStore) Save(_ context.Context, current order.Order) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.orders == nil {
		s.orders = make(map[uuid.UUID]order.Order)
	}
	s.orders[current.ID] = current
	return nil
}

func (s *InMemoryOrderStore) Get(_ context.Context, orderID uuid.UUID) (order.Order, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	current, ok := s.orders[orderID]
	if !ok {
		return order.Order{}, ErrOrderNotFound
	}
	return current, nil
}
