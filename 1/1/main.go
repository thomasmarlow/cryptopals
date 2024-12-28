package main

import (
	"encoding/base64"
	"encoding/hex"
)

func convertHexToBase64(originalValue string) (convertedValue string, err error) {
	hexBytes, err := hex.DecodeString(originalValue)
	if err == nil {
		convertedValue = base64.RawStdEncoding.EncodeToString(hexBytes)
	}
	return
}
