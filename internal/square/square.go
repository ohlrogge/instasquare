package square

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"
	"math"
	"os"
	"path/filepath"

	"github.com/disintegration/imageorient"
)

type img struct {
	squarepath   string
	filetype     string
	orignalimage image.Image
}

func GenerateAll(directory string) {
	files, err := os.ReadDir(directory)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		ext := filepath.Ext(file.Name())
		if ext == ".jpeg" || ext == ".jpg" || ext == ".png" {
			i := img{
				squarepath:   filepath.Join(directory, "square_"+file.Name()),
				filetype:     filepath.Ext(file.Name()),
				orignalimage: openImage(filepath.Join(directory, file.Name())),
			}
			createsquareimage(i)
		}
	}
}

func createsquareimage(i img) {

	squareimg := generatesquareimage(i.orignalimage)
	generatenewfile(i, squareimg)
}

func generatenewfile(i img, squareimg *image.RGBA) {
	fmt.Println("Generating" + i.squarepath)
	f, err := os.Create(i.squarepath)

	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	if i.filetype == ".jpeg" || i.filetype == ".jpg" {
		jpeg.Encode(f, squareimg, &jpeg.Options{Quality: 90})
	} else if i.filetype == ".png" {
		png.Encode(f, squareimg)
	}
}

func generatesquareimage(img image.Image) *image.RGBA {

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

func openImage(path string) image.Image {

	f, err := os.Open(path)

	if isError(err) {
		log.Fatalf("os.Open failed: %v", err)
	}

	img, _, err := imageorient.Decode(f)

	if isError(err) {
		log.Fatalf("imageorient.Decode failed: %v", err)
	}

	return img
}
