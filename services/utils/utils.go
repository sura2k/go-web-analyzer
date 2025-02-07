package utils

import "net/url"

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
