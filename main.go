package main

import (
	"bruno.works/urlshortener/types"
	"bruno.works/urlshortener/url-parser"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var fallbackUrl = ""
var redirectCode int
var urls types.Urls

func downloadConfigurationFile() []byte {
	configFile, err := http.Get(getConfigurationUrl())

	if err != nil {
		log.Fatal(fmt.Sprintf("Error while downloading config file: %s", err))
	}

	defer configFile.Body.Close()

	body, _ := ioutil.ReadAll(configFile.Body)

	return body
}

func handler(w http.ResponseWriter, r *http.Request) {
	destinationUrl := urls[r.URL.Path]
	if destinationUrl == "" {
		http.Redirect(w, r, fallbackUrl, redirectCode)
		return
	}

	http.Redirect(w, r, destinationUrl, http.StatusPermanentRedirect)
}

func main() {
	downloadedConfigurationFile := downloadConfigurationFile()
	var parseError error

	urls, fallbackUrl, redirectCode, parseError = url_parser.ConfigurationToObjects(downloadedConfigurationFile)
	if parseError != nil {
		log.Fatal(parseError)
	}

	http.HandleFunc("/", handler)

	log.Println(http.ListenAndServe(fmt.Sprintf(":%s", getPort()), nil))
}
