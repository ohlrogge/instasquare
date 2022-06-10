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
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Widget")
	myWindow.Resize(fyne.NewSize(768, 460))

	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}

	text1 := canvas.NewText("Open folder", green)
	buttonFolder := widget.NewButton("Open Folder", func() {
		fmt.Println("Open folder tapped.")
		fmt.Println(openFolderDialog(myWindow))
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

		directory = strings.TrimPrefix(dir.String(), "file://")

		fmt.Println(directory)

		// jpegSquarer.ConvertAllJpegsFromDirectory(directory)

		dialog.ShowInformation("Folder Open", directory, win)
	}, win)

	return directory

}