package usecase

import (
	"context"
	"fmt"

	userPb "github.com/ideaspaper/social-media-proto/user"
	"golang.org/x/exp/slog"
)

type userServiceUsecase struct {
	logger            *slog.Logger
	userServiceClient userPb.UserServiceClient
}

func NewUserServiceUsecase(logger *slog.Logger, userServiceClient userPb.UserServiceClient) IUserServiceUsecase {
	return &userServiceUsecase{
		logger:            logger,
		userServiceClient: userServiceClient,
	}
}

func (u *userServiceUsecase) FindUserByID(ctx context.Context, userPbReq *userPb.Req) (*userPb.Resp, error) {
	const scope = "userServiceUsecase#FindUserByID"
	user, err := u.userServiceClient.FindByID(ctx, userPbReq)
	if err != nil {
		u.logger.Error(
			"Got error from service",
			err,
			slog.String("request_id", ctx.Value("request_id").(string)),
			slog.String("scope", scope),
		)
		return nil, fmt.Errorf("%s: %w", scope, ErrClientService.SetError(err))
	}
	return user, nil
}

func (u *userServiceUsecase) CreateUser(ctx context.Context, userPbReq *userPb.Req) (*userPb.Resp, error) {
	const scope = "userServiceUsecase#CreateUser"
	user, err := u.userServiceClient.Create(ctx, userPbReq)
	if err != nil {
		u.logger.Error(
			"Got error from service",
			err,
			slog.String("request_id", ctx.Value("request_id").(string)),
			slog.String("scope", scope),
		)
		return nil, fmt.Errorf("%s: %w", scope, ErrClientService.SetError(err))
	}
	return user, nil
}

func (u *userServiceUsecase) DeleteUserByID(ctx context.Context, userPbReq *userPb.Req) (*userPb.Resp, error) {
	const scope = "userServiceUsecase#DeleteUserByID"
	user, err := u.userServiceClient.DeleteByID(ctx, userPbReq)
	if err != nil {
		u.logger.Error(
			"Got error from service",
			err,
			slog.String("request_id", ctx.Value("request_id").(string)),
			slog.String("scope", scope),
		)
		return nil, fmt.Errorf("%s: %w", scope, ErrClientService.SetError(err))
	}
	return user, nil
}

func (u *userServiceUsecase) DeleteUserPermanentlyByID(ctx context.Context, userPbReq *userPb.Req) (*userPb.Resp, error) {
	const scope = "userServiceUsecase#DeleteUserPermanentlyByID"
	user, err := u.userServiceClient.DeletePermanentlyByID(ctx, userPbReq)
	if err != nil {
		u.logger.Error(
			"Got error from service",
			err,
			slog.String("request_id", ctx.Value("request_id").(string)),
			slog.String("scope", scope),
		)
		return nil, fmt.Errorf("%s: %w", scope, ErrClientService.SetError(err))
	}
	return user, nil
}
