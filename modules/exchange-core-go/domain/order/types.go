package order

import "fmt"

type Symbol string

const (
	SymbolBTCUSDT Symbol = "BTC/USDT"
	SymbolETHUSDT Symbol = "ETH/USDT"
	SymbolSOLUSDT Symbol = "SOL/USDT"
)

func (s Symbol) Validate() error {
	switch s {
	case SymbolBTCUSDT, SymbolETHUSDT, SymbolSOLUSDT:
		return nil
	default:
		return fmt.Errorf("unsupported symbol %q", s)
	}
}

func SupportedSymbols() []Symbol {
	return []Symbol{SymbolBTCUSDT, SymbolETHUSDT, SymbolSOLUSDT}
}

type Side string

const (
	SideBuy  Side = "BUY"
	SideSell Side = "SELL"
)

func (s Side) Validate() error {
	switch s {
	case SideBuy, SideSell:
		return nil
	default:
		return fmt.Errorf("unsupported side %q", s)
	}
}

type Type string

const (
	TypeLimit  Type = "LIMIT"
	TypeMarket Type = "MARKET"
)

func (t Type) Validate() error {
	switch t {
	case TypeLimit, TypeMarket:
		return nil
	default:
		return fmt.Errorf("unsupported order type %q", t)
	}
}

type Status string

const (
	StatusPendingAccept   Status = "PENDING_ACCEPT"
	StatusOpen            Status = "OPEN"
	StatusPartiallyFilled Status = "PARTIALLY_FILLED"
	StatusFilled          Status = "FILLED"
	StatusCanceled        Status = "CANCELED"
	StatusRejected        Status = "REJECTED"
)
