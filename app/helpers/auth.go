package helpers

import "github.com/go-resty/resty/v2"

var client *resty.Client

func InitAPI(baseURL string, headers map[string]string) {
	client = resty.New()
	if baseURL != "" {
		client.SetBaseURL(baseURL)
	}
	if headers != nil {
		client.SetHeaders(headers)
	}
}
