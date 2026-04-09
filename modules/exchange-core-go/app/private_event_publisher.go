package app

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"local.exchange-demo/exchange-core-go/events"
	notificationv1 "local.exchange-demo/exchange-core-go/gen/proto/exchange/notification/v1"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type PrivateEvent struct {
	EventID       uuid.UUID   `json:"eventId"`
	EventType     events.Type `json:"eventType"`
	OccurredAt    time.Time   `json:"occurredAt"`
	CorrelationID uuid.UUID   `json:"correlationId"`
	CausationID   uuid.UUID   `json:"causationId"`
	Symbol        string      `json:"symbol,omitempty"`
	OrderID       uuid.UUID   `json:"orderId,omitempty"`
	UserID        uuid.UUID   `json:"userId,omitempty"`
	ShardID       string      `json:"shardId,omitempty"`
	Version       int         `json:"version,omitempty"`
	Payload       any         `json:"payload"`
}

type PrivateEventPublisher interface {
	Publish(ctx context.Context, event PrivateEvent) error
	Run() error
	Stop()
}

type NoopPrivateEventPublisher struct{}

func (p *NoopPrivateEventPublisher) Publish(_ context.Context, _ PrivateEvent) error { return nil }
func (p *NoopPrivateEventPublisher) Run() error                                      { return nil }
func (p *NoopPrivateEventPublisher) Stop()                                           {}

type GRPCPrivateEventPublisher struct {
	Address     string
	Timeout     time.Duration
	DialOptions []grpc.DialOption

	closeFn   func() error
	publishFn func(context.Context, *notificationv1.PublishPrivateEventRequest, ...grpc.CallOption) (*notificationv1.PublishPrivateEventResponse, error)
}

func (p *GRPCPrivateEventPublisher) Run() error {
	if p.Address == "" {
		return fmt.Errorf("notification grpc address is required")
	}
	if p.Timeout <= 0 {
		p.Timeout = 3 * time.Second
	}

	options := p.DialOptions
	if len(options) == 0 {
		options = []grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithBlock(),
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), p.Timeout)
	defer cancel()

	conn, err := grpc.DialContext(ctx, p.Address, options...)
	if err != nil {
		return fmt.Errorf("dial notification grpc service: %w", err)
	}

	client := notificationv1.NewPrivateNotificationServiceClient(conn)
	p.closeFn = conn.Close
	p.publishFn = client.Publish
	return nil
}

func (p *GRPCPrivateEventPublisher) Stop() {
	if p.closeFn != nil {
		_ = p.closeFn()
	}
}

func (p *GRPCPrivateEventPublisher) Publish(ctx context.Context, event PrivateEvent) error {
	if p.publishFn == nil {
		return fmt.Errorf("notification grpc client is not initialized")
	}

	request, err := privateEventRequest(event)
	if err != nil {
		return err
	}

	resp, err := p.publishFn(ctx, request)
	if err != nil {
		return fmt.Errorf("publish private event over grpc: %w", err)
	}
	if !resp.GetAccepted() {
		return fmt.Errorf("notification grpc service rejected private event")
	}
	return nil
}

func privateEventRequest(event PrivateEvent) (*notificationv1.PublishPrivateEventRequest, error) {
	payload, err := privateEventPayload(event.Payload)
	if err != nil {
		return nil, err
	}

	return &notificationv1.PublishPrivateEventRequest{
		Event: &notificationv1.PrivateEvent{
			EventId:       uuidString(event.EventID),
			EventType:     string(event.EventType),
			OccurredAt:    timestamppb.New(event.OccurredAt),
			CorrelationId: uuidString(event.CorrelationID),
			CausationId:   uuidString(event.CausationID),
			Symbol:        event.Symbol,
			OrderId:       uuidString(event.OrderID),
			UserId:        uuidString(event.UserID),
			ShardId:       event.ShardID,
			Version:       int32(event.Version),
			Payload:       payload,
		},
	}, nil
}

func privateEventPayload(payload any) (*structpb.Struct, error) {
	if payload == nil {
		return structpb.NewStruct(map[string]any{})
	}

	switch value := payload.(type) {
	case map[string]any:
		return structpb.NewStruct(value)
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("marshal private event payload: %w", err)
	}

	var object map[string]any
	if err := json.Unmarshal(body, &object); err == nil && object != nil {
		return structpb.NewStruct(object)
	}

	var scalar any
	if err := json.Unmarshal(body, &scalar); err != nil {
		return nil, fmt.Errorf("unmarshal private event payload: %w", err)
	}

	return structpb.NewStruct(map[string]any{"value": scalar})
}

func uuidString(id uuid.UUID) string {
	if id == uuid.Nil {
		return ""
	}
	return id.String()
}
