package types

type ShortUrl struct {
	Url         string `json:"url"`
	Destination string `json:"destination"`
}

type ConfigurationFile struct {
	FallbackUrl                 string     `json:"fallbackUrl"`
	IsFallbackPermanentRedirect bool       `json:"isFallbackPermanentRedirect"`
	Urls                        []ShortUrl `json:"urls"`
}

type Urls map[string]string

type Fallback struct {
	Url string
	RedirectCode int
}


type Configuration struct {
	Urls Urls
	Fallback Fallback
}