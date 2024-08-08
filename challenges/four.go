package challenges

import (
	"bufio"
	"cryptopals/utils"
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

func challengeFour() {
	file, err := os.Open("files/challenge_4_data.txt")
	if err != nil {
		log.Fatalf("Couldn't open file: %s", err.Error())
	}

	defer file.Close()

	possibleMatches := make([][]byte, 0)

	scanner := bufio.NewScanner(file)
	var i = 1
	for scanner.Scan() {
		line := scanner.Text()
		decoded, _ := hex.DecodeString(line)
		probableKey := utils.MostProbableKey(decoded)
		possibleMatch := utils.XorByteArray(decoded, probableKey)
		possibleMatches = append(possibleMatches, possibleMatch)
		i += 1
	}

	highest := 0
	var match []byte
	for _, possibleMatch := range possibleMatches {
		count := utils.CommonLetterCount(possibleMatch)
		if count > highest {
			highest = count
			match = possibleMatch
		}
	}

	fmt.Printf("Result: %s", string(match))
}
