package main

type ShortUrl struct {
	Url string `json:"url"`
	Destination string `json:"destination"`
}

type ConfigurationFile struct {
	FallbackUrl string `json:"fallbackUrl"`
	Urls []ShortUrl `json:"urls"`
}