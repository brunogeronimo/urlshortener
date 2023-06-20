package types

import "net/http"

type HTTPClient interface {
	Get(url string) (resp *http.Response, err error)
}

type ShortUrl struct {
	Url         string `json:"url"`
	Destination string `json:"destination"`
}

type Signature struct {
	Sha256 string `json:"sha256"`
}

type ConfigurationFile struct {
	FallbackUrl                 string     `json:"fallbackUrl"`
	IsFallbackPermanentRedirect bool       `json:"isFallbackPermanentRedirect"`
	Urls                        []ShortUrl `json:"urls"`
	Signature                   Signature  `json:"signature"`
}

type Urls map[string]string

type Fallback struct {
	Url          string
	RedirectCode int
}

type Configuration struct {
	Urls     Urls
	Fallback Fallback
}
