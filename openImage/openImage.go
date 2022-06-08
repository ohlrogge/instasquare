package openImage

import (
	"errors"
	"fmt"
	"image"
	_ "image/png"
	"os"
)

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func OpenJpeg(path string) (image.Image, error) {

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
