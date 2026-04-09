package order

import (
	"testing"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func TestValidateCreateInput(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name      string
		input     CreateInput
		expectErr bool
	}{
		{
			name: "valid limit order",
			input: CreateInput{
				UserID:   uuid.New(),
				Symbol:   SymbolBTCUSDT,
				Side:     SideBuy,
				Type:     TypeLimit,
				Price:    decimal.RequireFromString("60000"),
				Quantity: decimal.RequireFromString("0.5"),
			},
		},
		{
			name: "reject unsupported symbol",
			input: CreateInput{
				UserID:   uuid.New(),
				Symbol:   Symbol("DOGE/USDT"),
				Side:     SideBuy,
				Type:     TypeLimit,
				Price:    decimal.RequireFromString("1"),
				Quantity: decimal.RequireFromString("1"),
			},
			expectErr: true,
		},
		{
			name: "reject zero quantity",
			input: CreateInput{
				UserID:   uuid.New(),
				Symbol:   SymbolBTCUSDT,
				Side:     SideBuy,
				Type:     TypeLimit,
				Price:    decimal.RequireFromString("1"),
				Quantity: decimal.Zero,
			},
			expectErr: true,
		},
		{
			name: "reject limit order without price",
			input: CreateInput{
				UserID:   uuid.New(),
				Symbol:   SymbolBTCUSDT,
				Side:     SideBuy,
				Type:     TypeLimit,
				Price:    decimal.Zero,
				Quantity: decimal.RequireFromString("1"),
			},
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			err := ValidateCreateInput(tc.input)
			if tc.expectErr && err == nil {
				t.Fatalf("expected error but got nil")
			}
			if !tc.expectErr && err != nil {
				t.Fatalf("expected nil error but got %v", err)
			}
		})
	}
}
