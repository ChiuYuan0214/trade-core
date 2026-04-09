package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type DemoLedgerSeed struct {
	Database ConnectionProvider
}

func (s *DemoLedgerSeed) Run() error {
	if s.Database == nil || s.Database.Connection() == nil {
		return fmt.Errorf("demo ledger seed db is nil")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	now := time.Now().UTC()
	for _, userID := range []string{
		"11111111-1111-1111-1111-111111111111",
		"22222222-2222-2222-2222-222222222222",
	} {
		parsedUserID, err := uuid.Parse(userID)
		if err != nil {
			return err
		}
		if _, err := s.Database.Connection().ExecContext(ctx, `
			INSERT INTO users (user_id, created_at)
			VALUES ($1, $2)
			ON CONFLICT (user_id) DO NOTHING
		`, parsedUserID, now); err != nil {
			return err
		}
		for _, assetBalance := range []struct {
			asset     string
			available string
			frozen    string
		}{
			{asset: "USDT", available: "1000000", frozen: "0"},
			{asset: "BTC", available: "100", frozen: "0"},
			{asset: "ETH", available: "1000", frozen: "0"},
			{asset: "SOL", available: "10000", frozen: "0"},
		} {
			if _, err := s.Database.Connection().ExecContext(ctx, `
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
				parsedUserID,
				assetBalance.asset,
				assetBalance.available,
				assetBalance.frozen,
				now,
			); err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *DemoLedgerSeed) Stop() {}
