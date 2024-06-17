package setone

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
)

// Encodes a given hex string as base64.
//
// Hex String -> Bytes -> Base 64 String
func ChallengeOne() {
	answer := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	payload := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	bytes, _ := hex.DecodeString(payload)
	encoded := base64.StdEncoding.EncodeToString(bytes)

	if answer != encoded {
		log.Fatalf("Failure. Expected: %s Actual: %s", answer, encoded)
	} else {
		fmt.Printf("Challenge One: %s\n", encoded)
	}
}
