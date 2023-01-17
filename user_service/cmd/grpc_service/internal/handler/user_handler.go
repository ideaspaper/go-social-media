package handler

import (
	"context"
	"userservice/cmd/grpc_service/internal/util"

	userPb "github.com/ideaspaper/social-media-proto/user"

	"golang.org/x/exp/slog"
)

func (h *Handler) FindByID(ctx context.Context, in *userPb.Req) (*userPb.Resp, error) {
	const scope = "userHandler#FindUserByID"
	user, err := h.userUsecase.FindByID(ctx, int(in.GetUserID()))
	if err != nil {
		h.logger.Error(
			"Got error from usecase",
			err,
			slog.String("request_id", in.GetRequestID()),
			slog.String("scope", scope),
		)
		return nil, err
	}
	h.logger.Info(
		"Found a user by its ID",
		slog.String("request_id", in.GetRequestID()),
		slog.String("scope", scope),
	)
	return &userPb.Resp{
		Message:  "Found a user by its ID",
		UserResp: util.RespUserDtoToPb(user),
	}, nil
}

func (h *Handler) Create(ctx context.Context, in *userPb.Req) (*userPb.Resp, error) {
	const scope = "userHandler#CreateUser"
	userDto := util.ReqUserPbToDto(in)
	user, err := h.userUsecase.Create(ctx, userDto)
	if err != nil {
		h.logger.Error(
			"Got error from usecase",
			err,
			slog.String("request_id", in.GetRequestID()),
			slog.String("scope", scope),
		)
		return nil, err
	}
	h.logger.Info(
		"Created a user",
		slog.String("request_id", in.GetRequestID()),
		slog.String("scope", scope),
	)
	return &userPb.Resp{
		Message:  "Created a user",
		UserResp: util.RespUserDtoToPb(user),
	}, nil
}

func (h *Handler) DeleteByID(ctx context.Context, in *userPb.Req) (*userPb.Resp, error) {
	const scope = "userHandler#DeleteUserByID"
	user, err := h.userUsecase.DeleteByID(ctx, int(in.GetUserID()))
	if err != nil {
		h.logger.Error(
			"Got error from usecase",
			err,
			slog.String("request_id", in.GetRequestID()),
			slog.String("scope", scope),
		)
		return nil, err
	}
	h.logger.Info(
		"Soft deleted a user by its ID",
		slog.String("request_id", in.GetRequestID()),
		slog.String("scope", scope),
	)
	return &userPb.Resp{
		Message:  "Soft deleted a user by its ID",
		UserResp: util.RespUserDtoToPb(user),
	}, nil
}

func (h *Handler) DeletePermanentlyByID(ctx context.Context, in *userPb.Req) (*userPb.Resp, error) {
	const scope = "userHandler#DeleteUserPermanentlyByID"
	user, err := h.userUsecase.DeletePermanentlyByID(ctx, int(in.GetUserID()))
	if err != nil {
		h.logger.Error(
			"Got error from usecase",
			err,
			slog.String("request_id", in.GetRequestID()),
			slog.String("scope", scope),
		)
		return nil, err
	}
	h.logger.Info(
		"Deleted a user permanently by its ID",
		slog.String("request_id", in.GetRequestID()),
		slog.String("scope", scope),
	)
	return &userPb.Resp{
		Message:  "Deleted a user permanently by its ID",
		UserResp: util.RespUserDtoToPb(user),
	}, nil
}
