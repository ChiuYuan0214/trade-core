package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"local.exchange-demo/exchange-core-go/domain/account"
)

type BalanceStore struct {
	Database ConnectionProvider
}

func (s *BalanceStore) Run() error { return nil }
func (s *BalanceStore) Stop()      {}

func (s *BalanceStore) GetBalance(ctx context.Context, userID uuid.UUID, asset string) (account.Balance, error) {
	if s.Database == nil || s.Database.Connection() == nil {
		return account.Balance{}, errors.New("postgres balance store db is nil")
	}

	row := s.Database.Connection().QueryRowContext(ctx, `
		SELECT
			user_id,
			asset,
			available_balance,
			frozen_balance,
			updated_at
		FROM accounts
		WHERE user_id = $1 AND asset = $2
	`, userID, asset)

	var current account.Balance
	var availableRaw string
	var frozenRaw string

	if err := row.Scan(
		&current.UserID,
		&current.Asset,
		&availableRaw,
		&frozenRaw,
		&current.UpdatedAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return account.Balance{}, errors.New("balance not found")
		}
		return account.Balance{}, err
	}

	availableAmount, err := decimal.NewFromString(availableRaw)
	if err != nil {
		return account.Balance{}, err
	}
	frozenAmount, err := decimal.NewFromString(frozenRaw)
	if err != nil {
		return account.Balance{}, err
	}

	current.AvailableAmount = availableAmount
	current.FrozenAmount = frozenAmount
	return current, nil
}

func (s *BalanceStore) SaveBalance(ctx context.Context, balance account.Balance) error {
	if s.Database == nil || s.Database.Connection() == nil {
		return errors.New("postgres balance store db is nil")
	}

	_, err := s.Database.Connection().ExecContext(ctx, `
		INSERT INTO accounts (
			user_id,
			asset,
			available_balance,
			frozen_balance,
			updated_at
		) VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (user_id, asset) DO UPDATE SET
			available_balance = EXCLUDED.available_balance,
			frozen_balance = EXCLUDED.frozen_balance,
			updated_at = EXCLUDED.updated_at
	`,
		balance.UserID,
		balance.Asset,
		balance.AvailableAmount.String(),
		balance.FrozenAmount.String(),
		balance.UpdatedAt,
	)
	return err
}
