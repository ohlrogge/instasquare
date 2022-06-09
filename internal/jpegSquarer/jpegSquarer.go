package jpegSquarer

import (
	"fmt"
	"image"
	"image/jpeg"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/ohlrogge/instasquare/internal/openImage"
	"github.com/ohlrogge/instasquare/internal/squareImage"
)

func ConvertAllJpegsFromDirectory(directory string) {
	files, err := ioutil.ReadDir(directory)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".jpeg") || strings.HasSuffix(file.Name(), ".jpg") {
			makeJpegSquare(file)
		}
	}
}

func makeJpegSquare(file fs.FileInfo) {
	img, err := openImage.OpenJpeg(file.Name())

	if err != nil {
		fmt.Println(err)
	}

	squareImage := squareImage.Generate(img)
	generateJPEG(file.Name(), squareImage)
}

func generateJPEG(fileName string, squaredCanvas *image.RGBA) {
	squareJpeg, err := os.Create("square_" + fileName)

	if err != nil {
		fmt.Println(err)
	}

	defer squareJpeg.Close()

	jpeg.Encode(squareJpeg, squaredCanvas, &jpeg.Options{Quality: 90})
}
