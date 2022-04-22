package url_parser

import (
	"bruno.works/urlshortener/types"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

func getFallbackUrl(configurationFile types.ConfigurationFile) string {
	return configurationFile.FallbackUrl
}

func getFallbackRedirectCode(configurationFile types.ConfigurationFile) int {
	if configurationFile.IsFallbackPermanentRedirect {
		return http.StatusPermanentRedirect
	}

	return http.StatusTemporaryRedirect
}

func validateStructure(configurationFile types.ConfigurationFile) error {
	if configurationFile.FallbackUrl == "" {
		return errors.New("fallbackUrl is mandatory on configuration file")
	}

	return nil
}

func parseUrls(configurationFile types.ConfigurationFile) types.Urls {
	var urls = make(types.Urls)
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

	return urls
}

func parseConfigFile(downloadedConfigurationFile []byte) (types.ConfigurationFile, error) {
	var configurationFile types.ConfigurationFile
	parseErr := json.Unmarshal(downloadedConfigurationFile, &configurationFile)

	if parseErr != nil {
		return configurationFile, parseErr
	}

	return configurationFile, nil
}

func ConfigurationToObjects(downloadedConfigurationFile []byte) (types.Urls, string, int, error) {
	var urls = make(types.Urls)

	configurationFile, parseError := parseConfigFile(downloadedConfigurationFile)
	if parseError != nil {
		return urls, "", -1, errors.New(fmt.Sprintf("Error while parsing config file: %s", parseError))
	}

	validationError := validateStructure(configurationFile)
	if validationError != nil {
		return urls, "", -1, errors.New(fmt.Sprintf("Error while validating config file: %s", validationError))
	}

	urls = parseUrls(configurationFile)
	fallbackUrl := getFallbackUrl(configurationFile)
	redirectCode := getFallbackRedirectCode(configurationFile)

	return urls, fallbackUrl, redirectCode, nil
}