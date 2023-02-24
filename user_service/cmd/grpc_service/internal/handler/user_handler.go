package handler

import (
	"context"
	handlerUtil "userservice/cmd/grpc_service/internal/util"
	"userservice/internal/dto/req"
	internalUtil "userservice/internal/util"

	userPb "github.com/ideaspaper/social-media-proto/user"

	"golang.org/x/exp/slog"
)

func (h Handler) FindByID(ctx context.Context, in *userPb.FindByIDReq) (*userPb.FindByIDResp, error) {
	const scope = "userHandler#FindByID"
	requestID := ctx.Value(internalUtil.RequestID).(string)
	user, err := h.userUsecase.FindByID(ctx, int(in.GetId()))
	if err != nil {
		h.logger.Error(
			"Got error from usecase",
			err,
			slog.String("request_id", requestID),
			slog.String("scope", scope),
		)
		return nil, err
	}
	h.logger.Info(
		"Found a user by its ID",
		slog.String("request_id", requestID),
		slog.String("scope", scope),
	)
	return &userPb.FindByIDResp{
		Message:  "Found a user by its ID",
		UserResp: handlerUtil.RespUserDtoToPb(user),
	}, nil
}

func (h Handler) DeleteByID(ctx context.Context, in *userPb.DeleteByIDReq) (*userPb.DeleteByIDResp, error) {
	const scope = "userHandler#DeleteByID"
	requestID := ctx.Value(internalUtil.RequestID).(string)
	user, err := h.userUsecase.DeleteByID(ctx, int(in.GetId()))
	if err != nil {
		h.logger.Error(
			"Got error from usecase",
			err,
			slog.String("request_id", requestID),
			slog.String("scope", scope),
		)
		return nil, err
	}
	h.logger.Info(
		"Soft deleted a user by its ID",
		slog.String("request_id", requestID),
		slog.String("scope", scope),
	)
	return &userPb.DeleteByIDResp{
		Message:  "Soft deleted a user by its ID",
		UserResp: handlerUtil.RespUserDtoToPb(user),
	}, nil
}

func (h Handler) DeletePermanentlyByID(ctx context.Context, in *userPb.DeletePermanentlyByIDReq) (*userPb.DeletePermanentlyByIDResp, error) {
	const scope = "userHandler#DeletePermanentlyByID"
	requestID := ctx.Value(internalUtil.RequestID).(string)
	user, err := h.userUsecase.DeletePermanentlyByID(ctx, int(in.GetId()))
	if err != nil {
		h.logger.Error(
			"Got error from usecase",
			err,
			slog.String("request_id", requestID),
			slog.String("scope", scope),
		)
		return nil, err
	}
	h.logger.Info(
		"Deleted a user permanently by its ID",
		slog.String("request_id", requestID),
		slog.String("scope", scope),
	)
	return &userPb.DeletePermanentlyByIDResp{
		Message:  "Deleted a user permanently by its ID",
		UserResp: handlerUtil.RespUserDtoToPb(user),
	}, nil
}

func (h Handler) Register(ctx context.Context, in *userPb.RegisterReq) (*userPb.RegisterResp, error) {
	const scope = "userHandler#Register"
	requestID := ctx.Value(internalUtil.RequestID).(string)
	user, err := h.userUsecase.Register(ctx, &req.UserDto{
		Email:     in.GetEmail(),
		Password:  in.GetPassword(),
		FirstName: in.GetFirstName(),
		LastName:  in.GetLastName(),
	})
	if err != nil {
		h.logger.Error(
			"Got error from usecase",
			err,
			slog.String("request_id", requestID),
			slog.String("scope", scope),
		)
		return nil, err
	}
	h.logger.Info(
		"Registered a user",
		slog.String("request_id", requestID),
		slog.String("scope", scope),
	)
	return &userPb.RegisterResp{
		Message:  "Registered a user",
		UserResp: handlerUtil.RespUserDtoToPb(user),
	}, nil
}

func (h Handler) Login(ctx context.Context, in *userPb.LoginReq) (*userPb.LoginResp, error) {
	const scope = "userHandler#Login"
	requestID := ctx.Value(internalUtil.RequestID).(string)
	loginDto, err := h.userUsecase.Login(ctx, &req.LoginDto{
		Email:    in.GetEmail(),
		Password: in.GetPassword(),
	})
	if err != nil {
		h.logger.Error(
			"Got error from usecase",
			err,
			slog.String("request_id", requestID),
			slog.String("scope", scope),
		)
		return nil, err
	}
	return &userPb.LoginResp{
		Message: "User logged in",
		Id:      int64(loginDto.ID),
		Email:   loginDto.Email,
	}, nil
}
