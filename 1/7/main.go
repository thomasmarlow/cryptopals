package main

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	originalMessage, err := os.ReadFile("challenge.txt")
	if err != nil {
		fmt.Println("error reading file:", err.Error())
	}
	encryptedMessage := make([]byte, len(originalMessage))
	endIndex, err := base64.RawStdEncoding.Decode(encryptedMessage, originalMessage)
	if err != nil {
		fmt.Println("error decoding base64:", err.Error())
	}
	encryptedMessage = encryptedMessage[:endIndex]
	cb, err := aes.NewCipher([]byte(`YELLOW SUBMARINE`))
	if err != nil {
		fmt.Println("error getting cipher:", err.Error())
		return
	}
	decryptedBytes := make([]byte, len(encryptedMessage))
	blockSize := 16
	for i := 0; i < len(encryptedMessage); i += blockSize {
		cb.Decrypt(decryptedBytes[i:i+blockSize], encryptedMessage[i:i+blockSize])
	}
	fmt.Println(string(decryptedBytes))
}
