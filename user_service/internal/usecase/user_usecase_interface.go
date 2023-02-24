package usecase

import (
	"context"
	"userservice/internal/dto/req"
	"userservice/internal/dto/resp"
)

type IUserUsecase interface {
	FindByID(ctx context.Context, id int) (*resp.UserDto, error)
	DeleteByID(ctx context.Context, id int) (*resp.UserDto, error)
	DeletePermanentlyByID(ctx context.Context, id int) (*resp.UserDto, error)
	Register(ctx context.Context, userDto *req.UserDto) (*resp.UserDto, error)
	Login(ctx context.Context, loginDto *req.LoginDto) (*resp.LoginDto, error)
}
