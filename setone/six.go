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
	blocks := split(keySize, data)
	transposed := transpose(blocks)

	key := make([]byte, keySize)

	for i, v := range transposed {
		key[i] = utils.MostProbableKey(v)
	}

	fmt.Printf("Key: %s\n", key)

	result := make([]byte, len(data))

	for i := 0; i < len(data); i++ {
		result[i] = data[i] ^ key[i%keySize]
	}

	fmt.Printf("Result:\n\n%s\n", result)
}

// [[0 1 2 3] [4 5 6 7] [8 9 10 11] [12 13 14 15] [16 17 18 19]]
// [[0 4 8 12 16] [1 5 9 13 17] [2 6 10 14 18] [3 7 11 15 19]]
// Original dims: 5 arrays of 4 elems each
// Result dims: 4 arrays of 5 elems each
func transpose(blocks [][]byte) [][]byte {
	transposed := make([][]byte, len(blocks[0]))

	for i := 0; i < len(transposed); i++ {
		transposed[i] = make([]byte, len(blocks))
		inner := 0
		for _, v := range blocks {
			transposed[i][inner] = v[i]
			inner++
		}
	}

	return transposed
}

func split(size int, data []byte) [][]byte {
	blocks := make([][]byte, len(data)/size)

	for i := 0; i < len(blocks); i++ {
		beginIdx := (i * size)
		endIdx := min((i*size)+size, len(data))
		blocks[i] = data[beginIdx:endIdx]
	}

	return blocks
}

func findProbableKeySize(data []byte) int {
	type Pair struct {
		keySize  int
		distance float64
	}

	MAX_KEY_SIZE := 40
	distances := make([]Pair, 39)

	for i := 2; i <= MAX_KEY_SIZE; i++ {

		chunks := [][]byte{
			data[0:i],
			data[i:(i * 2)],
			data[(i * 2):(i * 3)],
			data[(i * 3):(i * 4)],
		}

		total := 0.0
		for j := 0; j < len(chunks); j++ {
			for k := (j + 1); k < len(chunks); k++ {
				weight, _ := utils.HammingDistance(chunks[j], chunks[k])
				total += float64(weight) / float64(i)
			}
		}

		distances[i-2] = Pair{i, total / 6.0}
	}

	slices.SortFunc(distances, func(a, b Pair) int {
		return cmp.Compare(a.distance, b.distance)
	})

	return distances[0].keySize
}
