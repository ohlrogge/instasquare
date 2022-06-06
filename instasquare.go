package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	_ "image/jpeg"
	"math"
	"os"

	"github.com/ohlrogge/instasquare/openImage"
)

var path = "samples/sample.jpg"

func main() {

	img, _ := openImage.OpenJpeg(path)

	fmt.Println(getSquareSideLength(img.Bounds()))

	squareSideLength := getSquareSideLength(img.Bounds())

	canvas := image.NewRGBA(image.Rect(0, 0, squareSideLength, squareSideLength))

	draw.Draw(canvas, canvas.Bounds(), &image.Uniform{color.RGBA{0, 0, 0, 255}}, image.Pt(0, 0), draw.Src)

	fmt.Println(getStartingPoint(img))

	draw.Draw(canvas, canvas.Bounds(), img, getStartingPoint(img), draw.Src)

	squareJpeg, _ := os.Create("samples/output.jpg")
	jpeg.Encode(squareJpeg, canvas, &jpeg.Options{Quality: 90})
}

func getSquareSideLength(rectangle image.Rectangle) (squareSideLength int) {
	if rectangle.Dx() > rectangle.Dy() {
		return rectangle.Dx()
	} else {
		return rectangle.Dy()
	}
}

func getStartingPoint(img image.Image) (startingPoint image.Point) {

	rectangle := img.Bounds()

	translateBy := -int(math.Abs(math.RoundToEven(float64((rectangle.Dx() - rectangle.Dy()) / 2))))

	if rectangle.Dx() > rectangle.Dy() {
		return image.Pt(0, translateBy)
	} else {
		return image.Pt(translateBy, 0)
	}

}
