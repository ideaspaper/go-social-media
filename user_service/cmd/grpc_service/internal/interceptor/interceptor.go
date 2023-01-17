package interceptor

import (
	"context"
	"time"

	"userservice/internal/util"

	userPb "github.com/ideaspaper/social-media-proto/user"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
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
	request := req.(*userPb.Req)
	ctx = context.WithValue(ctx, util.RequestID, request.RequestID)
	h, err := handler(ctx, req)
	if err != nil {
		err = i.ErrorHandler(err)
	}
	stop := time.Now()
	i.logger.Info(
		"Handle request",
		slog.String("request_id", request.RequestID),
		slog.String("scope", scope),
		slog.String("method", info.FullMethod),
		slog.String("latency", stop.Sub(start).String()),
	)
	return h, err
}
