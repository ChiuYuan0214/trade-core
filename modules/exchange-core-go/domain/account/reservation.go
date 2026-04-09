package account

import (
	"fmt"

	"github.com/shopspring/decimal"

	"local.exchange-demo/exchange-core-go/domain/order"
)

func ReservationAsset(current order.Order) (string, error) {
	switch current.Side {
	case order.SideBuy:
		return QuoteAsset(current.Symbol)
	case order.SideSell:
		return BaseAsset(current.Symbol)
	default:
		return "", fmt.Errorf("unsupported order side %s", current.Side)
	}
}

func ReservationAmount(current order.Order) (decimal.Decimal, error) {
	switch current.Side {
	case order.SideBuy:
		if current.Type == order.TypeMarket {
			return decimal.Zero, fmt.Errorf("market buy reservation amount requires a pricing policy")
		}
		return current.Price.Mul(current.Quantity), nil
	case order.SideSell:
		return current.Quantity, nil
	default:
		return decimal.Zero, fmt.Errorf("unsupported order side %s", current.Side)
	}
}

func BaseAsset(symbol order.Symbol) (string, error) {
	switch symbol {
	case order.SymbolBTCUSDT:
		return "BTC", nil
	case order.SymbolETHUSDT:
		return "ETH", nil
	case order.SymbolSOLUSDT:
		return "SOL", nil
	default:
		return "", fmt.Errorf("unsupported symbol %s", symbol)
	}
}

func QuoteAsset(symbol order.Symbol) (string, error) {
	switch symbol {
	case order.SymbolBTCUSDT, order.SymbolETHUSDT, order.SymbolSOLUSDT:
		return "USDT", nil
	default:
		return "", fmt.Errorf("unsupported symbol %s", symbol)
	}
}
