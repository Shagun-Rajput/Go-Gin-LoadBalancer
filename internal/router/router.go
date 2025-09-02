package router

import (
	"github.com/gin-gonic/gin"
	"Go-Gin-LoadBalancer/internal/handler"
)

func SetupRoutes(r *gin.Engine) {
	r.Any("/*path", handler.HandleRequest)
}
