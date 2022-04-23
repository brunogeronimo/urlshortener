package url_parser

import (
	"testing"
)

func TestParseToConfigWithEmptyFile(t *testing.T) {
	_, parseError := ParseToConfig([]byte(nil))

	if parseError == nil {
		t.Fatalf("Error should not be empty, %q found", parseError)
	}

	if parseError.Error() != "Error while parsing config file: unexpected end of JSON input" {
		t.Fatalf("Error response do not match, %q found", parseError)
	}
}

func TestParseToConfigWithInvalidJson(t *testing.T) {
	_, parseError := ParseToConfig([]byte("{not-a-json}"))
	if parseError == nil {
		t.Fatalf("Error should not be empty, %q found", parseError)
	}

	if parseError.Error() != "Error while parsing config file: invalid character 'n' looking for beginning of object key string" {
		t.Fatalf("Error response do not match, %q found", parseError)
	}
}

func TestParseToConfigWithoutFallbackUrl(t *testing.T) {
	_, parseError := ParseToConfig([]byte(`{"urls": []}`))
	if parseError == nil {
		t.Fatalf("Error should not be empty, %q found", parseError)
	}

	if parseError.Error() != "Error while validating config file: fallbackUrl is mandatory on configuration file" {
		t.Fatalf("Error response do not match, %q found", parseError)
	}
}

func TestParseToConfigWithInvalidUrls(t *testing.T) {
	configuration, parseError := ParseToConfig([]byte(`{"fallbackUrl": "https://fallback.url","urls": [{"url": "","destination": ""},{"url": "awesome-url","destination": ""},{"url": "","destination": "awesome-destination"},{"url": "/awesome-url","destination": "https://awesome.destination"}]}`))

	if parseError != nil {
		t.Fatalf("No errors were expected, found %s", parseError)
	}

	urlListSize := len(configuration.Urls)
	if urlListSize != 1 {
		t.Fatalf("Only one URL was expected, found %b", urlListSize)
	}

	originUrl := "/awesome-url"
	expectedDestinationUrl := "https://awesome.destination"
	destinationUrl := configuration.Urls[originUrl]

	if destinationUrl != expectedDestinationUrl {
		t.Fatalf("Found destination URL does not match with expected. Expected %s, found %s", expectedDestinationUrl, destinationUrl)
	}
}

func TestParseToConfigNotFallbackPermanentRedirect(t *testing.T) {
	testData := `{"fallbackUrl": "https://fallback.url", "isFallbackPermanentRedirect": false}`
	configuration, parseError := ParseToConfig([]byte(testData))

	if parseError != nil {
		t.Fatalf("No error expected, found %q", parseError)
	}

	if configuration.Fallback.RedirectCode != 307 {
		t.Fatalf("Invalid redirect code. Found %d", configuration.Fallback.RedirectCode)
	}
}

func TestParseToConfig(t *testing.T) {
	testData := `{"urls": [{"url": "/short-url","destination": "https://long-and-extense.url"}], "fallbackUrl": "https://fallback.url", "isFallbackPermanentRedirect": true}`
	configuration, parseError := ParseToConfig([]byte(testData))
	listSize := len(configuration.Urls)

	if parseError != nil {
		t.Fatalf("No error expected, found %q", parseError)
	}

	if configuration.Fallback.Url != "https://fallback.url" {
		t.Fatalf("Invalid fallback URL. Found %s", configuration.Fallback.Url)
	}

	if configuration.Fallback.RedirectCode != 308 {
		t.Fatalf("Invalid redirect code. Found %d", configuration.Fallback.RedirectCode)
	}

	if listSize != 1 {
		t.Fatalf("URL list size incorrect. Found %b", listSize)
	}

	originUrl := "/short-url"
	expectedDestinationUrl := "https://long-and-extense.url"
	destinationUrl := configuration.Urls[originUrl]
	if destinationUrl != expectedDestinationUrl {
		t.Fatalf("Destination URL should be %s, found %s", expectedDestinationUrl, destinationUrl)
	}
}
