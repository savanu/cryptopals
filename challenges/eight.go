package challenges

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func challengeEight() {

	file, err := os.Open("files/challenge_8_data.txt")
	if err != nil {
		log.Fatalf("Couldn't open file: %s", err.Error())
	}

	defer file.Close()

	i := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		freq := make(map[string]int)

		for i := 0; i < len(line); i += 32 {
			toScan := line[i : i+32]
			count := strings.Count(line, toScan)
			freq[toScan] = count
		}

		for k, v := range freq {
			if v > 1 {
				fmt.Printf("Line: %d Chunk: %s ; Count: %d\n", i, k, v)
			}
		}

		i++
	}
}
