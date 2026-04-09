package postgres

import (
	"context"
	"errors"

	"local.exchange-demo/exchange-core-go/domain/ledger"
)

type LedgerStore struct {
	Database ConnectionProvider
}

func (s *LedgerStore) Run() error { return nil }
func (s *LedgerStore) Stop()      {}

func (s *LedgerStore) AppendEntries(ctx context.Context, entries ...ledger.Entry) error {
	if s.Database == nil || s.Database.Connection() == nil {
		return errors.New("postgres ledger store db is nil")
	}
	if len(entries) == 0 {
		return nil
	}

	tx, err := s.Database.Connection().BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	for _, entry := range entries {
		if _, err := tx.ExecContext(ctx, `
			INSERT INTO ledger_entries (
				entry_id,
				user_id,
				asset,
				delta_available,
				delta_frozen,
				reference_type,
				reference_id,
				event_id,
				created_at
			) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
			ON CONFLICT (event_id) DO NOTHING
		`,
			entry.ID,
			entry.UserID,
			entry.Asset,
			entry.DeltaAvailable.String(),
			entry.DeltaFrozen.String(),
			entry.ReferenceType,
			entry.ReferenceID,
			entry.EventID,
			entry.CreatedAt,
		); err != nil {
			return err
		}
	}

	return tx.Commit()
}
