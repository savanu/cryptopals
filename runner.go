package main

import (
	"cryptopals/setone"
	"os"
	"strconv"
)

func main() {
	arg, _ := strconv.Atoi(os.Args[1])
	setone.Run(arg)
}
