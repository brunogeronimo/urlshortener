package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

type MockClient struct {
	GetFunc func(url string) (resp *http.Response, err error)
}

var (
	GetFunc func(url string) (*http.Response, error)
)

func (client *MockClient) Get(url string) (*http.Response, error) {
	return GetFunc(url)
}

func init() {
	Client = &MockClient{}
}

func TestFileRetrievalAndDataParsing(t *testing.T) {
	fileUrl := "https://assets.bruno.works/mock.json"

	t.Setenv(ConfigUrl, fileUrl)

	GetFunc = func(url string) (*http.Response, error) {
		if url != fileUrl {
			t.Fatalf("Received url does not match. Expected %s, received %s", fileUrl, url)
		}

		r := ioutil.NopCloser(bytes.NewReader([]byte(`{
				"fallbackUrl": "https://fallback.url",
				"urls": [
					{
						"url": "/short-url",
						"destination": "https://long-and-extense.url"
					}
				],
				"isFallbackPermanentRedirect": true
			}`)))

		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	prepareEnv()
}
