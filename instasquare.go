package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strings"

	"github.com/ohlrogge/instasquare/openImage"
)

func main() {
	// var path = "samples/sample.jpg"

	mydir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	files, err := ioutil.ReadDir(mydir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".jpeg") || strings.HasSuffix(file.Name(), ".jpg") {
			fmt.Println("Squaring " + file.Name())
			generateSquareImage(file.Name())
		}
	}
}

func generateSquareImage(path string) {

	img, _ := openImage.OpenJpeg(path)

	squareSideLength := getSquareSideLength(img.Bounds())

	// Generate
	squaredCanvas := image.NewRGBA(image.Rect(0, 0, squareSideLength, squareSideLength))

	// Colour canvas
	draw.Draw(squaredCanvas, squaredCanvas.Bounds(), &image.Uniform{color.RGBA{0, 0, 0, 255}}, image.Pt(0, 0), draw.Src)

	// Overlay image on the canvas
	draw.Draw(squaredCanvas, squaredCanvas.Bounds(), img, getStartingPoint(img), draw.Src)

	squareJpeg, _ := os.Create("square_" + path)
	jpeg.Encode(squareJpeg, squaredCanvas, &jpeg.Options{Quality: 90})
}

func getSquareSideLength(rectangle image.Rectangle) (squareSideLength int) {
	return int(math.Max(float64(rectangle.Dx()), float64(rectangle.Dy())))
}

func getStartingPoint(img image.Image) (startingPoint image.Point) {

	rectangle := img.Bounds()

	// Calculate translation
	translateBy := -int(math.Abs(math.RoundToEven(float64((rectangle.Dx() - rectangle.Dy()) / 2))))

	if rectangle.Dx() > rectangle.Dy() {
		return image.Pt(0, translateBy)
	} else {
		return image.Pt(translateBy, 0)
	}

}
