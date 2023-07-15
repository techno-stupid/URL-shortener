package helpers

import (
	"os"
	"strings"
)

func CheckDomainError(url string) bool {

	if url == os.Getenv("APP_URL") {
		return false
	}
	newURL := strings.Replace(url, "http://", "", 1)
	newURL = strings.Replace(newURL, "https://", "", 1)
	newURL = strings.Replace(newURL, "https://", "", 1)
	newURL = strings.Replace(newURL, "www.", "", 1)
	newURL = strings.Split(newURL, "/")[0]
	if newURL == os.Getenv("APP_URL") {
		return false
	}
	return true
}

func EnforceHTTP(url string) string {
	if url[:4] != "http" {
		return "http://" + url
	}
	return url
}
