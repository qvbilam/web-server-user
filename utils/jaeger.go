package utils

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"user/enum"
)

type MDCarrier struct {
	metadata.MD
}

func (m MDCarrier) ForeachKey(handler func(key, val string) error) error {
	for k, vs := range m.MD {
		for _, v := range vs {
			if err := handler(k, v); err != nil {
				return err
			}
		}
	}
	return nil
}

func (m MDCarrier) Set(key, val string) {
	m.MD[key] = append(m.MD[key], val)
}

func InitTracerContext(ginCtx *gin.Context, span opentracing.Span) context.Context {
	ctx := context.Background()
	ctxTracer := ginCtx.Value(enum.ContextTracerKey)
	if ctxTracer != nil {
		ctx = context.WithValue(ctx, enum.ContextTracerKey, ctxTracer)
	}
	ctxSpan := ginCtx.Value(enum.ContextSpanKey)
	if ctxSpan != nil {
		ctx = context.WithValue(ctx, enum.ContextSpanKey, ctxSpan)
	}

	opentracing.ContextWithSpan(ctx, span)
	return ctx
}

func ClientInterceptor(tracer opentracing.Tracer) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, request, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		//一个RPC调用的服务端的span，和RPC服务客户端的span构成ChildOf关系
		var parentCtx opentracing.SpanContext
		parentSpan := opentracing.SpanFromContext(ctx)
		if parentSpan != nil {
			parentCtx = parentSpan.Context()
		}

		// 通过 ctx 获取 tracer 后重新定义 parentSpanContext
		ctxSpan := ctx.Value(enum.ContextSpanKey)
		if ctxSpan != nil {
			parentCtx = ctxSpan.(*jaeger.Span).Context()
		}
		// 通过 ctx 获取 tracer 后重新定义
		ctxTracer := ctx.Value(enum.ContextTracerKey)
		if ctxTracer != nil {
			tracer = ctxTracer.(opentracing.Tracer)
		}

		span := tracer.StartSpan(
			method,
			opentracing.ChildOf(parentCtx),
			opentracing.Tag{Key: string(ext.Component), Value: "gRPC Client"},
			ext.SpanKindRPCClient,
		)

		defer span.Finish()
		md, ok := metadata.FromOutgoingContext(ctx)
		if !ok {
			md = metadata.New(nil)
		} else {
			md = md.Copy()
		}

		err := tracer.Inject(
			span.Context(),
			opentracing.TextMap,
			MDCarrier{md}, // 自定义 carrier
		)

		newCtx := metadata.NewOutgoingContext(ctx, md)
		err = invoker(newCtx, method, request, reply, cc, opts...)

		return err
	}
}

func ServerInterceptor(tracer opentracing.Tracer) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			md = metadata.New(nil)
		}
		spanContext, err := tracer.Extract(
			opentracing.TextMap,
			MDCarrier{md},
		)

		if err != nil && err != opentracing.ErrSpanContextNotFound {
			grpclog.Errorf("extract from metadata err: %v", err)
		} else {
			span := tracer.StartSpan(
				info.FullMethod,
				ext.RPCServerOption(spanContext),
				opentracing.Tag{Key: string(ext.Component), Value: "gRPC Server"},
				ext.SpanKindRPCServer,
			)
			defer span.Finish()

			ctx = opentracing.ContextWithSpan(ctx, span)
		}

		return handler(ctx, req)

	}

}
