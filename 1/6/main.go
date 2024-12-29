package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math"
	"os"
	"slices"
)

var (
	englishPlaintextLetterFrequency = func() map[byte]float64 {
		text := `Hey there! So, I was thinking about what happened last weekend, and man, it was wild. Remember when we all decided to head out to the beach, even though the weather forecast said it might rain? Yeah, that turned out to be such a good call. The sky cleared up just as we got there, and the sunset? Absolutely stunning.

Anyway, I can't believe how much fun we had. Everyone was just so relaxed and happy—something about the sea air, I guess. Oh, and the barbecue? Top-notch. Whoever brought those marinated ribs, hats off to you. Seriously, though, I think I ate way too much, but no regrets.

Oh, do you remember that game we played? The one where we had to guess random facts about each other? I totally didn't see it coming when Sarah said she once went skydiving. Like, what? How did I not know that about her? She's always so quiet. Guess you learn something new every day, huh?

By the way, are we doing something similar next month? I heard someone mention a camping trip, but I wasn't sure if that was serious or just one of those 'we should totally do this' kind of ideas. Let me know if it’s happening—I’ll make sure to bring marshmallows and a bunch of snacks.

Okay, I should probably get going. Just wanted to check in and reminisce for a bit. Talk soon!`
		countsByCharacter := map[byte]uint16{}
		for i := 0; i < len(text); i++ {
			countsByCharacter[text[i]]++
		}
		frequenciesByCharacter := map[byte]float64{}
		for character, count := range countsByCharacter {
			frequenciesByCharacter[character] = float64(count) / float64(len(text))
		}
		return frequenciesByCharacter
	}()
)

type byteSequence interface {
	string | []byte
}

type KSandND struct {
	KeySize            uint8
	NormalizedDistance float64
}

func main() {
	originalMessage, err := os.ReadFile("challenge.txt")
	if err != nil {
		fmt.Println("error reading file:", err.Error())
	}
	hexEncryptedMessage := make([]byte, len(originalMessage))
	endIndex, err := base64.StdEncoding.Decode(hexEncryptedMessage, originalMessage)
	hexEncryptedMessage = hexEncryptedMessage[:endIndex]
	if err != nil {
		fmt.Println("error decoding base64:", err.Error())
	}
	encryptedMessage := hexEncryptedMessage
	normalizedDistancesByKeySize := []KSandND{}
	for keySizeGuess := uint8(2); keySizeGuess <= 40; keySizeGuess++ {
		distance := uint64(0)
		for i := 0; i < 4; i++ {
			_1stChunk := encryptedMessage[i*int(keySizeGuess) : (i+1)*int(keySizeGuess)]
			_2ndChunk := encryptedMessage[(i+1)*int(keySizeGuess) : (i+2)*int(keySizeGuess)]
			distance += hammingDistanceBetween(_1stChunk, _2ndChunk)
		}
		normalizedDistancesByKeySize = append(normalizedDistancesByKeySize, KSandND{
			KeySize:            keySizeGuess,
			NormalizedDistance: float64(distance) / float64(keySizeGuess),
		})
	}
	slices.SortFunc(normalizedDistancesByKeySize, func(a, b KSandND) int {
		switch {
		case a.NormalizedDistance < b.NormalizedDistance:
			return -1
		case a.NormalizedDistance > b.NormalizedDistance:
			return 1
		default:
			return 0
		}
	})
	minNormalizedMinDistanceSum := 1_000.0
	bestRepeatingKeyGuess := []byte{}
	for _, ksAndND := range normalizedDistancesByKeySize[:4] {
		repeatingKey := []byte{}
		minDistanceSum := 0.0
		for i := 0; i < int(ksAndND.KeySize); i++ {
			transposedBlockBuffer := []byte{}
			for j := 0; j+i < len(encryptedMessage); j += int(ksAndND.KeySize) {
				transposedBlockBuffer = append(transposedBlockBuffer, encryptedMessage[j+i])
			}
			key, _, minDistance, err := crackSingleCharXOREncryptionFromBytes(transposedBlockBuffer)
			if err != nil {
				fmt.Println("error cracking single char XOR encryption:", err.Error())
				return
			}
			repeatingKey = append(repeatingKey, key)
			minDistanceSum += minDistance
		}
		normalizedMinDistanceSum := minDistanceSum / float64(ksAndND.KeySize)
		if normalizedMinDistanceSum < minNormalizedMinDistanceSum {
			minNormalizedMinDistanceSum = normalizedMinDistanceSum
			bestRepeatingKeyGuess = repeatingKey
		}
	}
	fmt.Println()
	fmt.Println(repeatingKeyXORDecrypt(bestRepeatingKeyGuess, encryptedMessage))
	fmt.Println(`Key best guess: "` + string(bestRepeatingKeyGuess) + `"`)
}

func repeatingKeyXORDecrypt(key, encryptedMessage []byte) (decryptedMessage string) {
	decryptedBytes := []byte{}
	for i := 0; i < len(encryptedMessage); i++ {
		decryptedBytes = append(decryptedBytes, encryptedMessage[i]^key[i%len(key)])
	}
	return string(decryptedBytes)
}

func hammingDistanceBetween[T byteSequence](a, b T) (hammingDistance uint64) {
	for i := 0; i < len(a); i++ {
		x := a[i] ^ b[i]
		for j := 0; j < 8; j++ {
			if x%2 == 1 {
				hammingDistance++
			}
			x = x >> 1
		}
	}
	return
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

func englishPlaintextScore(hexEncodedEnglishPlaintext []byte) (distance float64) {
	countsByCharacter := map[byte]uint16{}
	for i := 0; i < len(hexEncodedEnglishPlaintext); i++ {
		countsByCharacter[hexEncodedEnglishPlaintext[i]]++
	}
	for character, count := range countsByCharacter {
		frequency := float64(count) / float64(len(hexEncodedEnglishPlaintext))
		distance += math.Abs(englishPlaintextLetterFrequency[character] - frequency)
	}
	return
}

func crackSingleCharXOREncryption(encryptedMessage string) (key byte, decryptedMessage string, minDistance float64, err error) {
	encryptedMessageHexBytes, err := hex.DecodeString(encryptedMessage)
	if err != nil {
		fmt.Println("error decoding encryptedMessage:", err.Error())
		return
	}
	return crackSingleCharXOREncryptionFromBytes(encryptedMessageHexBytes)
}

func crackSingleCharXOREncryptionFromBytes(encryptedMessageBytes []byte) (key byte, decryptedMessage string, minDistance float64, err error) {
	minDistance = 1000.0
	bestCandidate := []byte{}
	bestCandidateKey := byte(0)
	for key := byte(0); key < 128; key++ {
		fullLengthKey := []byte{}
		for i := 0; i < len(encryptedMessageBytes); i += 1 {
			fullLengthKey = append(fullLengthKey, key)
		}
		decryptedMessageHexBytes := bufferXOR(encryptedMessageBytes, fullLengthKey)
		score := englishPlaintextScore(decryptedMessageHexBytes)
		if score < minDistance {
			minDistance = score
			bestCandidate = decryptedMessageHexBytes
			bestCandidateKey = key
		}
	}
	decryptedMessage = string(bestCandidate)
	key = bestCandidateKey
	return
}
