package config

import "os"

func GetServerPort() string {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080" // default port
	}
	return port
}
