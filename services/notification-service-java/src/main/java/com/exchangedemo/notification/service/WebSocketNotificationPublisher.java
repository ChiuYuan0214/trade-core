package com.exchangedemo.notification.service;

import com.exchangedemo.notification.model.UserNotification;
import org.springframework.messaging.simp.SimpMessagingTemplate;
import org.springframework.stereotype.Component;

@Component
public class WebSocketNotificationPublisher implements NotificationPublisher {

    private final SimpMessagingTemplate messagingTemplate;

    public WebSocketNotificationPublisher(SimpMessagingTemplate messagingTemplate) {
        this.messagingTemplate = messagingTemplate;
    }

    @Override
    public void publish(UserNotification notification) {
        messagingTemplate.convertAndSend(notification.topic(), notification);
    }
}
