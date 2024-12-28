package main

import "encoding/hex"

func bufferXOR(hex1, hex2 string) (hexR string, err error) {
	b1, err := hex.DecodeString(hex1)
	if err != nil {
		return
	}
	b2, err := hex.DecodeString(hex2)
	if err != nil {
		return
	}
	r := make([]byte, len(b1))
	for i, e1 := range b1 {
		r[i] = e1 ^ b2[i]
	}
	hexR = hex.EncodeToString(r)
	return
}
