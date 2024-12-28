package main

import (
	"fmt"
	"testing"
)

func TestBufferXOR(t *testing.T) {
	hex1 := "1c0111001f010100061a024b53535009181c"
	hex2 := "686974207468652062756c6c277320657965"
	expectedXORResult := "746865206b696420646f6e277420706c6179"

	obtainedXORResult, err := bufferXOR(hex1, hex2)
	if err != nil {
		t.Fatalf(`test failed: %s`, err.Error())
	}
	fmt.Println("expected value:", expectedXORResult)
	fmt.Println("obtained value:", obtainedXORResult)
	if expectedXORResult != obtainedXORResult {
		t.Fatalf(`test failed: resulting hex value mismatch`)
	}
}
