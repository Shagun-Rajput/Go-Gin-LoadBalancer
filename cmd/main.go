package main

import (
	"Go-Gin-LoadBalancer/internal/config"
	"Go-Gin-LoadBalancer/internal/router"
	"log"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Load server config from environment
	serverConfig := config.GetServerConfigFromEnv()
	log.Printf("Loaded Config: AppPort=%s, ServerList=%v, LoadBalancingStrategy=%s", serverConfig.AppPort, serverConfig.ServerList, serverConfig.LoadBalancingStrategy)

	// Initialize routes from internal/router
	router.SetupRoutes(r)

	// Start goroutine to print memory usage every minute
	go func() {
		for {
			var m runtime.MemStats
			runtime.ReadMemStats(&m)
			log.Printf("***** Memory Usage: Alloc = %v MiB, TotalAlloc = %v MiB, Sys = %v MiB, NumGC = %v", m.Alloc/1024/1024, m.TotalAlloc/1024/1024, m.Sys/1024/1024, m.NumGC)
			time.Sleep(time.Minute)
		}
	}()

	// Start server
	if err := r.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
	log.Println("***** Load Balancer started successfully *****")
}
