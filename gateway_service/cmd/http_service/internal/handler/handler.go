package handler

import (
	"gatewayservice/internal/usecase"

	"golang.org/x/exp/slog"

	"github.com/go-playground/validator/v10"
)

type Handler struct {
	logger             *slog.Logger
	validate           *validator.Validate
	userServiceUsecase usecase.IUserServiceUsecase
}

func New(logger *slog.Logger, validate *validator.Validate, userServiceUsecase usecase.IUserServiceUsecase) *Handler {
	return &Handler{
		logger:             logger,
		validate:           validate,
		userServiceUsecase: userServiceUsecase,
	}
}
