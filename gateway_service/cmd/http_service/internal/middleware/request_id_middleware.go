package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (m Middleware) RequestID(ctx *gin.Context) {
	requestID := uuid.New().String()
	ctx.Set("request_id", requestID)
	ctx.Next()
}
