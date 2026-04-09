package com.exchangedemo.notification.grpc;

import com.exchangedemo.notification.model.UserNotification;
import com.exchangedemo.notification.proto.PrivateNotificationServiceGrpc;
import com.exchangedemo.notification.proto.PublishPrivateEventRequest;
import com.exchangedemo.notification.proto.PublishPrivateEventResponse;
import com.exchangedemo.notification.service.NotificationDispatchService;
import io.grpc.stub.StreamObserver;
import java.util.Optional;
import org.springframework.stereotype.Component;

@Component
public class GrpcNotificationService extends PrivateNotificationServiceGrpc.PrivateNotificationServiceImplBase {

    private final NotificationDispatchService dispatchService;

    public GrpcNotificationService(NotificationDispatchService dispatchService) {
        this.dispatchService = dispatchService;
    }

    @Override
    public void publish(PublishPrivateEventRequest request, StreamObserver<PublishPrivateEventResponse> responseObserver) {
        Optional<UserNotification> notification = dispatchService.dispatch(GrpcEventMapper.toEnvelope(request.getEvent()));

        PublishPrivateEventResponse.Builder response = PublishPrivateEventResponse.newBuilder()
                .setAccepted(true)
                .setIgnored(notification.isEmpty());

        notification.ifPresent(value -> response
                .setTopic(value.topic())
                .setChannel(value.channel().name()));

        responseObserver.onNext(response.build());
        responseObserver.onCompleted();
    }
}
