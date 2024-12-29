package main

import (
	"fmt"
	"testing"
)

func TestSingleCharXOREncryptionCracker(t *testing.T) {
	key, decryptedMessage, _, err := crackSingleCharXOREncryption("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	if err != nil {
		t.Fatalf(`test failed: %s`, err.Error())
	}
	if key != byte(88) {
		t.Fatalf(`test failed: wrong key`)
	}
	if decryptedMessage != "Cooking MC's like a pound of bacon" {
		t.Fatalf(`test failed: wrong decrypted message`)
	}
}

func TestXORHexStrings(t *testing.T) {
	hex1 := "1c0111001f010100061a024b53535009181c"
	hex2 := "686974207468652062756c6c277320657965"
	expectedXORResult := "746865206b696420646f6e277420706c6179"

	obtainedXORResult, err := xorHexStrings(hex1, hex2)
	if err != nil {
		t.Fatalf(`test failed: %s`, err.Error())
	}
	fmt.Println("expected value:", expectedXORResult)
	fmt.Println("obtained value:", obtainedXORResult)
	if expectedXORResult != obtainedXORResult {
		t.Fatalf(`test failed: resulting hex value mismatch`)
	}
}
