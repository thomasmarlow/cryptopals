package main

import "testing"

func TestPKCS7(t *testing.T) {
	originalValue := `YELLOW SUBMARINE`
	paddedBytes := pkcs7([]byte(originalValue), 20)
	expectedValue := "YELLOW SUBMARINE\x04\x04\x04\x04"
	if string(paddedBytes) != expectedValue {
		t.Fatalf(`test failed:
expected: %v
got:      %v`, []byte(expectedValue), paddedBytes)
	}
}
