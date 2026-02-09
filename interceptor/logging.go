package interceptor

import (
	"context"
	"strings"
	"time"

	"github.com/jjtsksn/utils/logctx"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func LoggingUnaryInterceptor(logger *zap.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		if strings.HasPrefix(info.FullMethod, "/grpc.health.v1.") {
			return handler(ctx, req)
		}

		l := logger

		spanCtx := trace.SpanContextFromContext(ctx)
		if spanCtx.IsValid() {
			traceID := spanCtx.TraceID().String()
			spanID := spanCtx.SpanID().String()
			l = logger.With(
				zap.String("trace_id", traceID),
				zap.String("span_id", spanID))
		}

		ctx = logctx.WithLogger(ctx, l)

		start := time.Now()
		LogIncomingRequest(ctx, info.FullMethod, req)
		resp, err = handler(ctx, req)
		LogOutgoingResponse(ctx, info.FullMethod, resp, err, start)

		return resp, err
	}
}
