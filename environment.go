package main

import (
	"log"
	"os"
)

const ConfigUrl = "CONFIG_URL"
const Port = "PORT"
const CheckSignature = "CHECK_SIGNATURE"
const PublicKey = "PUBLIC_KEY"

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

func getPublicKey() string {
	return os.Getenv(PublicKey)
}

func isCheckSignatureEnabled() bool {
	getSignature := os.Getenv(CheckSignature)

	return getSignature == "true"
}
