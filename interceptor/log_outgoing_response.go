package interceptor

import (
	"context"
	"time"

	"github.com/jjtsksn/utils/logctx"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func LogOutgoingResponse(
	ctx context.Context,
	method string,
	resp any,
	err error,
	start time.Time,
) {
	log := logctx.FromContext(ctx)

	if err != nil {
		s, _ := status.FromError(err)

		level := zap.InfoLevel

		switch s.Code() {
		case codes.Canceled, codes.DeadlineExceeded:
			level = zap.WarnLevel
		case codes.Internal, codes.DataLoss, codes.Unavailable:
			level = zap.ErrorLevel
		}

		log.Check(level, "grpc request completed").
			Write(
				zap.String("grpc.method", method),
				zap.String("grpc.code", s.Code().String()),
				zap.Error(err),
				zap.Duration("duration", time.Since(start)),
			)
		return
	}

	log.Info(
		"grpc request completed",
		zap.String("grpc.method", method),
		zap.Any("response", resp),
		zap.Duration("duration", time.Since(start)),
	)
}
