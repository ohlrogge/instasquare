package main

import (
	"fmt"
	"image/color"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/ohlrogge/instasquare/internal/square"
)

var imageDirectory string

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("instasquare")
	myWindow.Resize(fyne.NewSize(768, 460))

	white := color.NRGBA{R: 255, G: 255, B: 255, A: 255}

	text1 := canvas.NewText("To square all images in a folder, open it.", white)
	buttonFolder := widget.NewButton("Open Folder", func() {
		openFolderDialog(myWindow)
	})

	content := container.NewGridWithRows(2, text1, buttonFolder)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

func openFolderDialog(win fyne.Window) (directory string) {

	dialog.ShowFolderOpen(func(dir fyne.ListableURI, err error) {
		if err != nil {
			dialog.ShowError(err, win)
			return
		}
		if dir == nil {
			fmt.Println("Cancelled")
			return
		}

		if err != nil {
			dialog.ShowError(err, win)
			return
		}

		imageDirectory = strings.TrimPrefix(dir.String(), "file://")

		square.GenerateAll(imageDirectory)

		dialog.ShowInformation("Folder Open", "All images squared from "+imageDirectory, win)
	}, win)

	return imageDirectory

}
