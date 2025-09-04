package router

import (
	"github.com/gin-gonic/gin"
	"Go-Gin-LoadBalancer/internal/handler"
	"Go-Gin-LoadBalancer/internal/config"
)

func SetupRoutes(r *gin.Engine, cfg config.ServerConfig) {
	r.Any("/*path", handler.HandleRequest(cfg))
}
