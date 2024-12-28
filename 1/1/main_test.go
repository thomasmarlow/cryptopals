package main

import (
	"fmt"
	"testing"
)

func TestConvertHexToBase64(t *testing.T) {
	hex := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	expectedBase64 := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	obtainedBase64, err := convertHexToBase64(hex)
	if err != nil {
		t.Fatalf(`test failed: %s`, err.Error())
	}
	fmt.Println("expected value:", expectedBase64)
	fmt.Println("obtained value:", obtainedBase64)
	if expectedBase64 != obtainedBase64 {
		t.Fatalf(`test failed: resulting base64 mismatch`)
	}
}
