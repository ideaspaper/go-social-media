package handler

import (
	"gatewayservice/internal/usecase"

	"golang.org/x/exp/slog"
)

type Handler struct {
	logger             *slog.Logger
	userServiceUsecase usecase.IUserServiceUsecase
}

func New(logger *slog.Logger, userServiceUsecase usecase.IUserServiceUsecase) *Handler {
	return &Handler{
		logger:             logger,
		userServiceUsecase: userServiceUsecase,
	}
}
