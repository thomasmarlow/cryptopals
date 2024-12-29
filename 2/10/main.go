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
	encryptedBytes, err := base64.StdEncoding.DecodeString(string(originalMessage))
	if err != nil {
		fmt.Println("error decoding base64:", err.Error())
	}
	decryptedBytes := aesCBCDecrypt([]byte(`YELLOW SUBMARINE`), []byte{
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
	}, encryptedBytes)
	fmt.Println(string(decryptedBytes))
}

func aesCBCEncrypt(key, iv, plainBytes []byte) (encryptedBytes []byte) {
	cb, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(`error getting cipher:`, err.Error())
		return
	}
	prevEncryptedBlock := iv
	var plainBlock []byte
	for i := 0; i < len(plainBytes); i += len(key) {
		if i+len(key) >= len(plainBytes) {
			plainBlock = pkcs7(plainBytes[i:], uint32(len(key)))
		} else {
			plainBlock = plainBytes[i : i+len(key)]
		}
		x := bufferXOR(prevEncryptedBlock, plainBlock)
		encryptedBlock := make([]byte, len(key))
		cb.Encrypt(encryptedBlock, x)
		encryptedBytes = append(encryptedBytes, encryptedBlock...)
		prevEncryptedBlock = encryptedBlock
	}
	return
}

func aesCBCDecrypt(key, iv, encryptedBytes []byte) (decryptedBytes []byte) {
	cb, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(`error getting cipher:`, err.Error())
		return
	}
	prevEncryptedBlock := iv
	for i := 0; i < len(encryptedBytes); i += len(key) {
		encryptedBlock := encryptedBytes[i : i+len(key)]
		decryptedBlock := make([]byte, len(key))
		cb.Decrypt(decryptedBlock, encryptedBlock)
		decryptedBlock = bufferXOR(prevEncryptedBlock, decryptedBlock)
		decryptedBytes = append(decryptedBytes, decryptedBlock...)
		prevEncryptedBlock = encryptedBlock
	}
	decryptedBytes = trimPadding(decryptedBytes)
	return
}

func trimPadding(b []byte) (trimmedB []byte) {
	paddingStart := len(b)
	for i := 0; i < len(b); i++ {
		if b[i] == 4 && paddingStart == len(b) {
			paddingStart = i
		}
		if b[i] != 4 && paddingStart != len(b) {
			paddingStart = len(b)
		}
	}
	return b[:paddingStart]
}

func pkcs7(originalBuffer []byte, blockSize uint32) (paddedBuffer []byte) {
	paddedBuffer = originalBuffer[:]
	for i := 0; i < (int(blockSize)-len(originalBuffer)%int(blockSize))%int(blockSize); i++ {
		paddedBuffer = append(paddedBuffer, 4)
	}
	return
}

func bufferXOR(b1, b2 []byte) (r []byte) {
	r = make([]byte, len(b1))
	for i, e1 := range b1 {
		r[i] = e1 ^ b2[i]
	}
	return
}
