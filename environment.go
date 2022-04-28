package main

import (
	"log"
	"os"
)

const ConfigUrl = "CONFIG_URL"
const Port = "PORT"

func getConfigurationUrl() string {
	url := os.Getenv(ConfigUrl)

	if url == "" {
		log.Fatalf("%s variable has not been provided", ConfigUrl)
	}

	return url
}

func getPort() string {
	port := os.Getenv(Port)

	if port == "" {
		return "8080"
	}

	return port
}
