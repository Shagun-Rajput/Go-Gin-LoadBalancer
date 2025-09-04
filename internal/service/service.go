package service

import (
	"Go-Gin-LoadBalancer/internal/config"
	"Go-Gin-LoadBalancer/internal/constants"
	"Go-Gin-LoadBalancer/internal/service/strategy"
	"log"
)

func SelectServer(cfg config.ServerConfig) string {
	log.Println("***** Inside SelectServer : Proxying request based on strategy *****")
	var selectedServer any
	switch cfg.LoadBalancingStrategy {
	case constants.StrategyRoundRobin:
		log.Println("***** Strategy selected: Round Robin *****")
		selectedServer = strategy.RoundRobin(cfg.ServerList)
	case constants.StrategyLeastConnections:
		log.Println("***** Strategy selected: Least Connections *****")
		selectedServer = strategy.LeastConnections(cfg.ServerList)
	case constants.StrategyRandom:
		log.Println("***** Strategy selected: Random *****")
		selectedServer = strategy.Random(cfg.ServerList)
	default:
		log.Println("***** No Strategy selected: Using Default (Round Robin) *****")
		selectedServer = strategy.RoundRobin(cfg.ServerList)
	}
	log.Println("***** Selected Server: ", selectedServer)
	return selectedServer.(string)
}
