package postgres

import (
	"context"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"local.exchange-demo/exchange-core-go/domain/ledger"
)

func TestLedgerStoreAppendEntries(t *testing.T) {
	conn, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmock.New: %v", err)
	}
	defer conn.Close()

	store := &LedgerStore{Database: &testConnectionProvider{conn: conn}}
	entry := ledger.Entry{
		ID:             uuid.New(),
		UserID:         uuid.New(),
		Asset:          "USDT",
		DeltaAvailable: decimal.RequireFromString("-10"),
		DeltaFrozen:    decimal.RequireFromString("10"),
		ReferenceType:  ledger.ReferenceTypeOrderReservation,
		ReferenceID:    uuid.New(),
		EventID:        uuid.New(),
		CreatedAt:      time.Now().UTC(),
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO ledger_entries").
		WithArgs(
			entry.ID,
			entry.UserID,
			entry.Asset,
			entry.DeltaAvailable.String(),
			entry.DeltaFrozen.String(),
			entry.ReferenceType,
			entry.ReferenceID,
			entry.EventID,
			entry.CreatedAt,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	if err := store.AppendEntries(context.Background(), entry); err != nil {
		t.Fatalf("AppendEntries: %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("ExpectationsWereMet: %v", err)
	}
}
