package handler

import (
	"Go-Gin-LoadBalancer/internal/client"
	"Go-Gin-LoadBalancer/internal/config"
	"Go-Gin-LoadBalancer/internal/constants"
	"Go-Gin-LoadBalancer/internal/service"
	"Go-Gin-LoadBalancer/internal/service/strategy"
	"io"
	"log"
	"net/http"
	"strings"

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
		// Logic : if live servers are provided ------------------------------------------------
		if strings.ToLower(c.Query(constants.AreServersLive)) == "true" {
			log.Println("****** Live servers are provided, Now hitting actual backend server ******")
			resp, err := client.ProxyToServer(selectedServer, c.Request)
			if err != nil {
				c.JSON(http.StatusBadGateway, gin.H{"error": "Failed to reach backend server"})
				return
			}
			defer resp.Body.Close()
			// Copy response headers
			for k, v := range resp.Header {
				for _, vv := range v {
					c.Writer.Header().Add(k, vv)
				}
			}
			c.Writer.WriteHeader(resp.StatusCode)
			io.Copy(c.Writer, resp.Body)
			return
		}
		// If no live servers are provided, return the selected server
		c.JSON(http.StatusOK, selectedServer)
	}
}
