package app

import (
	"context"
	"sync"

	"exchange-demo/internal/domain/ledger"
)

type InMemoryLedgerStore struct {
	mu      sync.RWMutex
	entries []ledger.Entry
}

func (s *InMemoryLedgerStore) Run() error { return nil }
func (s *InMemoryLedgerStore) Stop()      {}

func (s *InMemoryLedgerStore) AppendEntries(_ context.Context, entries ...ledger.Entry) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.entries = append(s.entries, entries...)
	return nil
}

func (s *InMemoryLedgerStore) Entries() []ledger.Entry {
	s.mu.RLock()
	defer s.mu.RUnlock()
	cloned := make([]ledger.Entry, len(s.entries))
	copy(cloned, s.entries)
	return cloned
}
