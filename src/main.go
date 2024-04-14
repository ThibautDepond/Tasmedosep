package main

import (
	"tasmedosep/src/view"

	"fyne.io/fyne/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("Tasmedosep")

	w.SetContent(view.LockBuild())
	w.ShowAndRun()
}
