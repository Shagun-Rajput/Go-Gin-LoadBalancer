package router

import (
	"Go-Gin-LoadBalancer/internal/config"
	"Go-Gin-LoadBalancer/internal/handler"
	"log"
	"time"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, cfg config.ServerConfig) {
	r.Any("/*path", func(c *gin.Context) {
		start := time.Now()
		log.Printf("***** Request received: %s %s", c.Request.Method, c.Request.URL.Path)
		handler.HandleRequest(cfg)(c)
		elapsed := time.Since(start).Seconds() * 1000
		log.Printf("Response sent: %s %s | Duration: %.3f ms", c.Request.Method, c.Request.URL.Path, elapsed)
	})
}
