package v3

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
)

func signature(message, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(message))
	signingKey := fmt.Sprintf("%x", mac.Sum(nil))
	return signingKey
}

func parseModel(in interface{}) (map[string]string, error) {
	result := make(map[string]string)
	inrec, _ := json.Marshal(in)
	json.Unmarshal(inrec, &result)

	for field, val := range result {
		fmt.Println("KV Pair: ", field, val)
	}
	return result, nil
}

func structToParams(v interface{}) url.Values {
	vals := url.Values{}
	vValue := reflect.ValueOf(v)
	vType := vValue.Type()
	for i := 0; i < vValue.NumField(); i++ {
		key := vType.Field(i).Name
		value := vValue.Field(i).Interface()
		vals.Add(key, fmt.Sprintf("%v", value))
	}
	return vals
}
