package main

import (
	"encoding/hex"
	"fmt"
)

func xorStrings(str1, str2 string) string {
	bytes1 := []byte(str1)
	bytes2 := []byte(str2)

	if len(bytes1) != len(bytes2) {
		panic("length mismatch")
	}

	xorResult := make([]byte, len(bytes1))

	for i := range bytes1 {
		xorResult[i] = bytes1[i] ^ bytes2[i]
	}

	// Encode the XOR result to hexadecimal format
	encodedStr := hex.EncodeToString(xorResult)

	// Decode the hexadecimal bytes to string
	decodedStr, _ := hex.DecodeString(encodedStr)

	return string(decodedStr)
}

func main() {

	fmt.Println(xorStrings("A", "B"))
}
