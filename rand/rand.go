package rand

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func Bytes(n int) ([]byte, error) {
	b := make([]byte, n)
	// rand.Read calls Reader.Read using io.ReadFull and creates random bytes slice
	nRead, err := rand.Read(b)
	if err != nil {
		return nil, fmt.Errorf("rand/bytes: %w", err)
	}
	if nRead < n {
		return nil, fmt.Errorf("rand/bytes: not enough random bytes read")
	}
	return b, nil
}

// creates random string using Bytes func
func String(n int) (string, error) {
	b, err := Bytes(n)
	if err != nil {
		return "", fmt.Errorf("rand/String: %w", err)
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
