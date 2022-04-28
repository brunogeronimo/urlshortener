package main

import (
	"bruno.works/urlshortener/types"
	"bruno.works/urlshortener/urlparser"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
		log.Fatal(fmt.Sprintf("Error while downloading config file: %s", err))
	}

	defer configFile.Body.Close()

	body, _ := ioutil.ReadAll(configFile.Body)

	return body
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
