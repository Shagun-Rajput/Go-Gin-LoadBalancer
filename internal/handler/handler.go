package handler

import (
	"Go-Gin-LoadBalancer/internal/config"
	"Go-Gin-LoadBalancer/internal/constants"
	"Go-Gin-LoadBalancer/internal/service"
	"Go-Gin-LoadBalancer/internal/service/strategy"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleRequest(cfg config.ServerConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		selectedServer := service.SelectServer(cfg)
		if cfg.LoadBalancingStrategy == constants.StrategyLeastConnections {
			// Increment connection count
			strategy.IncConnection(selectedServer)
			// Decrement connection count when request is done
			defer strategy.DecConnection(selectedServer)
		}
		c.JSON(http.StatusOK, selectedServer)
	}
}
