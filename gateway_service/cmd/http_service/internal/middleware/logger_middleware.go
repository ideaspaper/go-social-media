package middleware

import (
	"bytes"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (m Middleware) Logger(ctx *gin.Context) {
	start := time.Now()
	const scope = "middleware#Logger"
	ctx.Writer = &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
	ctx.Next()
	stop := time.Now()
	m.logger.Info(
		"Handle user request",
		slog.Any("request_id", ctx.Value("request_id")),
		slog.String("scope", scope),
		slog.String("ip", ctx.ClientIP()),
		slog.String("method", ctx.Request.Method),
		slog.String("path", ctx.Request.URL.Path),
		slog.Any("status_code", ctx.Writer.Status()),
		slog.String("latency", stop.Sub(start).String()),
	)
}
