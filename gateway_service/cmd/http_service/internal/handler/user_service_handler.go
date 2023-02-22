package handler

import (
	"errors"
	"gatewayservice/cmd/http_service/internal"
	"gatewayservice/cmd/http_service/internal/util"
	"gatewayservice/internal/dto/req"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	userPb "github.com/ideaspaper/social-media-proto/user"
	"golang.org/x/exp/slog"
)

func (h Handler) FindUserByID(ctx *gin.Context) {
	const scope = "userHandler#FindUserByID"
	userID, err := strconv.Atoi(ctx.Param("userID"))
	if err != nil {
		h.logger.Error(
			"Bad userID request param",
			err,
			slog.String("request_id", ctx.Value("request_id").(string)),
			slog.String("scope", scope),
		)
		ctx.Error(&internal.ErrBadParams)
		return
	}
	user, err := h.userServiceUsecase.FindUserByID(ctx, &userPb.Req{
		RequestID: ctx.Value("request_id").(string),
		UserID:    int64(userID),
	})
	if err != nil {
		h.logger.Error(
			"Got error from usecase",
			err,
			slog.String("request_id", ctx.Value("request_id").(string)),
			slog.String("scope", scope),
		)
		ctx.Error(err)
		return
	}
	h.logger.Info(
		"Found a user by its ID",
		slog.String("request_id", ctx.Value("request_id").(string)),
		slog.String("scope", scope),
	)
	ctx.JSON(
		http.StatusOK,
		&util.StandardResponse{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
			Data:    user,
		},
	)
}

func (h Handler) CreateUser(ctx *gin.Context) {
	const scope = "userHandler#CreateUser"
	body := req.UserDto{}
	ctx.ShouldBind(&body)
	err := h.validate.Struct(body)
	if err != nil {
		h.logger.Error(
			"Failed to validate input",
			err,
			slog.String("request_id", ctx.Value("request_id").(string)),
			slog.String("scope", scope),
		)
		validatorErrors := err.(validator.ValidationErrors)
		errorMessages := []string{}
		for _, validatorError := range validatorErrors {
			errorMessages = append(errorMessages, body.ErrorMessages(validatorError.Field(), validatorError.Tag()))
		}
		ctx.Error(internal.ErrFailToValidate.SetError(errors.New(strings.Join(errorMessages, ", "))))
		return
	}
	user, err := h.userServiceUsecase.CreateUser(ctx, &userPb.Req{
		RequestID: ctx.Value("request_id").(string),
		UserReq: &userPb.UserReq{
			Email:     body.Email,
			Password:  body.Password,
			FirstName: body.FirstName,
			LastName:  body.LastName,
		},
	})
	if err != nil {
		h.logger.Error(
			"Got error from usecase",
			err,
			slog.String("request_id", ctx.Value("request_id").(string)),
			slog.String("scope", scope),
		)
		ctx.Error(err)
		return
	}
	h.logger.Info(
		"Created a user",
		slog.String("request_id", ctx.Value("request_id").(string)),
		slog.String("scope", scope),
	)
	ctx.JSON(
		http.StatusOK,
		&util.StandardResponse{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusCreated),
			Data:    user,
		},
	)
}

func (h Handler) DeleteUserByID(ctx *gin.Context) {
	const scope = "userHandler#DeleteUserByID"
	userID, err := strconv.Atoi(ctx.Param("userID"))
	if err != nil {
		h.logger.Error(
			"Bad userID request param",
			err,
			slog.String("request_id", ctx.Value("request_id").(string)),
			slog.String("scope", scope),
		)
		ctx.Error(&internal.ErrBadParams)
		return
	}
	user, err := h.userServiceUsecase.DeleteUserByID(ctx, &userPb.Req{
		RequestID: ctx.Value("request_id").(string),
		UserID:    int64(userID),
	})
	if err != nil {
		h.logger.Error(
			"Got error from usecase",
			err,
			slog.String("request_id", ctx.Value("request_id").(string)),
			slog.String("scope", scope),
		)
		ctx.Error(err)
		return
	}
	h.logger.Info(
		"Soft deleted a user by its ID",
		slog.String("request_id", ctx.Value("request_id").(string)),
		slog.String("scope", scope),
	)
	ctx.JSON(
		http.StatusOK,
		&util.StandardResponse{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
			Data:    user,
		},
	)
}

func (h Handler) DeleteUserPermanentlyByID(ctx *gin.Context) {
	const scope = "userHandler#DeleteUserPermanentlyByID"
	userID, err := strconv.Atoi(ctx.Param("userID"))
	if err != nil {
		h.logger.Error(
			"Bad userID request param",
			err,
			slog.String("request_id", ctx.Value("request_id").(string)),
			slog.String("scope", scope),
		)
		ctx.Error(&internal.ErrBadParams)
		return
	}
	user, err := h.userServiceUsecase.DeleteUserPermanentlyByID(ctx, &userPb.Req{
		RequestID: ctx.Value("request_id").(string),
		UserID:    int64(userID),
	})
	if err != nil {
		h.logger.Error(
			"Got error from usecase",
			err,
			slog.String("request_id", ctx.Value("request_id").(string)),
			slog.String("scope", scope),
		)
		ctx.Error(err)
		return
	}
	h.logger.Info(
		"Deleted a user permanently by its ID",
		slog.String("request_id", ctx.Value("request_id").(string)),
		slog.String("scope", scope),
	)
	ctx.JSON(
		http.StatusOK,
		&util.StandardResponse{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
			Data:    user,
		},
	)
}
