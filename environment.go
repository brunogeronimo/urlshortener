package main

import (
	"log"
	"os"
)

func getConfigurationUrl() string {
	url := os.Getenv("CONFIG_URL")

	if url == "" {
		log.Fatal("CONFIG_URL variable has not been provided")
	}

	return url
}

func getPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		return "8080"
	}

	return port
}