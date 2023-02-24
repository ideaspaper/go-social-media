package usecase

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"userservice/internal/dto/req"
	"userservice/internal/dto/resp"
	"userservice/internal/repository"
	"userservice/internal/util"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/slog"
)

type userUsecase struct {
	logger         *slog.Logger
	validate       *validator.Validate
	userRepository repository.IUserRepository
}

func NewUserUsecase(logger *slog.Logger, validate *validator.Validate, userRepository repository.IUserRepository) IUserUsecase {
	return &userUsecase{
		logger:         logger,
		validate:       validate,
		userRepository: userRepository,
	}
}

func (uu userUsecase) FindByID(ctx context.Context, id int) (*resp.UserDto, error) {
	const scope = "userUsecase#FindByID"
	user, err := uu.userRepository.FindByID(ctx, id)
	if err != nil {
		uu.logger.Error(
			"Got error from repository",
			err,
			slog.String("request_id", ctx.Value(util.RequestID).(string)),
			slog.String("scope", scope),
		)
		if errors.Is(err, &repository.ErrDataNotFound) {
			return nil, fmt.Errorf("%s: %w", scope, ErrUserNotFound.SetError(err))
		}
		return nil, fmt.Errorf("%s: %w", scope, ErrUnknown.SetError(err))
	}
	uu.logger.Info(
		"Found a user by its ID",
		slog.String("request_id", ctx.Value(util.RequestID).(string)),
		slog.String("scope", scope),
	)
	return user.ToDto(), err
}

func (uu userUsecase) DeleteByID(ctx context.Context, id int) (*resp.UserDto, error) {
	const scope = "userUsecase#DeleteByID"
	user, err := uu.userRepository.DeleteByID(ctx, id)
	if err != nil {
		uu.logger.Error(
			"Got error from repository",
			err,
			slog.String("request_id", ctx.Value(util.RequestID).(string)),
			slog.String("scope", scope),
		)
		if errors.Is(err, &repository.ErrDataNotFound) {
			return nil, fmt.Errorf("%s: %w", scope, ErrUserNotFound.SetError(err))
		}
		return nil, fmt.Errorf("%s: %w", scope, ErrUnknown.SetError(err))
	}
	uu.logger.Info(
		"Soft deleted a user by its ID",
		slog.String("request_id", ctx.Value(util.RequestID).(string)),
		slog.String("scope", scope),
	)
	return user.ToDto(), err
}

func (uu userUsecase) DeletePermanentlyByID(ctx context.Context, id int) (*resp.UserDto, error) {
	const scope = "userUsecase#DeletePermanentlyByID"
	user, err := uu.userRepository.DeletePermanentlyByID(ctx, id)
	if err != nil {
		uu.logger.Error(
			"Got error from repository",
			err,
			slog.String("request_id", ctx.Value(util.RequestID).(string)),
			slog.String("scope", scope),
		)
		if errors.Is(err, &repository.ErrDataNotFound) {
			return nil, fmt.Errorf("%s: %w", scope, ErrUserNotFound.SetError(err))
		}
		return nil, fmt.Errorf("%s: %w", scope, ErrUnknown.SetError(err))
	}
	uu.logger.Info(
		"Deleted a user permanently by its ID",
		slog.String("request_id", ctx.Value(util.RequestID).(string)),
		slog.String("scope", scope),
	)
	return user.ToDto(), err
}

func (uu userUsecase) Register(ctx context.Context, userDto *req.UserDto) (*resp.UserDto, error) {
	const scope = "userUsecase#Register"
	err := uu.validate.Struct(userDto)
	if err != nil {
		uu.logger.Error(
			"Failed to validate input",
			err,
			slog.String("request_id", ctx.Value(util.RequestID).(string)),
			slog.String("scope", scope),
		)
		validatorErrors := err.(validator.ValidationErrors)
		errorMessages := []string{}
		for _, validatorError := range validatorErrors {
			errorMessages = append(errorMessages, userDto.ErrorMessages(validatorError.Field(), validatorError.Tag()))
		}
		return nil, fmt.Errorf("%s: %w", scope, ErrFailToValidate.SetError(errors.New(strings.Join(errorMessages, ", "))))
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userDto.Password), bcrypt.DefaultCost)
	if err != nil {
		uu.logger.Error(
			"Failed to hash password",
			err,
			slog.String("request_id", ctx.Value(util.RequestID).(string)),
			slog.String("scope", scope),
		)
		return nil, fmt.Errorf("%s: %w", scope, ErrFailHashingPassword.SetError(err))
	}
	userDto.Password = string(hashedPassword)
	user, err := uu.userRepository.Create(ctx, userDto)
	if err != nil {
		uu.logger.Error(
			"Got error from repository",
			err,
			slog.String("request_id", ctx.Value(util.RequestID).(string)),
			slog.String("scope", scope),
		)
		if errors.Is(err, &repository.ErrUniqueViolation) {
			return nil, fmt.Errorf("%s: %w", scope, ErrUserAlreadyExists.SetError(err))
		}
		return nil, fmt.Errorf("%s: %w", scope, ErrUnknown.SetError(err))
	}
	uu.logger.Info(
		"Created a user",
		slog.String("request_id", ctx.Value(util.RequestID).(string)),
		slog.String("scope", scope),
	)
	return user.ToDto(), err
}

func (uu userUsecase) Login(ctx context.Context, loginDto *req.LoginDto) (*resp.JwtDto, error) {
	const scope = "userUsecase#Login"
	err := uu.validate.Struct(loginDto)
	if err != nil {
		uu.logger.Error(
			"Failed to validate input",
			err,
			slog.String("request_id", ctx.Value(util.RequestID).(string)),
			slog.String("scope", scope),
		)
		validatorErrors := err.(validator.ValidationErrors)
		errorMessages := []string{}
		for _, validatorError := range validatorErrors {
			errorMessages = append(errorMessages, loginDto.ErrorMessages(validatorError.Field(), validatorError.Tag()))
		}
		return nil, fmt.Errorf("%s: %w", scope, ErrFailToValidate.SetError(errors.New(strings.Join(errorMessages, ", "))))
	}
	user, err := uu.userRepository.FindByEmail(ctx, loginDto.Email)
	if err != nil {
		uu.logger.Error(
			"Got error from repository",
			err,
			slog.String("request_id", ctx.Value(util.RequestID).(string)),
			slog.String("scope", scope),
		)
		if errors.Is(err, &repository.ErrDataNotFound) {
			return nil, fmt.Errorf("%s: %w", scope, ErrUserNotFound.SetError(err))
		}
		return nil, fmt.Errorf("%s: %w", scope, ErrUnknown.SetError(err))
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDto.Password))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", scope, ErrWrongEmailOrPassword.SetError(err))
	}
	ss, err := util.GenerateSignedJwt(user.ID, user.Email)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", scope, ErrFailSigningJWT.SetError(err))
	}
	return &resp.JwtDto{Token: ss}, nil
}
