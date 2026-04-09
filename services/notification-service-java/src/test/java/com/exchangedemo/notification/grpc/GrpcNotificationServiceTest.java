package com.exchangedemo.notification.grpc;

import com.exchangedemo.notification.model.EventEnvelope;
import com.exchangedemo.notification.model.NotificationChannel;
import com.exchangedemo.notification.model.UserNotification;
import com.exchangedemo.notification.proto.PrivateEvent;
import com.exchangedemo.notification.proto.PublishPrivateEventRequest;
import com.exchangedemo.notification.proto.PublishPrivateEventResponse;
import com.exchangedemo.notification.service.NotificationDispatchService;
import com.google.protobuf.Struct;
import com.google.protobuf.Timestamp;
import io.grpc.stub.StreamObserver;
import java.time.Instant;
import java.util.Map;
import java.util.Optional;
import java.util.UUID;
import org.junit.jupiter.api.Test;
import org.mockito.ArgumentCaptor;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertFalse;
import static org.junit.jupiter.api.Assertions.assertTrue;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;

class GrpcNotificationServiceTest {

    @Test
    void publishReturnsAcceptedNotificationMetadata() {
        NotificationDispatchService dispatchService = mock(NotificationDispatchService.class);
        GrpcNotificationService service = new GrpcNotificationService(dispatchService);
        RecordingObserver observer = new RecordingObserver();

        UUID eventId = UUID.randomUUID();
        UUID correlationId = UUID.randomUUID();
        UUID causationId = UUID.randomUUID();
        UUID orderId = UUID.randomUUID();
        UUID userId = UUID.fromString("11111111-1111-1111-1111-111111111111");
        Instant occurredAt = Instant.parse("2026-04-10T00:00:00Z");

        when(dispatchService.dispatch(any(EventEnvelope.class))).thenReturn(Optional.of(new UserNotification(
                userId,
                NotificationChannel.USER_TRADES,
                "/topic/private/" + userId + "/trades",
                "TradeExecuted",
                orderId,
                "BTC/USDT",
                occurredAt,
                Map.of("tradeId", UUID.randomUUID().toString())
        )));

        PublishPrivateEventRequest request = PublishPrivateEventRequest.newBuilder()
                .setEvent(PrivateEvent.newBuilder()
                        .setEventId(eventId.toString())
                        .setEventType("TradeExecuted")
                        .setOccurredAt(Timestamp.newBuilder()
                                .setSeconds(occurredAt.getEpochSecond())
                                .setNanos(occurredAt.getNano())
                                .build())
                        .setCorrelationId(correlationId.toString())
                        .setCausationId(causationId.toString())
                        .setSymbol("BTC/USDT")
                        .setOrderId(orderId.toString())
                        .setUserId(userId.toString())
                        .setShardId("shard-1")
                        .setVersion(1)
                        .setPayload(Struct.getDefaultInstance())
                        .build())
                .build();

        service.publish(request, observer);

        ArgumentCaptor<EventEnvelope> captor = ArgumentCaptor.forClass(EventEnvelope.class);
        verify(dispatchService).dispatch(captor.capture());
        assertEquals("TradeExecuted", captor.getValue().eventType());
        assertEquals(userId, captor.getValue().userId());

        assertTrue(observer.completed);
        assertFalse(observer.errored);
        assertEquals(1, observer.responses);
        assertTrue(observer.lastResponse.accepted());
        assertFalse(observer.lastResponse.ignored());
        assertEquals("/topic/private/" + userId + "/trades", observer.lastResponse.topic());
        assertEquals("USER_TRADES", observer.lastResponse.channel());
    }

    private static final class RecordingObserver implements StreamObserver<PublishPrivateEventResponse> {
        private PublishPrivateEventResponseView lastResponse;
        private int responses;
        private boolean completed;
        private boolean errored;

        @Override
        public void onNext(PublishPrivateEventResponse value) {
            responses++;
            lastResponse = new PublishPrivateEventResponseView(
                    value.getAccepted(),
                    value.getIgnored(),
                    value.getTopic(),
                    value.getChannel()
            );
        }

        @Override
        public void onError(Throwable throwable) {
            errored = true;
        }

        @Override
        public void onCompleted() {
            completed = true;
        }
    }

    private record PublishPrivateEventResponseView(
            boolean accepted,
            boolean ignored,
            String topic,
            String channel
    ) {
    }
}
