package utils

import (
	"crypto/rand"
	"errors"
)

var (
	// Letters defines the available characters for passwords.
	Letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ012345678a9_-"

	// ErrRandomData defines an error for missing random data.
	ErrRandomData = errors.New("failed to read random data")
)

// Password generates a new random password.
func Password() (string, error) {
	b := make([]byte, 16)
	l, err := rand.Read(b)

	if err != nil {
		return "", err
	}

	if l != 16 {
		return "", ErrRandomData
	}

	for c := range b {
		b[c] = Letters[b[c]&0x3F]
	}

	return string(b), nil
}
