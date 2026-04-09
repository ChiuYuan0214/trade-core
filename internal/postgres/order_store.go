package postgres

import (
	"context"
	"errors"
	"database/sql"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"exchange-demo/internal/app"
	"exchange-demo/internal/domain/order"
)

type OrderStore struct {
	Database ConnectionProvider
}

func (s *OrderStore) Run() error { return nil }
func (s *OrderStore) Stop()      {}

func (s *OrderStore) Save(ctx context.Context, current order.Order) error {
	if s.Database == nil || s.Database.Connection() == nil {
		return errors.New("postgres order store db is nil")
	}

	_, err := s.Database.Connection().ExecContext(ctx, `
		INSERT INTO orders (
			order_id,
			client_order_id,
			user_id,
			symbol,
			side,
			type,
			price,
			quantity,
			filled_quantity,
			status,
			rejection_reason,
			created_at,
			updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
		)
		ON CONFLICT (order_id) DO UPDATE SET
			client_order_id = EXCLUDED.client_order_id,
			user_id = EXCLUDED.user_id,
			symbol = EXCLUDED.symbol,
			side = EXCLUDED.side,
			type = EXCLUDED.type,
			price = EXCLUDED.price,
			quantity = EXCLUDED.quantity,
			filled_quantity = EXCLUDED.filled_quantity,
			status = EXCLUDED.status,
			rejection_reason = EXCLUDED.rejection_reason,
			created_at = EXCLUDED.created_at,
			updated_at = EXCLUDED.updated_at
	`,
		current.ID,
		nullIfEmpty(current.ClientOrderID),
		current.UserID,
		current.Symbol,
		current.Side,
		current.Type,
		current.Price.String(),
		current.Quantity.String(),
		current.FilledQuantity.String(),
		current.Status,
		nullIfEmpty(current.RejectionReason),
		current.CreatedAt,
		current.UpdatedAt,
	)
	return err
}

func (s *OrderStore) Get(ctx context.Context, orderID uuid.UUID) (order.Order, error) {
	if s.Database == nil || s.Database.Connection() == nil {
		return order.Order{}, errors.New("postgres order store db is nil")
	}

	row := s.Database.Connection().QueryRowContext(ctx, `
		SELECT
			order_id,
			client_order_id,
			user_id,
			symbol,
			side,
			type,
			price,
			quantity,
			filled_quantity,
			status,
			rejection_reason,
			created_at,
			updated_at
		FROM orders
		WHERE order_id = $1
	`, orderID)

	var current order.Order
	var clientOrderID sql.NullString
	var price sql.NullString
	var rejectionReason sql.NullString
	var symbol string
	var side string
	var orderType string
	var status string

	err := row.Scan(
		&current.ID,
		&clientOrderID,
		&current.UserID,
		&symbol,
		&side,
		&orderType,
		&price,
		&current.Quantity,
		&current.FilledQuantity,
		&status,
		&rejectionReason,
		&current.CreatedAt,
		&current.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return order.Order{}, app.ErrOrderNotFound
		}
		return order.Order{}, err
	}

	current.ClientOrderID = clientOrderID.String
	current.Symbol = order.Symbol(symbol)
	current.Side = order.Side(side)
	current.Type = order.Type(orderType)
	current.Status = order.Status(status)
	current.RejectionReason = rejectionReason.String
	if price.Valid {
		parsedPrice, err := decimal.NewFromString(price.String)
		if err != nil {
			return order.Order{}, err
		}
		current.Price = parsedPrice
	}

	return current, nil
}

func nullIfEmpty(value string) any {
	if value == "" {
		return nil
	}
	return value
}
