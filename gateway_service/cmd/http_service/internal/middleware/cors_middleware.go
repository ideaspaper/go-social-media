package middleware

import (
	"gatewayservice/cmd/http_service/internal"

	"github.com/gin-gonic/gin"
)

func (m Middleware) CORSMiddleware(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
	if ctx.Request.Method == "OPTIONS" {
		ctx.Error(&internal.ErrCors)
		ctx.Abort()
		return
	}
	ctx.Next()
}
