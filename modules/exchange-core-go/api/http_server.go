package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"local.exchange-demo/exchange-core-go/app"
	"local.exchange-demo/exchange-core-go/domain/order"
)

type HandlerProvider interface {
	Handler() http.Handler
	Run() error
	Stop()
}

type HTTPServer struct {
	GatewayApplication GatewayApplication
}

func (s *HTTPServer) Run() error { return nil }
func (s *HTTPServer) Stop()      {}

func (s *HTTPServer) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/v1/orders", s.handlePlaceOrder)
	mux.HandleFunc("DELETE /api/v1/orders/{order_id}", s.handleCancelOrder)
	mux.HandleFunc("GET /api/v1/orders/{order_id}", s.handleGetOrder)
	mux.HandleFunc("GET /api/v1/accounts/balances", s.handleGetBalance)
	return mux
}

type placeOrderRequest struct {
	ClientOrderID string `json:"client_order_id"`
	UserID        string `json:"user_id"`
	Symbol        string `json:"symbol"`
	Side          string `json:"side"`
	Type          string `json:"type"`
	Price         string `json:"price"`
	Quantity      string `json:"quantity"`
}

type orderResponse struct {
	OrderID         string `json:"order_id"`
	ClientOrderID   string `json:"client_order_id,omitempty"`
	UserID          string `json:"user_id"`
	Symbol          string `json:"symbol"`
	Side            string `json:"side"`
	Type            string `json:"type"`
	Price           string `json:"price"`
	Quantity        string `json:"quantity"`
	FilledQuantity  string `json:"filled_quantity"`
	Status          string `json:"status"`
	RejectionReason string `json:"rejection_reason,omitempty"`
}

type balanceResponse struct {
	UserID    string `json:"user_id"`
	Asset     string `json:"asset"`
	Available string `json:"available"`
	Frozen    string `json:"frozen"`
}

func (s *HTTPServer) handlePlaceOrder(w http.ResponseWriter, r *http.Request) {
	var req placeOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	price, err := parseDecimal(req.Price)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}
	quantity, err := parseDecimal(req.Quantity)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	placed, err := s.GatewayApplication.PlaceOrder(r.Context(), app.PlaceOrderInput{
		ClientOrderID: req.ClientOrderID,
		UserID:        userID,
		Symbol:        order.Symbol(req.Symbol),
		Side:          order.Side(req.Side),
		Type:          order.Type(req.Type),
		Price:         price,
		Quantity:      quantity,
	})
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	writeJSON(w, http.StatusAccepted, toOrderResponse(placed.Order))
}

func (s *HTTPServer) handleCancelOrder(w http.ResponseWriter, r *http.Request) {
	orderID, err := uuid.Parse(r.PathValue("order_id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	canceled, err := s.GatewayApplication.CancelOrder(r.Context(), orderID.String())
	if err != nil {
		if errors.Is(err, app.ErrOrderNotFound) || status.Code(err) == codes.NotFound {
			writeError(w, http.StatusNotFound, err)
			return
		}
		writeError(w, http.StatusBadRequest, err)
		return
	}

	writeJSON(w, http.StatusOK, toOrderResponse(canceled))
}

func (s *HTTPServer) handleGetOrder(w http.ResponseWriter, r *http.Request) {
	orderID, err := uuid.Parse(r.PathValue("order_id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	current, err := s.GatewayApplication.GetOrder(r.Context(), orderID.String())
	if err != nil {
		if errors.Is(err, app.ErrOrderNotFound) || status.Code(err) == codes.NotFound {
			writeError(w, http.StatusNotFound, err)
			return
		}
		writeError(w, http.StatusBadRequest, err)
		return
	}

	writeJSON(w, http.StatusOK, toOrderResponse(current))
}

func (s *HTTPServer) handleGetBalance(w http.ResponseWriter, r *http.Request) {
	userID, err := uuid.Parse(r.URL.Query().Get("user_id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}
	asset := r.URL.Query().Get("asset")
	if asset == "" {
		writeError(w, http.StatusBadRequest, errors.New("asset is required"))
		return
	}

	balance, err := s.GatewayApplication.GetBalance(r.Context(), userID.String(), asset)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	writeJSON(w, http.StatusOK, balanceResponse{
		UserID:    balance.UserID.String(),
		Asset:     balance.Asset,
		Available: balance.AvailableAmount.String(),
		Frozen:    balance.FrozenAmount.String(),
	})
}

func parseDecimal(raw string) (decimal.Decimal, error) {
	if raw == "" {
		return decimal.Zero, nil
	}
	return decimal.NewFromString(raw)
}

func writeJSON(w http.ResponseWriter, statusCode int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(payload)
}

func writeError(w http.ResponseWriter, statusCode int, err error) {
	writeJSON(w, statusCode, map[string]string{"error": err.Error()})
}

func toOrderResponse(current order.Order) orderResponse {
	return orderResponse{
		OrderID:         current.ID.String(),
		ClientOrderID:   current.ClientOrderID,
		UserID:          current.UserID.String(),
		Symbol:          string(current.Symbol),
		Side:            string(current.Side),
		Type:            string(current.Type),
		Price:           current.Price.String(),
		Quantity:        current.Quantity.String(),
		FilledQuantity:  current.FilledQuantity.String(),
		Status:          string(current.Status),
		RejectionReason: current.RejectionReason,
	}
}
