package main

import (
	"fmt"
	"os"

	"github.com/ohlrogge/instasquare/internal/jpegSquarer"
)

func main() {

	directory, err := os.Getwd()

	if err != nil {
		fmt.Println(err)
	}

	jpegSquarer.ConvertAllJpegsFromDirectory(directory)
}
