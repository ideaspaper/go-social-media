package handler

import (
	"userservice/internal/usecase"

	userPb "github.com/ideaspaper/social-media-proto/user"

	"golang.org/x/exp/slog"
)

type Handler struct {
	logger      *slog.Logger
	userUsecase usecase.IUserUsecase
	userPb.UnimplementedUserServiceServer
}

func New(logger *slog.Logger, userUsecase usecase.IUserUsecase) *Handler {
	return &Handler{
		logger:      logger,
		userUsecase: userUsecase,
	}
}
