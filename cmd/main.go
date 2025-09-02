package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Initialize routes from internal/router
	//router.SetupRoutes(r)

	// Start server
	if err := r.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
	log.Println("Load Balancer started successfully")
}
