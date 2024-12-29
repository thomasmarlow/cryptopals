package main

import "testing"

func TestAESCBCEncryption(t *testing.T) {
	key := []byte(`thefloorislavaaa`)
	iv := []byte(`noitsnottttttttt`)
	originalMessage := []byte(`This is my message.`)
	recoveredMessage := aesCBCDecrypt(key, iv, aesCBCEncrypt(key, iv, originalMessage))
	if string(originalMessage) != string(recoveredMessage) {
		t.Fatalf(`test failed:
expected: %v
got:      %v`, []byte(originalMessage), recoveredMessage)
	}
}

func TestPKCS7(t *testing.T) {
	for _, testCase := range []struct {
		OriginalValue, ExpectedValue string
		KeySize                      uint32
	}{
		{
			OriginalValue: `YELLOW SUBMARINE`,
			KeySize:       20,
			ExpectedValue: "YELLOW SUBMARINE\x04\x04\x04\x04",
		},
		{
			OriginalValue: `MyValue`,
			KeySize:       8,
			ExpectedValue: "MyValue\x04",
		},
		{
			OriginalValue: `NoPaddingNeeded`,
			KeySize:       15,
			ExpectedValue: `NoPaddingNeeded`,
		},
		{
			OriginalValue: `aaaa`,
			KeySize:       3,
			ExpectedValue: "aaaa\x04\x04",
		},
	} {
		paddedBytes := pkcs7([]byte(testCase.OriginalValue), testCase.KeySize)
		if string(paddedBytes) != testCase.ExpectedValue {
			t.Fatalf(`test failed:
	expected: %v
	got:      %v`, []byte(testCase.ExpectedValue), paddedBytes)
		}
	}
}
