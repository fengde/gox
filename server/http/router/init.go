package router

import (
	"server/http/handler"
	"server/http/middleware"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	r.Use(middleware.AccessLog())
	r.GET("/health", handler.Handler(handler.Health))
}
