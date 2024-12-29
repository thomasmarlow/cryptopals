package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	fmt.Println()
	fmt.Println(repeatingKeyXOREncrypt("P3P3", "pepe.ricardo@gmail.com"))
	fmt.Println()
	fmt.Println(repeatingKeyXOREncrypt("myV3rYS3cr37k3y", "My very important message."))
	fmt.Println()
	fmt.Println(repeatingKeyXOREncrypt(
		"ajs0p9fjpas9edja8w9efjp9wejfjkawenfewuJWFKJSEJFuioWJEF8oijwef",
		`Hey there! So, I was thinking about what happened last weekend, and man, it was wild. Remember when we all decided to head out to the beach, even though the weather forecast said it might rain? Yeah, that turned out to be such a good call. The sky cleared up just as we got there, and the sunset? Absolutely stunning.

Anyway, I can't believe how much fun we had. Everyone was just so relaxed and happy—something about the sea air, I guess. Oh, and the barbecue? Top-notch. Whoever brought those marinated ribs, hats off to you. Seriously, though, I think I ate way too much, but no regrets.

Oh, do you remember that game we played? The one where we had to guess random facts about each other? I totally didn't see it coming when Sarah said she once went skydiving. Like, what? How did I not know that about her? She's always so quiet. Guess you learn something new every day, huh?

By the way, are we doing something similar next month? I heard someone mention a camping trip, but I wasn't sure if that was serious or just one of those 'we should totally do this' kind of ideas. Let me know if it’s happening—I’ll make sure to bring marshmallows and a bunch of snacks.

Okay, I should probably get going. Just wanted to check in and reminisce for a bit. Talk soon!`,
	))
	fmt.Println()
}

func repeatingKeyXOREncrypt(key, message string) (encryptedMessage string) {
	for i := 0; i < len(message); i++ {
		encryptedMessage += hex.EncodeToString([]byte{message[i] ^ (key[i%len(key)])})
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
