package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Authorization: APIKey
func GetApiKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no Authorization header")
	}

	vals := strings.Split(val, " ")

	if len(vals) != 2 {
		return "", errors.New("malformed Authorization header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("invalid Authorization header")
	}

	return vals[1], nil
}
