package com.exchangedemo.matching.proto;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@io.grpc.stub.annotations.GrpcGenerated
public final class MatchingEngineServiceGrpc {

  private MatchingEngineServiceGrpc() {}

  public static final java.lang.String SERVICE_NAME = "exchange.matching.v1.MatchingEngineService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.exchangedemo.matching.proto.ApplyOrderRequest,
      com.exchangedemo.matching.proto.ApplyOrderResponse> getApplyOrderMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ApplyOrder",
      requestType = com.exchangedemo.matching.proto.ApplyOrderRequest.class,
      responseType = com.exchangedemo.matching.proto.ApplyOrderResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.exchangedemo.matching.proto.ApplyOrderRequest,
      com.exchangedemo.matching.proto.ApplyOrderResponse> getApplyOrderMethod() {
    io.grpc.MethodDescriptor<com.exchangedemo.matching.proto.ApplyOrderRequest, com.exchangedemo.matching.proto.ApplyOrderResponse> getApplyOrderMethod;
    if ((getApplyOrderMethod = MatchingEngineServiceGrpc.getApplyOrderMethod) == null) {
      synchronized (MatchingEngineServiceGrpc.class) {
        if ((getApplyOrderMethod = MatchingEngineServiceGrpc.getApplyOrderMethod) == null) {
          MatchingEngineServiceGrpc.getApplyOrderMethod = getApplyOrderMethod =
              io.grpc.MethodDescriptor.<com.exchangedemo.matching.proto.ApplyOrderRequest, com.exchangedemo.matching.proto.ApplyOrderResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ApplyOrder"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.exchangedemo.matching.proto.ApplyOrderRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.exchangedemo.matching.proto.ApplyOrderResponse.getDefaultInstance()))
              .setSchemaDescriptor(new MatchingEngineServiceMethodDescriptorSupplier("ApplyOrder"))
              .build();
        }
      }
    }
    return getApplyOrderMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.exchangedemo.matching.proto.CancelOrderRequest,
      com.exchangedemo.matching.proto.CancelOrderResponse> getCancelOrderMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "CancelOrder",
      requestType = com.exchangedemo.matching.proto.CancelOrderRequest.class,
      responseType = com.exchangedemo.matching.proto.CancelOrderResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.exchangedemo.matching.proto.CancelOrderRequest,
      com.exchangedemo.matching.proto.CancelOrderResponse> getCancelOrderMethod() {
    io.grpc.MethodDescriptor<com.exchangedemo.matching.proto.CancelOrderRequest, com.exchangedemo.matching.proto.CancelOrderResponse> getCancelOrderMethod;
    if ((getCancelOrderMethod = MatchingEngineServiceGrpc.getCancelOrderMethod) == null) {
      synchronized (MatchingEngineServiceGrpc.class) {
        if ((getCancelOrderMethod = MatchingEngineServiceGrpc.getCancelOrderMethod) == null) {
          MatchingEngineServiceGrpc.getCancelOrderMethod = getCancelOrderMethod =
              io.grpc.MethodDescriptor.<com.exchangedemo.matching.proto.CancelOrderRequest, com.exchangedemo.matching.proto.CancelOrderResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "CancelOrder"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.exchangedemo.matching.proto.CancelOrderRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.exchangedemo.matching.proto.CancelOrderResponse.getDefaultInstance()))
              .setSchemaDescriptor(new MatchingEngineServiceMethodDescriptorSupplier("CancelOrder"))
              .build();
        }
      }
    }
    return getCancelOrderMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static MatchingEngineServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<MatchingEngineServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<MatchingEngineServiceStub>() {
        @java.lang.Override
        public MatchingEngineServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new MatchingEngineServiceStub(channel, callOptions);
        }
      };
    return MatchingEngineServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports all types of calls on the service
   */
  public static MatchingEngineServiceBlockingV2Stub newBlockingV2Stub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<MatchingEngineServiceBlockingV2Stub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<MatchingEngineServiceBlockingV2Stub>() {
        @java.lang.Override
        public MatchingEngineServiceBlockingV2Stub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new MatchingEngineServiceBlockingV2Stub(channel, callOptions);
        }
      };
    return MatchingEngineServiceBlockingV2Stub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static MatchingEngineServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<MatchingEngineServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<MatchingEngineServiceBlockingStub>() {
        @java.lang.Override
        public MatchingEngineServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new MatchingEngineServiceBlockingStub(channel, callOptions);
        }
      };
    return MatchingEngineServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static MatchingEngineServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<MatchingEngineServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<MatchingEngineServiceFutureStub>() {
        @java.lang.Override
        public MatchingEngineServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new MatchingEngineServiceFutureStub(channel, callOptions);
        }
      };
    return MatchingEngineServiceFutureStub.newStub(factory, channel);
  }

  /**
   */
  public interface AsyncService {

    /**
     */
    default void applyOrder(com.exchangedemo.matching.proto.ApplyOrderRequest request,
        io.grpc.stub.StreamObserver<com.exchangedemo.matching.proto.ApplyOrderResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getApplyOrderMethod(), responseObserver);
    }

    /**
     */
    default void cancelOrder(com.exchangedemo.matching.proto.CancelOrderRequest request,
        io.grpc.stub.StreamObserver<com.exchangedemo.matching.proto.CancelOrderResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCancelOrderMethod(), responseObserver);
    }
  }

  /**
   * Base class for the server implementation of the service MatchingEngineService.
   */
  public static abstract class MatchingEngineServiceImplBase
      implements io.grpc.BindableService, AsyncService {

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return MatchingEngineServiceGrpc.bindService(this);
    }
  }

  /**
   * A stub to allow clients to do asynchronous rpc calls to service MatchingEngineService.
   */
  public static final class MatchingEngineServiceStub
      extends io.grpc.stub.AbstractAsyncStub<MatchingEngineServiceStub> {
    private MatchingEngineServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected MatchingEngineServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new MatchingEngineServiceStub(channel, callOptions);
    }

    /**
     */
    public void applyOrder(com.exchangedemo.matching.proto.ApplyOrderRequest request,
        io.grpc.stub.StreamObserver<com.exchangedemo.matching.proto.ApplyOrderResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getApplyOrderMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void cancelOrder(com.exchangedemo.matching.proto.CancelOrderRequest request,
        io.grpc.stub.StreamObserver<com.exchangedemo.matching.proto.CancelOrderResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCancelOrderMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * A stub to allow clients to do synchronous rpc calls to service MatchingEngineService.
   */
  public static final class MatchingEngineServiceBlockingV2Stub
      extends io.grpc.stub.AbstractBlockingStub<MatchingEngineServiceBlockingV2Stub> {
    private MatchingEngineServiceBlockingV2Stub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected MatchingEngineServiceBlockingV2Stub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new MatchingEngineServiceBlockingV2Stub(channel, callOptions);
    }

    /**
     */
    public com.exchangedemo.matching.proto.ApplyOrderResponse applyOrder(com.exchangedemo.matching.proto.ApplyOrderRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getApplyOrderMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.exchangedemo.matching.proto.CancelOrderResponse cancelOrder(com.exchangedemo.matching.proto.CancelOrderRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getCancelOrderMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do limited synchronous rpc calls to service MatchingEngineService.
   */
  public static final class MatchingEngineServiceBlockingStub
      extends io.grpc.stub.AbstractBlockingStub<MatchingEngineServiceBlockingStub> {
    private MatchingEngineServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected MatchingEngineServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new MatchingEngineServiceBlockingStub(channel, callOptions);
    }

    /**
     */
    public com.exchangedemo.matching.proto.ApplyOrderResponse applyOrder(com.exchangedemo.matching.proto.ApplyOrderRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getApplyOrderMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.exchangedemo.matching.proto.CancelOrderResponse cancelOrder(com.exchangedemo.matching.proto.CancelOrderRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCancelOrderMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do ListenableFuture-style rpc calls to service MatchingEngineService.
   */
  public static final class MatchingEngineServiceFutureStub
      extends io.grpc.stub.AbstractFutureStub<MatchingEngineServiceFutureStub> {
    private MatchingEngineServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected MatchingEngineServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new MatchingEngineServiceFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.exchangedemo.matching.proto.ApplyOrderResponse> applyOrder(
        com.exchangedemo.matching.proto.ApplyOrderRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getApplyOrderMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.exchangedemo.matching.proto.CancelOrderResponse> cancelOrder(
        com.exchangedemo.matching.proto.CancelOrderRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCancelOrderMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_APPLY_ORDER = 0;
  private static final int METHODID_CANCEL_ORDER = 1;

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
        case METHODID_APPLY_ORDER:
          serviceImpl.applyOrder((com.exchangedemo.matching.proto.ApplyOrderRequest) request,
              (io.grpc.stub.StreamObserver<com.exchangedemo.matching.proto.ApplyOrderResponse>) responseObserver);
          break;
        case METHODID_CANCEL_ORDER:
          serviceImpl.cancelOrder((com.exchangedemo.matching.proto.CancelOrderRequest) request,
              (io.grpc.stub.StreamObserver<com.exchangedemo.matching.proto.CancelOrderResponse>) responseObserver);
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
          getApplyOrderMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.exchangedemo.matching.proto.ApplyOrderRequest,
              com.exchangedemo.matching.proto.ApplyOrderResponse>(
                service, METHODID_APPLY_ORDER)))
        .addMethod(
          getCancelOrderMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.exchangedemo.matching.proto.CancelOrderRequest,
              com.exchangedemo.matching.proto.CancelOrderResponse>(
                service, METHODID_CANCEL_ORDER)))
        .build();
  }

  private static abstract class MatchingEngineServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    MatchingEngineServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return com.exchangedemo.matching.proto.MatchingEngineProto.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("MatchingEngineService");
    }
  }

  private static final class MatchingEngineServiceFileDescriptorSupplier
      extends MatchingEngineServiceBaseDescriptorSupplier {
    MatchingEngineServiceFileDescriptorSupplier() {}
  }

  private static final class MatchingEngineServiceMethodDescriptorSupplier
      extends MatchingEngineServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final java.lang.String methodName;

    MatchingEngineServiceMethodDescriptorSupplier(java.lang.String methodName) {
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
      synchronized (MatchingEngineServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new MatchingEngineServiceFileDescriptorSupplier())
              .addMethod(getApplyOrderMethod())
              .addMethod(getCancelOrderMethod())
              .build();
        }
      }
    }
    return result;
  }
}
