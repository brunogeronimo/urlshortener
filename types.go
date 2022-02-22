package main

type ShortUrl struct {
	Url         string `json:"url"`
	Destination string `json:"destination"`
}

type ConfigurationFile struct {
	FallbackUrl                 string     `json:"fallbackUrl"`
	IsFallbackPermanentRedirect bool       `json:"isFallbackPermanentRedirect"`
	Urls                        []ShortUrl `json:"urls"`
}
