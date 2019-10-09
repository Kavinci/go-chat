package main

import "fyne.io/fyne/widget"

func main() {
	app := app.new()
	app.Icon()
	window := app.NewWindow("GoChat")
	window.SetContent(widget.NewVBox(
		widget.NewLabel("Sign in to GoChat"),
		widget.NewButton("Log in", func() {

		}),
		widget.NewButton("Sign up", func() {
			app.OpenURL()
		}),
		widget.NewButton("Cancel", func() {

		}),
	))
}
