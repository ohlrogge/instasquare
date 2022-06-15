package square

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"log"
	"math"
	"os"
	"path/filepath"
	"strings"
)

func ConvertAllJpegsFromDirectory(directory string) {
	files, err := os.ReadDir(directory)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".jpeg") || strings.HasSuffix(file.Name(), ".jpg") {
			makeJpegSquare(filepath.Join(directory, file.Name()))
		}
	}
}

func makeJpegSquare(file string) {
	fmt.Println(file)

	img, err := openJpeg(file)

	if err != nil {
		fmt.Println(err)
	}

	squareImage := generateSquareImage(img)
	generateJPEG(file, squareImage)
}

func generateJPEG(fileName string, squaredCanvas *image.RGBA) {
	squareJpeg, err := os.Create(fileName + "_square.jpg")

	if err != nil {
		fmt.Println(err)
	}

	defer squareJpeg.Close()

	jpeg.Encode(squareJpeg, squaredCanvas, &jpeg.Options{Quality: 90})
}

func generateSquareImage(img image.Image) *image.RGBA {

	squareSideLength := getSquareSideLength(img.Bounds())

	// Generate
	squaredCanvas := image.NewRGBA(image.Rect(0, 0, squareSideLength, squareSideLength))

	// Colour canvas
	draw.Draw(squaredCanvas, squaredCanvas.Bounds(), &image.Uniform{color.RGBA{0, 0, 0, 255}}, image.Pt(0, 0), draw.Src)

	// Overlay image on the canvas
	draw.Draw(squaredCanvas, squaredCanvas.Bounds(), img, getStartingPoint(img), draw.Src)

	return squaredCanvas
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

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func openJpeg(path string) (image.Image, error) {

	var file, err = os.Open(path)

	if isError(err) {
		return nil, err
	}

	img, format, err := image.Decode(file)

	if isError(err) {
		return nil, err
	}

	if format != "jpeg" {
		fmt.Println("Image format is not jpeg.")
		return nil, errors.New("image format is not jpeg")
	}

	return img, nil
}
