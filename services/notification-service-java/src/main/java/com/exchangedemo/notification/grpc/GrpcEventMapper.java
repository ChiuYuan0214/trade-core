package com.exchangedemo.notification.grpc;

import com.exchangedemo.notification.model.EventEnvelope;
import com.exchangedemo.notification.proto.PrivateEvent;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.google.protobuf.InvalidProtocolBufferException;
import com.google.protobuf.util.JsonFormat;
import io.grpc.Status;
import java.io.IOException;
import java.time.Instant;
import java.util.UUID;

public final class GrpcEventMapper {

    private static final ObjectMapper objectMapper = new ObjectMapper();
    private static final JsonFormat.Printer printer = JsonFormat.printer();

    private GrpcEventMapper() {
    }

    public static EventEnvelope toEnvelope(PrivateEvent event) {
        return new EventEnvelope(
                parseUuid(event.getEventId()),
                event.getEventType(),
                toInstant(event),
                parseUuid(event.getCorrelationId()),
                parseUuid(event.getCausationId()),
                emptyToNull(event.getSymbol()),
                parseUuid(event.getOrderId()),
                parseUuid(event.getUserId()),
                emptyToNull(event.getShardId()),
                event.getVersion(),
                toJsonNode(event)
        );
    }

    private static Instant toInstant(PrivateEvent event) {
        if (!event.hasOccurredAt()) {
            return null;
        }
        return Instant.ofEpochSecond(event.getOccurredAt().getSeconds(), event.getOccurredAt().getNanos());
    }

    private static JsonNode toJsonNode(PrivateEvent event) {
        if (!event.hasPayload()) {
            return objectMapper.createObjectNode();
        }
        try {
            return objectMapper.readTree(printer.print(event.getPayload()));
        } catch (InvalidProtocolBufferException e) {
            throw Status.INVALID_ARGUMENT.withDescription("invalid protobuf payload").withCause(e).asRuntimeException();
        } catch (IOException e) {
            throw Status.INVALID_ARGUMENT.withDescription("unable to parse payload as JSON").withCause(e).asRuntimeException();
        }
    }

    private static UUID parseUuid(String raw) {
        if (raw == null || raw.isBlank()) {
            return null;
        }
        try {
            return UUID.fromString(raw);
        } catch (IllegalArgumentException error) {
            throw Status.INVALID_ARGUMENT.withDescription("invalid uuid: " + raw).withCause(error).asRuntimeException();
        }
    }

    private static String emptyToNull(String value) {
        return value == null || value.isBlank() ? null : value;
    }
}
