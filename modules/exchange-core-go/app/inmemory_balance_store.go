package app

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"local.exchange-demo/exchange-core-go/domain/account"
)

type InMemoryBalanceStore struct {
	mu       sync.RWMutex
	balances map[string]account.Balance
}

func (s *InMemoryBalanceStore) Run() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.balances == nil {
		s.balances = make(map[string]account.Balance)
	}
	return nil
}

func (s *InMemoryBalanceStore) Stop() {}

func (s *InMemoryBalanceStore) GetBalance(_ context.Context, userID uuid.UUID, asset string) (account.Balance, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	current, ok := s.balances[balanceKey(userID, asset)]
	if !ok {
		return account.Balance{}, fmt.Errorf("balance not found for user=%s asset=%s", userID, asset)
	}
	return current, nil
}

func (s *InMemoryBalanceStore) SaveBalance(_ context.Context, balance account.Balance) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.balances == nil {
		s.balances = make(map[string]account.Balance)
	}
	s.balances[balanceKey(balance.UserID, balance.Asset)] = balance
	return nil
}

func (s *InMemoryBalanceStore) Seed(userID uuid.UUID, asset string, available string, frozen string) error {
	availableAmount, err := decimal.NewFromString(available)
	if err != nil {
		return err
	}
	frozenAmount, err := decimal.NewFromString(frozen)
	if err != nil {
		return err
	}
	return s.SaveBalance(context.Background(), account.Balance{
		UserID:          userID,
		Asset:           asset,
		AvailableAmount: availableAmount,
		FrozenAmount:    frozenAmount,
		UpdatedAt:       time.Now().UTC(),
	})
}

func (s *InMemoryBalanceStore) SeedString(rawUserID string, asset string, available string, frozen string) error {
	userID, err := uuid.Parse(rawUserID)
	if err != nil {
		return err
	}
	return s.Seed(userID, asset, available, frozen)
}

func balanceKey(userID uuid.UUID, asset string) string {
	return userID.String() + ":" + asset
}
