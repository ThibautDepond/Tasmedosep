package view

import (
	"crypto/md5"

	"fyne.io/fyne"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"

	"tasmedosep/src/storage"
)

func LockBuild() *fyne.Container {
	keyInput := widget.NewEntry()
	keyInput.SetPlaceHolder("Enter Key ...")
	hello := widget.NewLabel("Welkam to Tasmedosep")

	return container.NewVBox(
		hello,
		keyInput,
		widget.NewButton("Hi!", func() {
			mdFiver := md5.New()
			storage.CurrentKey = mdFiver.Sum([]byte(keyInput.Text))
		}),
	)
}
