package com.exchangedemo.ledger.proto;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@io.grpc.stub.annotations.GrpcGenerated
public final class LedgerServiceGrpc {

  private LedgerServiceGrpc() {}

  public static final java.lang.String SERVICE_NAME = "exchange.ledger.v1.LedgerService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.exchangedemo.ledger.proto.ReserveFundsRequest,
      com.exchangedemo.ledger.proto.ReserveFundsResponse> getReserveFundsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ReserveFunds",
      requestType = com.exchangedemo.ledger.proto.ReserveFundsRequest.class,
      responseType = com.exchangedemo.ledger.proto.ReserveFundsResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.exchangedemo.ledger.proto.ReserveFundsRequest,
      com.exchangedemo.ledger.proto.ReserveFundsResponse> getReserveFundsMethod() {
    io.grpc.MethodDescriptor<com.exchangedemo.ledger.proto.ReserveFundsRequest, com.exchangedemo.ledger.proto.ReserveFundsResponse> getReserveFundsMethod;
    if ((getReserveFundsMethod = LedgerServiceGrpc.getReserveFundsMethod) == null) {
      synchronized (LedgerServiceGrpc.class) {
        if ((getReserveFundsMethod = LedgerServiceGrpc.getReserveFundsMethod) == null) {
          LedgerServiceGrpc.getReserveFundsMethod = getReserveFundsMethod =
              io.grpc.MethodDescriptor.<com.exchangedemo.ledger.proto.ReserveFundsRequest, com.exchangedemo.ledger.proto.ReserveFundsResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ReserveFunds"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.exchangedemo.ledger.proto.ReserveFundsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.exchangedemo.ledger.proto.ReserveFundsResponse.getDefaultInstance()))
              .setSchemaDescriptor(new LedgerServiceMethodDescriptorSupplier("ReserveFunds"))
              .build();
        }
      }
    }
    return getReserveFundsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.exchangedemo.ledger.proto.ReleaseFundsRequest,
      com.exchangedemo.ledger.proto.ReleaseFundsResponse> getReleaseFundsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ReleaseFunds",
      requestType = com.exchangedemo.ledger.proto.ReleaseFundsRequest.class,
      responseType = com.exchangedemo.ledger.proto.ReleaseFundsResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.exchangedemo.ledger.proto.ReleaseFundsRequest,
      com.exchangedemo.ledger.proto.ReleaseFundsResponse> getReleaseFundsMethod() {
    io.grpc.MethodDescriptor<com.exchangedemo.ledger.proto.ReleaseFundsRequest, com.exchangedemo.ledger.proto.ReleaseFundsResponse> getReleaseFundsMethod;
    if ((getReleaseFundsMethod = LedgerServiceGrpc.getReleaseFundsMethod) == null) {
      synchronized (LedgerServiceGrpc.class) {
        if ((getReleaseFundsMethod = LedgerServiceGrpc.getReleaseFundsMethod) == null) {
          LedgerServiceGrpc.getReleaseFundsMethod = getReleaseFundsMethod =
              io.grpc.MethodDescriptor.<com.exchangedemo.ledger.proto.ReleaseFundsRequest, com.exchangedemo.ledger.proto.ReleaseFundsResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ReleaseFunds"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.exchangedemo.ledger.proto.ReleaseFundsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.exchangedemo.ledger.proto.ReleaseFundsResponse.getDefaultInstance()))
              .setSchemaDescriptor(new LedgerServiceMethodDescriptorSupplier("ReleaseFunds"))
              .build();
        }
      }
    }
    return getReleaseFundsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.exchangedemo.ledger.proto.ApplyTradeRequest,
      com.exchangedemo.ledger.proto.ApplyTradeResponse> getApplyTradeMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ApplyTrade",
      requestType = com.exchangedemo.ledger.proto.ApplyTradeRequest.class,
      responseType = com.exchangedemo.ledger.proto.ApplyTradeResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.exchangedemo.ledger.proto.ApplyTradeRequest,
      com.exchangedemo.ledger.proto.ApplyTradeResponse> getApplyTradeMethod() {
    io.grpc.MethodDescriptor<com.exchangedemo.ledger.proto.ApplyTradeRequest, com.exchangedemo.ledger.proto.ApplyTradeResponse> getApplyTradeMethod;
    if ((getApplyTradeMethod = LedgerServiceGrpc.getApplyTradeMethod) == null) {
      synchronized (LedgerServiceGrpc.class) {
        if ((getApplyTradeMethod = LedgerServiceGrpc.getApplyTradeMethod) == null) {
          LedgerServiceGrpc.getApplyTradeMethod = getApplyTradeMethod =
              io.grpc.MethodDescriptor.<com.exchangedemo.ledger.proto.ApplyTradeRequest, com.exchangedemo.ledger.proto.ApplyTradeResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ApplyTrade"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.exchangedemo.ledger.proto.ApplyTradeRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.exchangedemo.ledger.proto.ApplyTradeResponse.getDefaultInstance()))
              .setSchemaDescriptor(new LedgerServiceMethodDescriptorSupplier("ApplyTrade"))
              .build();
        }
      }
    }
    return getApplyTradeMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.exchangedemo.ledger.proto.GetBalanceRequest,
      com.exchangedemo.ledger.proto.GetBalanceResponse> getGetBalanceMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetBalance",
      requestType = com.exchangedemo.ledger.proto.GetBalanceRequest.class,
      responseType = com.exchangedemo.ledger.proto.GetBalanceResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.exchangedemo.ledger.proto.GetBalanceRequest,
      com.exchangedemo.ledger.proto.GetBalanceResponse> getGetBalanceMethod() {
    io.grpc.MethodDescriptor<com.exchangedemo.ledger.proto.GetBalanceRequest, com.exchangedemo.ledger.proto.GetBalanceResponse> getGetBalanceMethod;
    if ((getGetBalanceMethod = LedgerServiceGrpc.getGetBalanceMethod) == null) {
      synchronized (LedgerServiceGrpc.class) {
        if ((getGetBalanceMethod = LedgerServiceGrpc.getGetBalanceMethod) == null) {
          LedgerServiceGrpc.getGetBalanceMethod = getGetBalanceMethod =
              io.grpc.MethodDescriptor.<com.exchangedemo.ledger.proto.GetBalanceRequest, com.exchangedemo.ledger.proto.GetBalanceResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetBalance"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.exchangedemo.ledger.proto.GetBalanceRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.exchangedemo.ledger.proto.GetBalanceResponse.getDefaultInstance()))
              .setSchemaDescriptor(new LedgerServiceMethodDescriptorSupplier("GetBalance"))
              .build();
        }
      }
    }
    return getGetBalanceMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static LedgerServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<LedgerServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<LedgerServiceStub>() {
        @java.lang.Override
        public LedgerServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new LedgerServiceStub(channel, callOptions);
        }
      };
    return LedgerServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports all types of calls on the service
   */
  public static LedgerServiceBlockingV2Stub newBlockingV2Stub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<LedgerServiceBlockingV2Stub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<LedgerServiceBlockingV2Stub>() {
        @java.lang.Override
        public LedgerServiceBlockingV2Stub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new LedgerServiceBlockingV2Stub(channel, callOptions);
        }
      };
    return LedgerServiceBlockingV2Stub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static LedgerServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<LedgerServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<LedgerServiceBlockingStub>() {
        @java.lang.Override
        public LedgerServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new LedgerServiceBlockingStub(channel, callOptions);
        }
      };
    return LedgerServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static LedgerServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<LedgerServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<LedgerServiceFutureStub>() {
        @java.lang.Override
        public LedgerServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new LedgerServiceFutureStub(channel, callOptions);
        }
      };
    return LedgerServiceFutureStub.newStub(factory, channel);
  }

  /**
   */
  public interface AsyncService {

    /**
     */
    default void reserveFunds(com.exchangedemo.ledger.proto.ReserveFundsRequest request,
        io.grpc.stub.StreamObserver<com.exchangedemo.ledger.proto.ReserveFundsResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getReserveFundsMethod(), responseObserver);
    }

    /**
     */
    default void releaseFunds(com.exchangedemo.ledger.proto.ReleaseFundsRequest request,
        io.grpc.stub.StreamObserver<com.exchangedemo.ledger.proto.ReleaseFundsResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getReleaseFundsMethod(), responseObserver);
    }

    /**
     */
    default void applyTrade(com.exchangedemo.ledger.proto.ApplyTradeRequest request,
        io.grpc.stub.StreamObserver<com.exchangedemo.ledger.proto.ApplyTradeResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getApplyTradeMethod(), responseObserver);
    }

    /**
     */
    default void getBalance(com.exchangedemo.ledger.proto.GetBalanceRequest request,
        io.grpc.stub.StreamObserver<com.exchangedemo.ledger.proto.GetBalanceResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetBalanceMethod(), responseObserver);
    }
  }

  /**
   * Base class for the server implementation of the service LedgerService.
   */
  public static abstract class LedgerServiceImplBase
      implements io.grpc.BindableService, AsyncService {

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return LedgerServiceGrpc.bindService(this);
    }
  }

  /**
   * A stub to allow clients to do asynchronous rpc calls to service LedgerService.
   */
  public static final class LedgerServiceStub
      extends io.grpc.stub.AbstractAsyncStub<LedgerServiceStub> {
    private LedgerServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected LedgerServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new LedgerServiceStub(channel, callOptions);
    }

    /**
     */
    public void reserveFunds(com.exchangedemo.ledger.proto.ReserveFundsRequest request,
        io.grpc.stub.StreamObserver<com.exchangedemo.ledger.proto.ReserveFundsResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getReserveFundsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void releaseFunds(com.exchangedemo.ledger.proto.ReleaseFundsRequest request,
        io.grpc.stub.StreamObserver<com.exchangedemo.ledger.proto.ReleaseFundsResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getReleaseFundsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void applyTrade(com.exchangedemo.ledger.proto.ApplyTradeRequest request,
        io.grpc.stub.StreamObserver<com.exchangedemo.ledger.proto.ApplyTradeResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getApplyTradeMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getBalance(com.exchangedemo.ledger.proto.GetBalanceRequest request,
        io.grpc.stub.StreamObserver<com.exchangedemo.ledger.proto.GetBalanceResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetBalanceMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * A stub to allow clients to do synchronous rpc calls to service LedgerService.
   */
  public static final class LedgerServiceBlockingV2Stub
      extends io.grpc.stub.AbstractBlockingStub<LedgerServiceBlockingV2Stub> {
    private LedgerServiceBlockingV2Stub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected LedgerServiceBlockingV2Stub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new LedgerServiceBlockingV2Stub(channel, callOptions);
    }

    /**
     */
    public com.exchangedemo.ledger.proto.ReserveFundsResponse reserveFunds(com.exchangedemo.ledger.proto.ReserveFundsRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getReserveFundsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.exchangedemo.ledger.proto.ReleaseFundsResponse releaseFunds(com.exchangedemo.ledger.proto.ReleaseFundsRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getReleaseFundsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.exchangedemo.ledger.proto.ApplyTradeResponse applyTrade(com.exchangedemo.ledger.proto.ApplyTradeRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getApplyTradeMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.exchangedemo.ledger.proto.GetBalanceResponse getBalance(com.exchangedemo.ledger.proto.GetBalanceRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetBalanceMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do limited synchronous rpc calls to service LedgerService.
   */
  public static final class LedgerServiceBlockingStub
      extends io.grpc.stub.AbstractBlockingStub<LedgerServiceBlockingStub> {
    private LedgerServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected LedgerServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new LedgerServiceBlockingStub(channel, callOptions);
    }

    /**
     */
    public com.exchangedemo.ledger.proto.ReserveFundsResponse reserveFunds(com.exchangedemo.ledger.proto.ReserveFundsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getReserveFundsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.exchangedemo.ledger.proto.ReleaseFundsResponse releaseFunds(com.exchangedemo.ledger.proto.ReleaseFundsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getReleaseFundsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.exchangedemo.ledger.proto.ApplyTradeResponse applyTrade(com.exchangedemo.ledger.proto.ApplyTradeRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getApplyTradeMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.exchangedemo.ledger.proto.GetBalanceResponse getBalance(com.exchangedemo.ledger.proto.GetBalanceRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetBalanceMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do ListenableFuture-style rpc calls to service LedgerService.
   */
  public static final class LedgerServiceFutureStub
      extends io.grpc.stub.AbstractFutureStub<LedgerServiceFutureStub> {
    private LedgerServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected LedgerServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new LedgerServiceFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.exchangedemo.ledger.proto.ReserveFundsResponse> reserveFunds(
        com.exchangedemo.ledger.proto.ReserveFundsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getReserveFundsMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.exchangedemo.ledger.proto.ReleaseFundsResponse> releaseFunds(
        com.exchangedemo.ledger.proto.ReleaseFundsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getReleaseFundsMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.exchangedemo.ledger.proto.ApplyTradeResponse> applyTrade(
        com.exchangedemo.ledger.proto.ApplyTradeRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getApplyTradeMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.exchangedemo.ledger.proto.GetBalanceResponse> getBalance(
        com.exchangedemo.ledger.proto.GetBalanceRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetBalanceMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_RESERVE_FUNDS = 0;
  private static final int METHODID_RELEASE_FUNDS = 1;
  private static final int METHODID_APPLY_TRADE = 2;
  private static final int METHODID_GET_BALANCE = 3;

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
        case METHODID_RESERVE_FUNDS:
          serviceImpl.reserveFunds((com.exchangedemo.ledger.proto.ReserveFundsRequest) request,
              (io.grpc.stub.StreamObserver<com.exchangedemo.ledger.proto.ReserveFundsResponse>) responseObserver);
          break;
        case METHODID_RELEASE_FUNDS:
          serviceImpl.releaseFunds((com.exchangedemo.ledger.proto.ReleaseFundsRequest) request,
              (io.grpc.stub.StreamObserver<com.exchangedemo.ledger.proto.ReleaseFundsResponse>) responseObserver);
          break;
        case METHODID_APPLY_TRADE:
          serviceImpl.applyTrade((com.exchangedemo.ledger.proto.ApplyTradeRequest) request,
              (io.grpc.stub.StreamObserver<com.exchangedemo.ledger.proto.ApplyTradeResponse>) responseObserver);
          break;
        case METHODID_GET_BALANCE:
          serviceImpl.getBalance((com.exchangedemo.ledger.proto.GetBalanceRequest) request,
              (io.grpc.stub.StreamObserver<com.exchangedemo.ledger.proto.GetBalanceResponse>) responseObserver);
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
          getReserveFundsMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.exchangedemo.ledger.proto.ReserveFundsRequest,
              com.exchangedemo.ledger.proto.ReserveFundsResponse>(
                service, METHODID_RESERVE_FUNDS)))
        .addMethod(
          getReleaseFundsMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.exchangedemo.ledger.proto.ReleaseFundsRequest,
              com.exchangedemo.ledger.proto.ReleaseFundsResponse>(
                service, METHODID_RELEASE_FUNDS)))
        .addMethod(
          getApplyTradeMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.exchangedemo.ledger.proto.ApplyTradeRequest,
              com.exchangedemo.ledger.proto.ApplyTradeResponse>(
                service, METHODID_APPLY_TRADE)))
        .addMethod(
          getGetBalanceMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.exchangedemo.ledger.proto.GetBalanceRequest,
              com.exchangedemo.ledger.proto.GetBalanceResponse>(
                service, METHODID_GET_BALANCE)))
        .build();
  }

  private static abstract class LedgerServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    LedgerServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return com.exchangedemo.ledger.proto.LedgerServiceProto.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("LedgerService");
    }
  }

  private static final class LedgerServiceFileDescriptorSupplier
      extends LedgerServiceBaseDescriptorSupplier {
    LedgerServiceFileDescriptorSupplier() {}
  }

  private static final class LedgerServiceMethodDescriptorSupplier
      extends LedgerServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final java.lang.String methodName;

    LedgerServiceMethodDescriptorSupplier(java.lang.String methodName) {
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
      synchronized (LedgerServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new LedgerServiceFileDescriptorSupplier())
              .addMethod(getReserveFundsMethod())
              .addMethod(getReleaseFundsMethod())
              .addMethod(getApplyTradeMethod())
              .addMethod(getGetBalanceMethod())
              .build();
        }
      }
    }
    return result;
  }
}
