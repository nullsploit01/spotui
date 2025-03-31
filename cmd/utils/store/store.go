package utils

import (
	"fmt"

	"github.com/zalando/go-keyring"
)

const service = "spotui"

func Set(key string, value string) error {
	return keyring.Set(service, key, value)
}

func Get(key string) (string, error) {
	value, err := keyring.Get(service, key)
	if err != nil {
		return "", fmt.Errorf("could not get value for key '%s': %w", key, err)
	}
	return value, nil
}

func Delete(key string) error {
	return keyring.Delete(service, key)
}
