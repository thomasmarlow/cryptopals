package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("challenge.txt")
	if err != nil {
		fmt.Println("error opening file:", err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	repeatedBlockCountPerLine := map[string]uint8{}
	for scanner.Scan() {
		line := []byte(scanner.Text())

		encryptedMessage := make([]byte, len(line)/2)
		_, err = hex.Decode(encryptedMessage, line)
		if err != nil {
			fmt.Println("error decoding hex:", err.Error())
		}
		equalBlockCount := map[string]uint8{}
		for i := 0; i < len(encryptedMessage); i += 16 {
			equalBlockCount[hex.EncodeToString(encryptedMessage[i:i+16])]++
		}
		repeatedBlockCount := uint8(0)
		for _, count := range equalBlockCount {
			repeatedBlockCount += count - 1
		}
		if repeatedBlockCount > 0 {
			repeatedBlockCountPerLine[hex.EncodeToString(encryptedMessage)] = repeatedBlockCount
		}
	}
	for line, repeatedBlockCount := range repeatedBlockCountPerLine {
		fmt.Println(repeatedBlockCount, line)
	}
}
