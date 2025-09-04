package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"Go-Gin-LoadBalancer/internal/service"
	"Go-Gin-LoadBalancer/internal/config"
)

func HandleRequest(cfg config.ServerConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, service.ProxyRequest(cfg))
	}
}
