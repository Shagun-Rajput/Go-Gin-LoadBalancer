package service

import (
	"Go-Gin-LoadBalancer/internal/config"
	"Go-Gin-LoadBalancer/internal/service/strategy"
)

func ProxyRequest(cfg config.ServerConfig) string {
	if cfg.LoadBalancingStrategy == "round_robin" {
		return strategy.RoundRobin(cfg.ServerList)
	}
	return "Unknown load balancing strategy"
}