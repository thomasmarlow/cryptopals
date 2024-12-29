package main

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestHammingDistanceBetweenStrings(t *testing.T) {
	distance := hammingDistanceBetween(
		`this is a test`,
		`wokka wokka!!!`,
	)
	if distance != 37 {
		t.Fatalf(`test failed:
expected: %d
got: %d`, 37, distance,
		)
	}
}

func TestRepeatingKeyXORDecryption(t *testing.T) {
	encryptedMessage := `0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f`
	expectedDecryptedMessage := `Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`
	hexEncryptedMessage, err := hex.DecodeString(encryptedMessage)
	if err != nil {
		t.Fatalf(`test failed: %s`, err.Error())
	}
	decryptedMessage := repeatingKeyXORDecrypt([]byte(`ICE`), hexEncryptedMessage)
	if decryptedMessage != expectedDecryptedMessage {
		t.Fatalf(`test failed:
expected: %s
got: %s`, expectedDecryptedMessage, decryptedMessage)
	}
}

func TestRepeatingKeyXOREncryption(t *testing.T) {
	message := `Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`
	expectedEncryptedMessage := `0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f`
	encryptedMessage := repeatingKeyXOREncrypt(`ICE`, message)
	if encryptedMessage != expectedEncryptedMessage {
		t.Fatalf(`test failed:
expected: %s
got: %s`, expectedEncryptedMessage, encryptedMessage)
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
	if expectedXORResult != obtainedXORResult {
		fmt.Println("expected value:", expectedXORResult)
		fmt.Println("obtained value:", obtainedXORResult)
		t.Fatalf(`test failed: resulting hex value mismatch`)
	}
}
