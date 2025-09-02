package main

import (
	"log"
	"Go-Gin-LoadBalancer/internal/router"
	"github.com/gin-gonic/gin"
	"runtime"
	"time"
)

func main() {
	r := gin.Default()

	// Initialize routes from internal/router
	router.SetupRoutes(r)

	// Start goroutine to print memory usage every minute
	go func() {
		for {
			var m runtime.MemStats
			runtime.ReadMemStats(&m)
			log.Printf("Memory Usage: Alloc = %v MiB, TotalAlloc = %v MiB, Sys = %v MiB, NumGC = %v", m.Alloc/1024/1024, m.TotalAlloc/1024/1024, m.Sys/1024/1024, m.NumGC)
			time.Sleep(time.Minute)
		}
	}()

	// Start server
	if err := r.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
	log.Println("Load Balancer started successfully")
}
