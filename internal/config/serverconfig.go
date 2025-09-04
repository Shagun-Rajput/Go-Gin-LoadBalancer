package config

import (
	"log"
	"os"
	"strings"
)

type ServerConfig struct {
	AppPort               string
	ServerList            []string
	LoadBalancingStrategy string
}

func GetServerConfigFromEnv() ServerConfig {
	log.Println("***** Inside GetServerConfigFromEnv :Loading server configuration from environment variables...")
	serverListStr := os.Getenv("SERVER_LIST")
	var serverList []string
	if serverListStr != "" {
		// Split by comma and trim spaces
		rawList := splitAndTrim(serverListStr, ",")
		serverList = rawList
	}
	return ServerConfig{
		AppPort:               os.Getenv("APP_PORT"),
		ServerList:            serverList,
		LoadBalancingStrategy: os.Getenv("LOAD_BALANCING_STRATEGY"),
	}
}

func splitAndTrim(s, sep string) []string {
	var result []string
	for _, part := range strings.Split(s, sep) {
		trimmed := strings.TrimSpace(part)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}
