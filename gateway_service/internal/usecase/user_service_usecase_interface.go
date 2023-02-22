package usecase

import (
	"context"

	userPb "github.com/ideaspaper/social-media-proto/user"
)

type IUserServiceUsecase interface {
	FindUserByID(ctx context.Context, userPbReq *userPb.Req) (*userPb.Resp, error)
	CreateUser(ctx context.Context, userPbReq *userPb.Req) (*userPb.Resp, error)
	DeleteUserByID(ctx context.Context, userPbReq *userPb.Req) (*userPb.Resp, error)
	DeleteUserPermanentlyByID(ctx context.Context, userPbReq *userPb.Req) (*userPb.Resp, error)
}
