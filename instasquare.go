package main

import (
	_ "image/png"

	"github.com/ohlrogge/instasquare/openImage/openImage"
)

var path = "samples/sample.jpg"

func main() {

	img, _ := openImage.OpenJpeg(path)

}
