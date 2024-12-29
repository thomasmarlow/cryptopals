package main

import (
	"encoding/hex"
	"strings"
)

func main() {}

func repeatingKeyXOREncrypt(key, message string) (encryptedMessage string) {
	lines := strings.Split(message, "\n")
	for _, line := range lines {
		encryptedMessage += repeatingKeyXOREncryptLine(key, line) + "\n"
	}
	return
}

func repeatingKeyXOREncryptLine(key, line string) (encryptedLine string) {
	for i := 0; i < len(line); i++ {
		encryptedLine += hex.EncodeToString([]byte{line[i] ^ (key[i%len(key)])})
	}
	return
}

func xorHexStrings(hex1, hex2 string) (hexR string, err error) {
	b1, err := hex.DecodeString(hex1)
	if err != nil {
		return
	}
	b2, err := hex.DecodeString(hex2)
	if err != nil {
		return
	}
	hexR = hex.EncodeToString(bufferXOR(b1, b2))
	return
}

func bufferXOR(b1, b2 []byte) (r []byte) {
	r = make([]byte, len(b1))
	for i, e1 := range b1 {
		r[i] = e1 ^ b2[i]
	}
	return
}
