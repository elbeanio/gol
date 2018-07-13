package io

import (
	"bufio"
	"log"
	"os"
)

// GetInput is a goroutine for non-blocking IO, terminal needs to be in raw mode
func GetInput(input chan rune) {
	for {
		in := bufio.NewReader(os.Stdin)
		result, _, err := in.ReadRune()
		if err != nil {
			log.Fatal(err)
		}
		input <- result
	}
}
