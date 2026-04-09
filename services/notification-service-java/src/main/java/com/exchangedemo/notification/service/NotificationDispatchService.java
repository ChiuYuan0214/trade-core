package com.exchangedemo.notification.service;

import com.exchangedemo.notification.model.EventEnvelope;
import com.exchangedemo.notification.model.UserNotification;
import java.util.Optional;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Service;

@Service
public class NotificationDispatchService {

    private static final Logger logger = LoggerFactory.getLogger(NotificationDispatchService.class);

    private final NotificationRouter router;
    private final NotificationPublisher publisher;

    public NotificationDispatchService(NotificationRouter router, NotificationPublisher publisher) {
        this.router = router;
        this.publisher = publisher;
    }

    public Optional<UserNotification> dispatch(EventEnvelope envelope) {
        Optional<UserNotification> routed = router.route(envelope);
        routed.ifPresent(notification -> {
            publisher.publish(notification);
            logger.info("published private notification channel={} userId={} eventType={}",
                    notification.channel(), notification.userId(), notification.eventType());
        });
        return routed;
    }
}
