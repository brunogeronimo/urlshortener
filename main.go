package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var configFilePath = "configs/urls.json"
var urls = make(map[string]string)
var fallbackUrl = ""
var configurationFile ConfigurationFile

func handler(w http.ResponseWriter, r *http.Request) {
	destinationUrl := urls[r.URL.Path]
	if destinationUrl == "" {
		http.Redirect(w, r, fallbackUrl, http.StatusTemporaryRedirect)
		return
	}

	http.Redirect(w, r, destinationUrl, http.StatusPermanentRedirect)
}

func validateStructure() {
	if configurationFile.FallbackUrl == "" {
		log.Fatal("fallbackUrl is mandatory on configuration file")
		os.Exit(1)
	}
}

func parseUrls() {
	for i, url := range configurationFile.Urls {
		if url.Url == "" {
			log.Println(`Url attribute is not set on url. Skipping...`, i)
			continue
		}

		if url.Destination == "" {
			log.Println(`Destination attribute is not set on url. Skipping`, i)
			continue
		}

		urls[url.Url] = url.Destination
	}
}

func main() {
	configs, _ := ioutil.ReadFile(configFilePath)
	err := json.Unmarshal(configs, &configurationFile)

	if err != nil {
		log.Fatal("Error while parsing config file", err)
		os.Exit(1)
	}

	validateStructure()
	parseUrls()

	fallbackUrl = configurationFile.FallbackUrl

	http.HandleFunc("/", handler)

	log.Println(http.ListenAndServe("", nil))
}