package handler

import (
	"gatewayservice/cmd/http_service/internal"

	"github.com/gin-gonic/gin"
)

func (h Handler) NoRoute(ctx *gin.Context) {
	ctx.Error(&internal.ErrNoRoute)
}
