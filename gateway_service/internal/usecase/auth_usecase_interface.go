package usecase

import (
	"context"
	"gatewayservice/internal/dto/req"
)

type IAuthUsecase interface {
	Login(ctx context.Context, loginDto req.LoginDto)
}
