package interceptor

import (
	"context"
	"time"

	"userservice/internal/util"

	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Interceptor struct {
	logger *slog.Logger
}

func NewInterceptor(logger *slog.Logger) *Interceptor {
	return &Interceptor{
		logger: logger,
	}
}

func (i Interceptor) Intercept(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	const scope = "interceptor#Intercept"
	start := time.Now()
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Internal, "No metadata provided")
	}
	requestID := md["request-id"]
	if len(requestID) == 0 {
		return nil, status.Error(codes.Internal, "No request ID provided")
	}
	ctx = context.WithValue(ctx, util.RequestID, requestID[0])
	h, err := handler(ctx, req)
	if err != nil {
		err = i.ErrorHandler(err)
	}
	stop := time.Now()
	i.logger.Info(
		"Handle request",
		slog.String("request_id", requestID[0]),
		slog.String("scope", scope),
		slog.String("method", info.FullMethod),
		slog.String("latency", stop.Sub(start).String()),
	)
	return h, err
}
