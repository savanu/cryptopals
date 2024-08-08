package main

import (
	"cryptopals/challenges"
	"os"
	"strconv"
)

func main() {
	arg, _ := strconv.Atoi(os.Args[1])
	challenges.Run(arg)
}
