package handler

import (
	"gatewayservice/cmd/http_service/internal"
	handlerUtil "gatewayservice/cmd/http_service/internal/util"
	"gatewayservice/internal/dto/req"
	internalUtil "gatewayservice/internal/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func (h Handler) FindUserByID(ctx *gin.Context) {
	const scope = "userHandler#FindUserByID"
	requestID := ctx.Value(internalUtil.RequestID).(string)
	userID, err := strconv.Atoi(ctx.Param("userID"))
	if err != nil {
		h.logger.Error(
			"Bad userID request param",
			err,
			slog.String("request_id", requestID),
			slog.String("scope", scope),
		)
		ctx.Error(&internal.ErrBadParams)
		return
	}
	user, err := h.userServiceUsecase.FindUserByID(ctx, userID)
	if err != nil {
		h.logger.Error(
			"Got error from usecase",
			err,
			slog.String("request_id", requestID),
			slog.String("scope", scope),
		)
		ctx.Error(err)
		return
	}
	h.logger.Info(
		"Found a user by its ID",
		slog.String("request_id", requestID),
		slog.String("scope", scope),
	)
	ctx.JSON(
		http.StatusOK,
		&handlerUtil.StandardResponse{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
			Data:    user,
		},
	)
}

func (h Handler) DeleteUserByID(ctx *gin.Context) {
	const scope = "userHandler#DeleteUserByID"
	requestID := ctx.Value(internalUtil.RequestID).(string)
	userID, err := strconv.Atoi(ctx.Param("userID"))
	if err != nil {
		h.logger.Error(
			"Bad userID request param",
			err,
			slog.String("request_id", requestID),
			slog.String("scope", scope),
		)
		ctx.Error(&internal.ErrBadParams)
		return
	}
	user, err := h.userServiceUsecase.DeleteUserByID(ctx, userID)
	if err != nil {
		h.logger.Error(
			"Got error from usecase",
			err,
			slog.String("request_id", requestID),
			slog.String("scope", scope),
		)
		ctx.Error(err)
		return
	}
	h.logger.Info(
		"Soft deleted a user by its ID",
		slog.String("request_id", requestID),
		slog.String("scope", scope),
	)
	ctx.JSON(
		http.StatusOK,
		&handlerUtil.StandardResponse{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
			Data:    user,
		},
	)
}

func (h Handler) DeleteUserPermanentlyByID(ctx *gin.Context) {
	const scope = "userHandler#DeleteUserPermanentlyByID"
	requestID := ctx.Value(internalUtil.RequestID).(string)
	userID, err := strconv.Atoi(ctx.Param("userID"))
	if err != nil {
		h.logger.Error(
			"Bad userID request param",
			err,
			slog.String("request_id", requestID),
			slog.String("scope", scope),
		)
		ctx.Error(&internal.ErrBadParams)
		return
	}
	user, err := h.userServiceUsecase.DeleteUserPermanentlyByID(ctx, userID)
	if err != nil {
		h.logger.Error(
			"Got error from usecase",
			err,
			slog.String("request_id", requestID),
			slog.String("scope", scope),
		)
		ctx.Error(err)
		return
	}
	h.logger.Info(
		"Deleted a user permanently by its ID",
		slog.String("request_id", requestID),
		slog.String("scope", scope),
	)
	ctx.JSON(
		http.StatusOK,
		&handlerUtil.StandardResponse{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
			Data:    user,
		},
	)
}

func (h Handler) RegisterUser(ctx *gin.Context) {
	const scope = "userHandler#RegisterUser"
	requestID := ctx.Value(internalUtil.RequestID).(string)
	userDto := req.UserDto{}
	ctx.ShouldBind(&userDto)
	user, err := h.userServiceUsecase.Register(ctx, &userDto)
	if err != nil {
		h.logger.Error(
			"Got error from usecase",
			err,
			slog.String("request_id", requestID),
			slog.String("scope", scope),
		)
		ctx.Error(err)
		return
	}
	h.logger.Info(
		"Created a user",
		slog.String("request_id", requestID),
		slog.String("scope", scope),
	)
	ctx.JSON(
		http.StatusCreated,
		&handlerUtil.StandardResponse{
			Code:    http.StatusCreated,
			Message: http.StatusText(http.StatusCreated),
			Data:    user,
		},
	)
}

func (h Handler) LoginUser(ctx *gin.Context) {
	const scope = "userHandler#LoginUser"
	requestID := ctx.Value(internalUtil.RequestID).(string)
	loginDto := req.LoginDto{}
	ctx.ShouldBind(&loginDto)
	user, err := h.userServiceUsecase.Login(ctx, &loginDto)
	if err != nil {
		h.logger.Error(
			"Got error from usecase",
			err,
			slog.String("request_id", requestID),
			slog.String("scope", scope),
		)
		ctx.Error(err)
		return
	}
	h.logger.Info(
		"Login success",
		slog.String("request_id", requestID),
		slog.String("scope", scope),
	)
	ctx.JSON(
		http.StatusOK,
		&handlerUtil.StandardResponse{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
			Data:    user,
		},
	)
}
