package main

import (
	"fmt"
	"gatewayservice/cmd/http_service/internal/handler"
	"gatewayservice/cmd/http_service/internal/middleware"
	"gatewayservice/cmd/http_service/internal/router"
	"gatewayservice/internal/usecase"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	userPb "github.com/ideaspaper/social-media-proto/user"
)

func initLogger() *slog.Logger {
	logStringLevel := map[string]slog.Level{
		"DEBUG": slog.LevelDebug,
		"INFO":  slog.LevelInfo,
		"WARN":  slog.LevelWarn,
		"ERROR": slog.LevelError,
	}
	opts := slog.HandlerOptions{
		Level:     logStringLevel[os.Getenv("LOG_LEVEL")],
		AddSource: true,
	}
	textHandler := opts.NewTextHandler(os.Stdout).WithAttrs(
		[]slog.Attr{
			slog.String("app-name", os.Getenv("APP_NAME")),
			slog.String("app-version", os.Getenv("APP_VERSION")),
		},
	)
	return slog.New(textHandler)
}

func main() {
	const appPort = "8081"
	gin.SetMode(gin.ReleaseMode)
	logger := initLogger()
	userServiceConn, err := grpc.Dial(
		fmt.Sprintf(
			"%s:%s",
			os.Getenv("USER_SERVICE_HOST"),
			os.Getenv("USER_SERVICE_PORT"),
		),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		logger.Error("Connecting to gRPC service failed", err)
		os.Exit(1)
	}
	defer userServiceConn.Close()
	validate := validator.New()
	userService := userPb.NewUserServiceClient(userServiceConn)
	userServiceUsecase := usecase.NewUserServiceUsecase(logger, validate, userService)
	handler := handler.New(logger, userServiceUsecase)
	middleware := middleware.New(logger)
	router := router.New(handler, middleware)
	logger.Info("Server listening", slog.String("port", appPort))
	if err := router.Run(fmt.Sprintf(":%s", appPort)); err != nil {
		logger.Error("Failed to serve", err)
		os.Exit(1)
	}
}
