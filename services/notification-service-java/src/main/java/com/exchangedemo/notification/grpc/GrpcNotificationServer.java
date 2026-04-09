package com.exchangedemo.notification.grpc;

import com.exchangedemo.notification.config.NotificationProperties;
import io.grpc.Server;
import io.grpc.ServerBuilder;
import java.io.IOException;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.TimeUnit;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.context.SmartLifecycle;
import org.springframework.stereotype.Component;

@Component
public class GrpcNotificationServer implements SmartLifecycle {

    private static final Logger logger = LoggerFactory.getLogger(GrpcNotificationServer.class);

    private final NotificationProperties properties;
    private final GrpcNotificationService service;
    private final ExecutorService notificationVirtualThreadExecutor;

    private volatile boolean running;
    private Server server;

    public GrpcNotificationServer(
            NotificationProperties properties,
            GrpcNotificationService service,
            ExecutorService notificationVirtualThreadExecutor
    ) {
        this.properties = properties;
        this.service = service;
        this.notificationVirtualThreadExecutor = notificationVirtualThreadExecutor;
    }

    @Override
    public void start() {
        if (running) {
            return;
        }
        try {
            server = ServerBuilder.forPort(properties.grpcPort())
                    .executor(notificationVirtualThreadExecutor)
                    .addService(service)
                    .build()
                    .start();
            running = true;
            logger.info("notification gRPC server listening on port {} with virtual threads", properties.grpcPort());
        } catch (IOException error) {
            throw new IllegalStateException("failed to start notification gRPC server", error);
        }
    }

    @Override
    public void stop() {
        if (!running || server == null) {
            return;
        }
        server.shutdown();
        try {
            if (!server.awaitTermination(5, TimeUnit.SECONDS)) {
                server.shutdownNow();
            }
        } catch (InterruptedException error) {
            Thread.currentThread().interrupt();
            server.shutdownNow();
        } finally {
            running = false;
        }
    }

    @Override
    public boolean isRunning() {
        return running;
    }

    @Override
    public boolean isAutoStartup() {
        return true;
    }

    @Override
    public int getPhase() {
        return Integer.MAX_VALUE;
    }
}
