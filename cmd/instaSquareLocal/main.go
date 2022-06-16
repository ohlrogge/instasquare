package main

import (
	"fmt"
	"os"

	"github.com/ohlrogge/instasquare/internal/square"
)

func main() {

	directory, err := os.Getwd()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(directory)

	square.GenerateAll(directory)
}
