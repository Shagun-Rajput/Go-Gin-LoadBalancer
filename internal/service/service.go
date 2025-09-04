package service

import (
	"Go-Gin-LoadBalancer/internal/config"
	"Go-Gin-LoadBalancer/internal/service/strategy"
	"log"
)

func ProxyRequest(cfg config.ServerConfig) string {
	log.Println("***** Inside ProxyRequest : Proxying request based on strategy *****")
	switch cfg.LoadBalancingStrategy {
	case "round_robin":
		log.Println("***** Strategy selected: Round Robin *****")
		return strategy.RoundRobin(cfg.ServerList)
	case "least_connections":
		log.Println("***** Strategy selected: Least Connections *****")
		return strategy.LeastConnections(cfg.ServerList)
	case "random":
		log.Println("***** Strategy selected: Random *****")
		return strategy.Random(cfg.ServerList)
	case "weighted_round_robin":
		log.Println("***** Strategy selected: Weighted Round Robin *****")
		return strategy.WeightedRoundRobin(cfg.ServerList)
	default:
		log.Println("***** No Strategy selected: Using Default (Round Robin) *****")
		return strategy.RoundRobin(cfg.ServerList)
	}
}
