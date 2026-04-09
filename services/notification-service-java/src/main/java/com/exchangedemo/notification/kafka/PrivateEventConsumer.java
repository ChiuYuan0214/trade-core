package com.exchangedemo.notification.kafka;

import com.exchangedemo.notification.model.EventEnvelope;
import com.exchangedemo.notification.service.NotificationDispatchService;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.stereotype.Component;

@Component
public class PrivateEventConsumer {

    private final NotificationDispatchService dispatchService;

    public PrivateEventConsumer(NotificationDispatchService dispatchService) {
        this.dispatchService = dispatchService;
    }

    @KafkaListener(
            topics = "${exchange.notification.event-topic}",
            groupId = "${spring.kafka.consumer.group-id}"
    )
    public void consume(EventEnvelope envelope) {
        dispatchService.dispatch(envelope);
    }
}
