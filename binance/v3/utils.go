package v3

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

func signature(message, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(message))
	signingKey := fmt.Sprintf("%x", mac.Sum(nil))
	return signingKey
}
