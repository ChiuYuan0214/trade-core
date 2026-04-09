package com.exchangedemo.notification.web;

import com.exchangedemo.notification.model.EventEnvelope;
import com.exchangedemo.notification.model.UserNotification;
import com.exchangedemo.notification.service.NotificationDispatchService;
import jakarta.validation.Valid;
import java.util.Map;
import java.util.Optional;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/internal/notifications")
public class NotificationIngestController {

    private final NotificationDispatchService dispatchService;

    public NotificationIngestController(NotificationDispatchService dispatchService) {
        this.dispatchService = dispatchService;
    }

    @PostMapping("/events")
    public ResponseEntity<?> ingest(@Valid @RequestBody EventEnvelope envelope) {
        Optional<UserNotification> notification = dispatchService.dispatch(envelope);
        if (notification.isEmpty()) {
            return ResponseEntity.accepted().body(Map.of("status", "ignored"));
        }
        return ResponseEntity.accepted().body(notification.get());
    }
}
