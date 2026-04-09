package com.exchangedemo.notification.service;

import com.exchangedemo.notification.model.UserNotification;

public interface NotificationPublisher {
    void publish(UserNotification notification);
}
