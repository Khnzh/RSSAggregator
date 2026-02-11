package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Authorization: ApiKey {"your_api_key"}
func GetApiKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authorization credentials")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of header")
	}
	return vals[1], nil
}
