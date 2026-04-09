package engine

import (
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"

	"exchange-demo/internal/domain/order"
	"exchange-demo/internal/matching/book"
)

type Shard struct {
	mu       sync.Mutex
	shardID  string
	symbol   order.Symbol
	book     *book.Book
	sequence uint64
}

func NewShard(shardID string, symbol order.Symbol) (*Shard, error) {
	if shardID == "" {
		return nil, fmt.Errorf("shard_id is required")
	}
	orderBook, err := book.New(symbol)
	if err != nil {
		return nil, err
	}
	return &Shard{
		shardID: shardID,
		symbol:  symbol,
		book:    orderBook,
	}, nil
}

func (s *Shard) Apply(incoming order.Order, now time.Time) (book.ApplyResult, uint64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	result, err := s.book.Apply(incoming, now)
	if err != nil {
		return book.ApplyResult{}, s.sequence, err
	}
	s.sequence++
	return result, s.sequence, nil
}

func (s *Shard) Cancel(orderID uuid.UUID, now time.Time) (*order.Order, uint64, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	canceled, ok := s.book.Cancel(orderID, now)
	if !ok {
		return nil, s.sequence, false
	}
	s.sequence++
	return canceled, s.sequence, true
}

func (s *Shard) Snapshot(depth int) ([]book.Level, []book.Level, uint64) {
	s.mu.Lock()
	defer s.mu.Unlock()

	bids, asks := s.book.Snapshot(depth)
	return bids, asks, s.sequence
}

func (s *Shard) Symbol() order.Symbol {
	return s.symbol
}

func (s *Shard) ID() string {
	return s.shardID
}
