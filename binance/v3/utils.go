package v3

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

func signature(message, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(message))
	signingKey := fmt.Sprintf("%x", mac.Sum(nil))
	return signingKey
}

func parseModel(in interface{}) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	inrec, _ := json.Marshal(in)
	json.Unmarshal(inrec, &result)

	for field, val := range result {
		fmt.Println("KV Pair: ", field, val)
	}
	return result, nil
}
