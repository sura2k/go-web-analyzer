package utils

import (
	"net/http"
	"net/url"
	"time"

	"github.com/sura2k/go-web-analyzer/config"
)

// Check whether the url is valid
func IsValidUrl(rawUrl string) (bool, error) {
	u, err := url.Parse(rawUrl)

	if err != nil {
		return false, err
	}

	return u.Scheme != "" && u.Host != "", nil
}

// Derive base url from the given url
func DeriveBaseUrl(rawUrl string) (string, error) {
	parsedUrl, err := url.Parse(rawUrl)
	if err != nil {
		return "", err
	}

	// Construct base url using scheme and host
	baseUrl := parsedUrl.Scheme + "://" + parsedUrl.Host
	return baseUrl, nil
}

// Derive host from the given url
func DeriveHost(rawUrl string) (string, error) {
	parsedUrl, err := url.Parse(rawUrl)
	if err != nil {
		return "", err
	}

	return parsedUrl.Host, nil
}

// Derive direct url for relative urls
func DeriveDirectUrl(relativeUrl string, baseUrl string) string {
	parsedBaseUrl, err := url.Parse(baseUrl)
	if err != nil {
		return ""
	}
	// Resolve relative url
	return parsedBaseUrl.ResolveReference(&url.URL{Path: relativeUrl}).String()
}

// Check if the url is accessible
// Assumptions:
//   - If the GET url returns with either 2xx or 3xx status codes, considered the url is accessible
func IsUrlAccessible(url string) bool {
	client := &http.Client{
		Timeout: time.Second * time.Duration(config.Config.Defaults.HTTP.Timeout.Seconds),
	}

	resp, err := client.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	// Check the response status code is 2xx or 3xx
	if resp.StatusCode/100 == 2 || resp.StatusCode/100 == 3 {
		return true
	} else {
		return false
	}
}
