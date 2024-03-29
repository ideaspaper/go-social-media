package middleware

import (
	"gatewayservice/internal/util"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (m Middleware) RequestID(ctx *gin.Context) {
	requestID := uuid.New().String()
	ctx.Set(util.RequestID, requestID)
	ctx.Next()
}
