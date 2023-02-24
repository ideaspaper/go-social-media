package usecase

import (
	"context"
	"gatewayservice/internal/dto/req"

	userPb "github.com/ideaspaper/social-media-proto/user"
)

type IUserServiceUsecase interface {
	FindUserByID(ctx context.Context, userID int) (*userPb.FindByIDResp, error)
	DeleteUserByID(ctx context.Context, userID int) (*userPb.DeleteByIDResp, error)
	DeleteUserPermanentlyByID(ctx context.Context, userID int) (*userPb.DeletePermanentlyByIDResp, error)
	Register(ctx context.Context, registerDto *req.UserDto) (*userPb.RegisterResp, error)
	Login(ctx context.Context, loginDto *req.LoginDto) (*userPb.LoginResp, error)
}
