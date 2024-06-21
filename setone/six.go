package setone

import (
	"cmp"
	"cryptopals/utils"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"slices"
)

func ChallengeSix() {
	encoded, err := os.ReadFile("files/challenge_6_data.txt")
	if err != nil {
		log.Fatalf("Couldn't open file: %s", err.Error())
	}

	data, err := base64.StdEncoding.DecodeString(string(encoded))
	if err != nil {
		log.Fatalf("Couldn't decode file: %s", err.Error())
	}

	keySize := findProbableKeySize(data)
	fmt.Println(keySize)
}

func findProbableKeySize(data []byte) int {

	type Pair struct {
		keySize  int
		distance float64
	}

	MAX_KEY_SIZE := 40

	distances := make([]Pair, 39)

	for i := 2; i <= MAX_KEY_SIZE; i++ {
		chunks := split(data, i)
		total := 0.0
		combo := 0.0
		for j := 0; j < len(chunks); j++ {
			for k := (j + 1); k < len(chunks); k++ {
				weight, _ := utils.HammingDistance(chunks[j], chunks[k])
				total += float64(weight) / float64(i)
				combo++
			}
		}

		distances[i-2] = Pair{i, total / combo}
	}

	slices.SortFunc(distances, func(a, b Pair) int {
		return cmp.Compare(a.distance, b.distance)
	})

	return distances[0].keySize
}

func split(data []byte, size int) [][]byte {
	chunks := make([][]byte, 4)
	chunks[0] = data[0:size]
	chunks[1] = data[size:(size * 2)]
	chunks[2] = data[(size * 2):(size * 3)]
	chunks[3] = data[(size * 3):(size * 4)]
	return chunks
}
