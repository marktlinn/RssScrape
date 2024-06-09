package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// GetAPIKey extracts and returns the the API key associated with a user held in the HTTP headers.
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("failed to find Authorization Header")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		err := fmt.Sprintf("malformed Authorization Header: %s\n", val)
		return "", errors.New(err)
	}

	if vals[0] != "ApiKey" {
		err := fmt.Sprintf("malformed Authorization Header: %s\n", vals[0])
		return "", errors.New(err)
	}

	return vals[1], nil
}
