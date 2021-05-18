package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

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

func downloadConfigFile() []byte {
	configFile, err := http.Get(getConfigUrl())

	if err != nil {
		log.Fatal(fmt.Sprintf("Error while downloading config file: %s", err))
	}

	defer configFile.Body.Close()

	body, _ := ioutil.ReadAll(configFile.Body)

	return body
}

func parseConfigFile() {
	parseErr := json.Unmarshal(downloadConfigFile(), &configurationFile)

	if parseErr != nil {
		log.Fatal("Error while parsing config file: ", parseErr)
	}
}

func getConfigUrl() string {
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

func main() {
	parseConfigFile()
	validateStructure()
	parseUrls()

	fallbackUrl = configurationFile.FallbackUrl

	http.HandleFunc("/", handler)

	log.Println(http.ListenAndServe(fmt.Sprintf(":%s", getPort()), nil))
}