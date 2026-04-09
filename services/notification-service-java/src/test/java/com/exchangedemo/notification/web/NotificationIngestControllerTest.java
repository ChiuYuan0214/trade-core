package com.exchangedemo.notification.web;

import com.exchangedemo.notification.model.EventEnvelope;
import com.exchangedemo.notification.model.NotificationChannel;
import com.exchangedemo.notification.model.UserNotification;
import com.exchangedemo.notification.service.NotificationDispatchService;
import com.fasterxml.jackson.databind.ObjectMapper;
import java.time.Instant;
import java.util.Optional;
import java.util.UUID;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.WebMvcTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.http.MediaType;
import org.springframework.test.web.servlet.MockMvc;

import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.when;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.post;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.jsonPath;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

@WebMvcTest(NotificationIngestController.class)
class NotificationIngestControllerTest {

    @Autowired
    private MockMvc mockMvc;

    @Autowired
    private ObjectMapper objectMapper;

    @MockBean
    private NotificationDispatchService dispatchService;

    @Test
    void ingestReturnsAcceptedNotificationPayload() throws Exception {
        UUID userId = UUID.fromString("11111111-1111-1111-1111-111111111111");
        when(dispatchService.dispatch(any(EventEnvelope.class))).thenReturn(Optional.of(new UserNotification(
                userId,
                NotificationChannel.USER_BALANCES,
                "/topic/private/" + userId + "/balances",
                "FundsReserved",
                UUID.randomUUID(),
                "BTC/USDT",
                Instant.parse("2026-04-10T00:00:00Z"),
                objectMapper.createObjectNode().put("asset", "USDT")
        )));

        EventEnvelope envelope = new EventEnvelope(
                UUID.randomUUID(),
                "FundsReserved",
                Instant.parse("2026-04-10T00:00:00Z"),
                UUID.randomUUID(),
                UUID.randomUUID(),
                "BTC/USDT",
                UUID.randomUUID(),
                userId,
                "shard-1",
                1,
                objectMapper.createObjectNode().put("amount", "60000")
        );

        mockMvc.perform(post("/internal/notifications/events")
                        .contentType(MediaType.APPLICATION_JSON)
                        .content(objectMapper.writeValueAsBytes(envelope)))
                .andExpect(status().isAccepted())
                .andExpect(jsonPath("$.channel").value("USER_BALANCES"))
                .andExpect(jsonPath("$.topic").value("/topic/private/" + userId + "/balances"))
                .andExpect(jsonPath("$.eventType").value("FundsReserved"));
    }

    @Test
    void ingestReturnsIgnoredWhenRouterDropsEvent() throws Exception {
        when(dispatchService.dispatch(any(EventEnvelope.class))).thenReturn(Optional.empty());

        EventEnvelope envelope = new EventEnvelope(
                UUID.randomUUID(),
                "IgnoredEvent",
                Instant.parse("2026-04-10T00:00:00Z"),
                UUID.randomUUID(),
                UUID.randomUUID(),
                "BTC/USDT",
                UUID.randomUUID(),
                UUID.fromString("11111111-1111-1111-1111-111111111111"),
                "shard-1",
                1,
                objectMapper.createObjectNode()
        );

        mockMvc.perform(post("/internal/notifications/events")
                        .contentType(MediaType.APPLICATION_JSON)
                        .content(objectMapper.writeValueAsBytes(envelope)))
                .andExpect(status().isAccepted())
                .andExpect(jsonPath("$.status").value("ignored"));
    }
}
