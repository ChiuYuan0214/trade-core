package app

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"

	"local.exchange-demo/exchange-core-go/events"
	notificationv1 "local.exchange-demo/exchange-core-go/gen/proto/exchange/notification/v1"
)

func TestGRPCPrivateEventPublisherPublishesExpectedPayload(t *testing.T) {
	t.Parallel()

	listener := bufconn.Listen(1024 * 1024)
	server := grpc.NewServer()
	recorder := &recordingNotificationServer{}
	notificationv1.RegisterPrivateNotificationServiceServer(server, recorder)
	go func() {
		_ = server.Serve(listener)
	}()
	defer server.Stop()

	publisher := &GRPCPrivateEventPublisher{
		Address: "bufnet",
		DialOptions: []grpc.DialOption{
			grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
				return listener.Dial()
			}),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		},
	}
	if err := publisher.Run(); err != nil {
		t.Fatalf("publisher run: %v", err)
	}
	defer publisher.Stop()

	eventID := uuid.New()
	userID := uuid.New()
	err := publisher.Publish(context.Background(), PrivateEvent{
		EventID:       eventID,
		EventType:     events.TypeTradeExecuted,
		OccurredAt:    time.Date(2026, 4, 9, 0, 0, 0, 0, time.UTC),
		CorrelationID: uuid.New(),
		CausationID:   uuid.New(),
		Symbol:        "BTC/USDT",
		OrderID:       uuid.New(),
		UserID:        userID,
		Version:       1,
		Payload: map[string]any{
			"price": "60000",
		},
	})
	if err != nil {
		t.Fatalf("publish event: %v", err)
	}

	if recorder.lastRequest == nil || recorder.lastRequest.Event == nil {
		t.Fatal("expected grpc request to be recorded")
	}
	if recorder.lastRequest.Event.EventId != eventID.String() {
		t.Fatalf("unexpected event id: %s", recorder.lastRequest.Event.EventId)
	}
	if recorder.lastRequest.Event.UserId != userID.String() {
		t.Fatalf("unexpected user id: %s", recorder.lastRequest.Event.UserId)
	}
	if recorder.lastRequest.Event.EventType != string(events.TypeTradeExecuted) {
		t.Fatalf("unexpected event type: %s", recorder.lastRequest.Event.EventType)
	}
}

type recordingNotificationServer struct {
	notificationv1.UnimplementedPrivateNotificationServiceServer
	lastRequest *notificationv1.PublishPrivateEventRequest
}

func (s *recordingNotificationServer) Publish(_ context.Context, request *notificationv1.PublishPrivateEventRequest) (*notificationv1.PublishPrivateEventResponse, error) {
	s.lastRequest = request
	return &notificationv1.PublishPrivateEventResponse{Accepted: true}, nil
}
