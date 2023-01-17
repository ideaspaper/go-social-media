package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"userservice/cmd/config"
	"userservice/cmd/grpc_service/internal/handler"
	"userservice/cmd/grpc_service/internal/interceptor"
	"userservice/internal/repository/pg"
	"userservice/internal/usecase"

	userPb "github.com/ideaspaper/social-media-proto/user"

	"github.com/go-playground/validator/v10"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
)

var logStringLevel = map[string]slog.Level{
	"DEBUG": slog.LevelDebug,
	"INFO":  slog.LevelInfo,
	"WARN":  slog.LevelWarn,
	"ERROR": slog.LevelError,
}

func initLogger() *slog.Logger {
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
	const appPort = "50051"
	if err := config.ConnectDB(); err != nil {
		log.Fatalln(err)
	}
	db := config.GetDB()
	logger := initLogger()
	validate := validator.New()
	userRepository := pg.NewUserRepository(logger, db)
	userUsecase := usecase.NewUserUsecase(logger, validate, userRepository)
	handler := handler.New(logger, userUsecase)
	interceptor := interceptor.NewInterceptor(logger)
	lis, err := net.Listen(
		"tcp",
		fmt.Sprintf(":%s", appPort),
	)
	if err != nil {
		logger.Error("Failed to listen", err)
		os.Exit(1)
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(interceptor.Intercept))
	userPb.RegisterUserServiceServer(s, handler)
	logger.Info("Server listening", slog.String("port", appPort))
	if err := s.Serve(lis); err != nil {
		logger.Error("Failed to serve", err)
		os.Exit(1)
	}
}
