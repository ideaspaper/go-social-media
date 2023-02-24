package usecase

import (
	"context"
	"errors"
	"fmt"
	"gatewayservice/internal/dto/req"
	"gatewayservice/internal/util"
	"strings"

	"github.com/go-playground/validator/v10"
	userPb "github.com/ideaspaper/social-media-proto/user"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc/metadata"
)

type userServiceUsecase struct {
	logger            *slog.Logger
	validate          *validator.Validate
	userServiceClient userPb.UserServiceClient
}

func NewUserServiceUsecase(logger *slog.Logger, validate *validator.Validate, userServiceClient userPb.UserServiceClient) IUserServiceUsecase {
	return &userServiceUsecase{
		logger:            logger,
		validate:          validate,
		userServiceClient: userServiceClient,
	}
}

func (u *userServiceUsecase) FindUserByID(ctx context.Context, userID int) (*userPb.FindByIDResp, error) {
	const scope = "userServiceUsecase#FindUserByID"
	requestID := ctx.Value(util.RequestID).(string)
	mdCtx := metadata.NewOutgoingContext(ctx, metadata.Pairs("request-id", requestID))
	user, err := u.userServiceClient.FindByID(mdCtx, &userPb.FindByIDReq{
		Id: int64(userID),
	})
	if err != nil {
		u.logger.Error(
			"Got error from service",
			err,
			slog.String("request_id", requestID),
			slog.String("scope", scope),
		)
		return nil, fmt.Errorf("%s: %w", scope, ErrClientService.SetError(err))
	}
	return user, nil
}

func (u *userServiceUsecase) DeleteUserByID(ctx context.Context, userID int) (*userPb.DeleteByIDResp, error) {
	const scope = "userServiceUsecase#DeleteUserByID"
	requestID := ctx.Value(util.RequestID).(string)
	mdCtx := metadata.NewOutgoingContext(ctx, metadata.Pairs("request-id", requestID))
	user, err := u.userServiceClient.DeleteByID(mdCtx, &userPb.DeleteByIDReq{
		Id: int64(userID),
	})
	if err != nil {
		u.logger.Error(
			"Got error from service",
			err,
			slog.String("request_id", requestID),
			slog.String("scope", scope),
		)
		return nil, fmt.Errorf("%s: %w", scope, ErrClientService.SetError(err))
	}
	return user, nil
}

func (u *userServiceUsecase) DeleteUserPermanentlyByID(ctx context.Context, userID int) (*userPb.DeletePermanentlyByIDResp, error) {
	const scope = "userServiceUsecase#DeleteUserPermanentlyByID"
	requestID := ctx.Value(util.RequestID).(string)
	mdCtx := metadata.NewOutgoingContext(ctx, metadata.Pairs("request-id", requestID))
	user, err := u.userServiceClient.DeletePermanentlyByID(mdCtx, &userPb.DeletePermanentlyByIDReq{
		Id: int64(userID),
	})
	if err != nil {
		u.logger.Error(
			"Got error from service",
			err,
			slog.String("request_id", requestID),
			slog.String("scope", scope),
		)
		return nil, fmt.Errorf("%s: %w", scope, ErrClientService.SetError(err))
	}
	return user, nil
}

func (u *userServiceUsecase) Register(ctx context.Context, userDto *req.UserDto) (*userPb.RegisterResp, error) {
	const scope = "userServiceUsecase#CreateUser"
	requestID := ctx.Value(util.RequestID).(string)
	mdCtx := metadata.NewOutgoingContext(ctx, metadata.Pairs("request-id", requestID))
	err := u.validate.Struct(userDto)
	if err != nil {
		u.logger.Error(
			"Failed to validate input",
			err,
			slog.String("request_id", requestID),
			slog.String("scope", scope),
		)
		validatorErrors := err.(validator.ValidationErrors)
		errorMessages := []string{}
		for _, validatorError := range validatorErrors {
			errorMessages = append(errorMessages, userDto.ErrorMessages(validatorError.Field(), validatorError.Tag()))
		}
		return nil, fmt.Errorf("%s: %w", scope, ErrFailToValidate.SetError(errors.New(strings.Join(errorMessages, ", "))))
	}
	user, err := u.userServiceClient.Register(mdCtx, &userPb.RegisterReq{
		Email:     userDto.Email,
		Password:  userDto.Password,
		FirstName: userDto.FirstName,
		LastName:  userDto.LastName,
	})
	if err != nil {
		u.logger.Error(
			"Got error from service",
			err,
			slog.String("request_id", requestID),
			slog.String("scope", scope),
		)
		return nil, fmt.Errorf("%s: %w", scope, ErrClientService.SetError(err))
	}
	return user, nil
}

func (u *userServiceUsecase) Login(ctx context.Context, loginDto *req.LoginDto) (*userPb.LoginResp, error) {
	const scope = "userServiceUsecase#Login"
	requestID := ctx.Value(util.RequestID).(string)
	mdCtx := metadata.NewOutgoingContext(ctx, metadata.Pairs("request-id", requestID))
	err := u.validate.Struct(loginDto)
	if err != nil {
		u.logger.Error(
			"Failed to validate input",
			err,
			slog.String("request_id", requestID),
			slog.String("scope", scope),
		)
		validatorErrors := err.(validator.ValidationErrors)
		errorMessages := []string{}
		for _, validatorError := range validatorErrors {
			errorMessages = append(errorMessages, loginDto.ErrorMessages(validatorError.Field(), validatorError.Tag()))
		}
		return nil, fmt.Errorf("%s: %w", scope, ErrFailToValidate.SetError(errors.New(strings.Join(errorMessages, ", "))))
	}
	user, err := u.userServiceClient.Login(mdCtx, &userPb.LoginReq{
		Email:    loginDto.Email,
		Password: loginDto.Password,
	})
	if err != nil {
		u.logger.Error(
			"Got error from service",
			err,
			slog.String("request_id", requestID),
			slog.String("scope", scope),
		)
		return nil, fmt.Errorf("%s: %w", scope, ErrClientService.SetError(err))
	}
	return user, nil
}
