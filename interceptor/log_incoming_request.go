package interceptor

import (
	"context"

	"github.com/jjtsksn/utils/logctx"
	"go.uber.org/zap"
)

func LogIncomingRequest(
	ctx context.Context,
	method string,
	req any,
) {
	log := logctx.FromContext(ctx)

	log.Info(
		"recieved incoming grpc request",
		zap.String("grpc.method", method),
		zap.Any("request", req),
	)
}
