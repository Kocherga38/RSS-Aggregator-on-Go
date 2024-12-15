package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts and API key from
// the headers of the HTTP request
func GetAPIKey(headers http.Header) (string, error) {
	value := headers.Get("Authorization")
	if value == "" {
		return "", errors.New("no authentication info found")
	}

	values := strings.Split(value, " ")
	if len(values) != 2 {
		return "", errors.New("malformed auth header")
	}

	if values[0] != "ApiKey" {
		return "", errors.New("malformed first part of auth header")
	}

	return values[1], nil
}
