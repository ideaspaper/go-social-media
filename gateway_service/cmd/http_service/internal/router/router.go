package router

import (
	"gatewayservice/cmd/http_service/internal/handler"
	"gatewayservice/cmd/http_service/internal/middleware"

	"github.com/gin-gonic/gin"
)

func New(h *handler.Handler, m *middleware.Middleware) *gin.Engine {
	r := gin.New()
	r.Use(m.Logger, m.ErrorHandler, m.CORSMiddleware, m.RequestID)
	r.GET("/users/:userID", h.FindUserByID)
	r.POST("/users", h.CreateUser)
	r.POST("/users/:userID/softdelete", h.DeleteUserByID)
	r.DELETE("/users/:userID", h.DeleteUserPermanentlyByID)
	r.NoRoute(h.NoRoute)
	return r
}
