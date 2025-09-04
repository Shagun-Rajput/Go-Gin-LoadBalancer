package service

import (
	"Go-Gin-LoadBalancer/internal/config"
	"Go-Gin-LoadBalancer/internal/service/strategy"
	"log"
)

func SelectServer(cfg config.ServerConfig) string {
	log.Println("***** Inside SelectServer : Proxying request based on strategy *****")
	var selectedServer any
	switch cfg.LoadBalancingStrategy {
	case "round_robin":
		log.Println("***** Strategy selected: Round Robin *****")
		selectedServer = strategy.RoundRobin(cfg.ServerList)
	case "least_connections":
		log.Println("***** Strategy selected: Least Connections *****")
		selectedServer = strategy.LeastConnections(cfg.ServerList)
	case "random":
		log.Println("***** Strategy selected: Random *****")
		selectedServer = strategy.Random(cfg.ServerList)
	default:
		log.Println("***** No Strategy selected: Using Default (Round Robin) *****")
		selectedServer = strategy.RoundRobin(cfg.ServerList)
	}
	log.Println("***** Selected Server: ", selectedServer)
	return selectedServer.(string)
}
