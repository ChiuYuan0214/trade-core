package postgres

import (
	"context"
	"database/sql"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"local.exchange-demo/exchange-core-go/domain/account"
)

type testConnectionProvider struct {
	conn *sql.DB
}

func (p *testConnectionProvider) Connection() *sql.DB { return p.conn }
func (p *testConnectionProvider) Run() error          { return nil }
func (p *testConnectionProvider) Stop()               {}

func TestBalanceStoreSaveBalance(t *testing.T) {
	conn, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmock.New: %v", err)
	}
	defer conn.Close()

	userID := uuid.New()
	updatedAt := time.Now().UTC()
	store := &BalanceStore{Database: &testConnectionProvider{conn: conn}}

	mock.ExpectExec("INSERT INTO accounts").
		WithArgs(userID, "USDT", "100.5", "12", updatedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = store.SaveBalance(context.Background(), account.Balance{
		UserID:          userID,
		Asset:           "USDT",
		AvailableAmount: decimal.RequireFromString("100.5"),
		FrozenAmount:    decimal.RequireFromString("12"),
		UpdatedAt:       updatedAt,
	})
	if err != nil {
		t.Fatalf("SaveBalance: %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("ExpectationsWereMet: %v", err)
	}
}

func TestBalanceStoreGetBalance(t *testing.T) {
	conn, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmock.New: %v", err)
	}
	defer conn.Close()

	userID := uuid.New()
	updatedAt := time.Now().UTC()
	store := &BalanceStore{Database: &testConnectionProvider{conn: conn}}

	rows := sqlmock.NewRows([]string{
		"user_id", "asset", "available_balance", "frozen_balance", "updated_at",
	}).AddRow(userID, "BTC", "7.5", "1.25", updatedAt)

	mock.ExpectQuery("SELECT\\s+user_id,\\s+asset,\\s+available_balance,\\s+frozen_balance,\\s+updated_at\\s+FROM accounts").
		WithArgs(userID, "BTC").
		WillReturnRows(rows)

	current, err := store.GetBalance(context.Background(), userID, "BTC")
	if err != nil {
		t.Fatalf("GetBalance: %v", err)
	}
	if !current.AvailableAmount.Equal(decimal.RequireFromString("7.5")) {
		t.Fatalf("available mismatch: %s", current.AvailableAmount)
	}
	if !current.FrozenAmount.Equal(decimal.RequireFromString("1.25")) {
		t.Fatalf("frozen mismatch: %s", current.FrozenAmount)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("ExpectationsWereMet: %v", err)
	}
}
