package main

import (
	"database/sql"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var db *sql.DB

func main() {
	info := "aall good"

	myApp := app.New()
	myWindow := myApp.NewWindow("t")
	content := initWindow()

	db = initdb()

	myWindow.SetContent(content)
	dialog.ShowInformation("Check 12", info, myWindow)

	myWindow.ShowAndRun()

}

func initWindow() *fyne.Container { // pass default content later
	label1 := "Set name"
	image := canvas.NewImageFromFile("/Users/peterlorimer/iconImages/original/happy.png")
	image.FillMode = canvas.ImageFillOriginal
	//image.SetMinSize(fyne.Size(200))
	imageCont := container.NewVBox(
		image,
		widget.NewLabel("36                                x"), //manual widen Box
	)

	content1 := container.NewVBox(
		container.NewHBox(
			widget.NewLabel("Icon Editor"),
			widget.NewLabel(label1),
		),
		container.NewHBox(
			container.NewVBox(
				widget.NewButton("DO something", func() {}),
				widget.NewLabel("Label 1"),
				widget.NewLabel("Label 2"),
				widget.NewLabel("Label 3"),
				widget.NewLabel("Label 4"),
			),
			imageCont,
		),
	)

	//      content1.Add(widget.NewButton("Add more items", func() {
	//              content1.Add(widget.NewLabel("Added"))
	//      }))
	return content1
}
