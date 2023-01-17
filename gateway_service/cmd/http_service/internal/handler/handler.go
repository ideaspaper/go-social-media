package handler

import (
	"golang.org/x/exp/slog"

	"github.com/go-playground/validator/v10"
	userPb "github.com/ideaspaper/social-media-proto/user"
)

type Handler struct {
	logger            *slog.Logger
	validate          *validator.Validate
	userServiceClient userPb.UserServiceClient
}

func New(logger *slog.Logger, validate *validator.Validate, userServiceClient userPb.UserServiceClient) *Handler {
	return &Handler{
		logger:            logger,
		validate:          validate,
		userServiceClient: userServiceClient,
	}
}
