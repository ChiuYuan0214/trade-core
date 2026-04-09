package com.exchangedemo.notification.proto;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@io.grpc.stub.annotations.GrpcGenerated
public final class PrivateNotificationServiceGrpc {

  private PrivateNotificationServiceGrpc() {}

  public static final java.lang.String SERVICE_NAME = "exchange.notification.v1.PrivateNotificationService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.exchangedemo.notification.proto.PublishPrivateEventRequest,
      com.exchangedemo.notification.proto.PublishPrivateEventResponse> getPublishMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Publish",
      requestType = com.exchangedemo.notification.proto.PublishPrivateEventRequest.class,
      responseType = com.exchangedemo.notification.proto.PublishPrivateEventResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.exchangedemo.notification.proto.PublishPrivateEventRequest,
      com.exchangedemo.notification.proto.PublishPrivateEventResponse> getPublishMethod() {
    io.grpc.MethodDescriptor<com.exchangedemo.notification.proto.PublishPrivateEventRequest, com.exchangedemo.notification.proto.PublishPrivateEventResponse> getPublishMethod;
    if ((getPublishMethod = PrivateNotificationServiceGrpc.getPublishMethod) == null) {
      synchronized (PrivateNotificationServiceGrpc.class) {
        if ((getPublishMethod = PrivateNotificationServiceGrpc.getPublishMethod) == null) {
          PrivateNotificationServiceGrpc.getPublishMethod = getPublishMethod =
              io.grpc.MethodDescriptor.<com.exchangedemo.notification.proto.PublishPrivateEventRequest, com.exchangedemo.notification.proto.PublishPrivateEventResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Publish"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.exchangedemo.notification.proto.PublishPrivateEventRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.exchangedemo.notification.proto.PublishPrivateEventResponse.getDefaultInstance()))
              .setSchemaDescriptor(new PrivateNotificationServiceMethodDescriptorSupplier("Publish"))
              .build();
        }
      }
    }
    return getPublishMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static PrivateNotificationServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<PrivateNotificationServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<PrivateNotificationServiceStub>() {
        @java.lang.Override
        public PrivateNotificationServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new PrivateNotificationServiceStub(channel, callOptions);
        }
      };
    return PrivateNotificationServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports all types of calls on the service
   */
  public static PrivateNotificationServiceBlockingV2Stub newBlockingV2Stub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<PrivateNotificationServiceBlockingV2Stub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<PrivateNotificationServiceBlockingV2Stub>() {
        @java.lang.Override
        public PrivateNotificationServiceBlockingV2Stub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new PrivateNotificationServiceBlockingV2Stub(channel, callOptions);
        }
      };
    return PrivateNotificationServiceBlockingV2Stub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static PrivateNotificationServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<PrivateNotificationServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<PrivateNotificationServiceBlockingStub>() {
        @java.lang.Override
        public PrivateNotificationServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new PrivateNotificationServiceBlockingStub(channel, callOptions);
        }
      };
    return PrivateNotificationServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static PrivateNotificationServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<PrivateNotificationServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<PrivateNotificationServiceFutureStub>() {
        @java.lang.Override
        public PrivateNotificationServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new PrivateNotificationServiceFutureStub(channel, callOptions);
        }
      };
    return PrivateNotificationServiceFutureStub.newStub(factory, channel);
  }

  /**
   */
  public interface AsyncService {

    /**
     */
    default void publish(com.exchangedemo.notification.proto.PublishPrivateEventRequest request,
        io.grpc.stub.StreamObserver<com.exchangedemo.notification.proto.PublishPrivateEventResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getPublishMethod(), responseObserver);
    }
  }

  /**
   * Base class for the server implementation of the service PrivateNotificationService.
   */
  public static abstract class PrivateNotificationServiceImplBase
      implements io.grpc.BindableService, AsyncService {

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return PrivateNotificationServiceGrpc.bindService(this);
    }
  }

  /**
   * A stub to allow clients to do asynchronous rpc calls to service PrivateNotificationService.
   */
  public static final class PrivateNotificationServiceStub
      extends io.grpc.stub.AbstractAsyncStub<PrivateNotificationServiceStub> {
    private PrivateNotificationServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected PrivateNotificationServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new PrivateNotificationServiceStub(channel, callOptions);
    }

    /**
     */
    public void publish(com.exchangedemo.notification.proto.PublishPrivateEventRequest request,
        io.grpc.stub.StreamObserver<com.exchangedemo.notification.proto.PublishPrivateEventResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getPublishMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * A stub to allow clients to do synchronous rpc calls to service PrivateNotificationService.
   */
  public static final class PrivateNotificationServiceBlockingV2Stub
      extends io.grpc.stub.AbstractBlockingStub<PrivateNotificationServiceBlockingV2Stub> {
    private PrivateNotificationServiceBlockingV2Stub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected PrivateNotificationServiceBlockingV2Stub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new PrivateNotificationServiceBlockingV2Stub(channel, callOptions);
    }

    /**
     */
    public com.exchangedemo.notification.proto.PublishPrivateEventResponse publish(com.exchangedemo.notification.proto.PublishPrivateEventRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getPublishMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do limited synchronous rpc calls to service PrivateNotificationService.
   */
  public static final class PrivateNotificationServiceBlockingStub
      extends io.grpc.stub.AbstractBlockingStub<PrivateNotificationServiceBlockingStub> {
    private PrivateNotificationServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected PrivateNotificationServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new PrivateNotificationServiceBlockingStub(channel, callOptions);
    }

    /**
     */
    public com.exchangedemo.notification.proto.PublishPrivateEventResponse publish(com.exchangedemo.notification.proto.PublishPrivateEventRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getPublishMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do ListenableFuture-style rpc calls to service PrivateNotificationService.
   */
  public static final class PrivateNotificationServiceFutureStub
      extends io.grpc.stub.AbstractFutureStub<PrivateNotificationServiceFutureStub> {
    private PrivateNotificationServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected PrivateNotificationServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new PrivateNotificationServiceFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.exchangedemo.notification.proto.PublishPrivateEventResponse> publish(
        com.exchangedemo.notification.proto.PublishPrivateEventRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getPublishMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_PUBLISH = 0;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final AsyncService serviceImpl;
    private final int methodId;

    MethodHandlers(AsyncService serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_PUBLISH:
          serviceImpl.publish((com.exchangedemo.notification.proto.PublishPrivateEventRequest) request,
              (io.grpc.stub.StreamObserver<com.exchangedemo.notification.proto.PublishPrivateEventResponse>) responseObserver);
          break;
        default:
          throw new AssertionError();
      }
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        default:
          throw new AssertionError();
      }
    }
  }

  public static final io.grpc.ServerServiceDefinition bindService(AsyncService service) {
    return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
        .addMethod(
          getPublishMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.exchangedemo.notification.proto.PublishPrivateEventRequest,
              com.exchangedemo.notification.proto.PublishPrivateEventResponse>(
                service, METHODID_PUBLISH)))
        .build();
  }

  private static abstract class PrivateNotificationServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    PrivateNotificationServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return com.exchangedemo.notification.proto.PrivateEventProto.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("PrivateNotificationService");
    }
  }

  private static final class PrivateNotificationServiceFileDescriptorSupplier
      extends PrivateNotificationServiceBaseDescriptorSupplier {
    PrivateNotificationServiceFileDescriptorSupplier() {}
  }

  private static final class PrivateNotificationServiceMethodDescriptorSupplier
      extends PrivateNotificationServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final java.lang.String methodName;

    PrivateNotificationServiceMethodDescriptorSupplier(java.lang.String methodName) {
      this.methodName = methodName;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.MethodDescriptor getMethodDescriptor() {
      return getServiceDescriptor().findMethodByName(methodName);
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (PrivateNotificationServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new PrivateNotificationServiceFileDescriptorSupplier())
              .addMethod(getPublishMethod())
              .build();
        }
      }
    }
    return result;
  }
}
