package request

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func GetURLParamInt(r *http.Request, key string) (int, error) {
	value, err := GetURLParamString(r, key)
	if err != nil {
		return 0, err
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}

	return intValue, nil
}

func GetURLParamString(r *http.Request, key string) (string, error) {
	value := chi.URLParam(r, key)

	if value == "" {
		return "", fmt.Errorf("missing %s parameter", key)
	}

	return value, nil
}

func GetURLParamBool(r *http.Request, key string) (bool, error) {
	value, err := GetURLParamString(r, key)
	if err != nil {
		return false, err
	}

	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		return false, err
	}

	return boolValue, nil
}
