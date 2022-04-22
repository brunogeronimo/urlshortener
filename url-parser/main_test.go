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
