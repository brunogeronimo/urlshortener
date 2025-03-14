package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"bruno.works/urlshortener/signaturevalidator"
	"bruno.works/urlshortener/types"
	"bruno.works/urlshortener/urlparser"
)

var (
	globalConfiguration types.Configuration
	Client              types.HTTPClient
)

func init() {
	Client = &http.Client{}
}

func downloadConfigurationFile() []byte {
	configFile, err := Client.Get(getConfigurationUrl())

	if err != nil {
		log.Fatalf("Error while downloading config file: %s", err)
	}

	defer configFile.Body.Close()

	body, _ := ioutil.ReadAll(configFile.Body)

	return body
}

func validateConfigSignature(downloadedConfigurationFile []byte) {
	publicKey := getPublicKey()
	if publicKey == "" {
		log.Fatalf("Security error! %s is set to true, but no public key has been defined", PublicKey)
	}
	signaturevalidator.IsSignatureValid(publicKey, downloadedConfigurationFile)
}

func handler(w http.ResponseWriter, r *http.Request) {
	destinationUrl := globalConfiguration.Urls[r.URL.Path]
	if destinationUrl == "" {
		http.Redirect(w, r, globalConfiguration.Fallback.Url, globalConfiguration.Fallback.RedirectCode)
		return
	}

	http.Redirect(w, r, destinationUrl, http.StatusPermanentRedirect)
}

func prepareEnv() {
	downloadedConfigurationFile := downloadConfigurationFile()

	if isCheckSignatureEnabled() {
		validateConfigSignature(downloadedConfigurationFile)
	}

	configuration, parseError := urlparser.ParseToConfig(downloadedConfigurationFile)
	if parseError != nil {
		log.Fatal(parseError)
	}

	globalConfiguration = configuration
}

func main() {
	prepareEnv()

	http.HandleFunc("/", handler)

	log.Println(http.ListenAndServe(fmt.Sprintf(":%s", getPort()), nil))
}
