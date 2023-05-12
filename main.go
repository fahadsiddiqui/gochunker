package main

import (
	"fmt"

	"github.com/fahadsiddiqui/gochunker/chunker"
)

func main() {
	chunks, err := chunker.Chunker("hey, what's up right now? what are you doing these days?", 16)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		for _, ch := range chunks {
			if ch != "" {
				fmt.Println(ch)
			}
		}
	}
}
