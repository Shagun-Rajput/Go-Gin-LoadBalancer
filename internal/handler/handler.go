package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"Go-Gin-LoadBalancer/internal/service"
)

func HandleRequest(c *gin.Context) {
	c.JSON(http.StatusOK, service.ProxyRequest())
}
