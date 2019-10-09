package main

import (
	"log"
	"net/url"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	app := app.New()
	signupURL, err := url.Parse("https://kentutterback.com/GoChat/SignUp")
	if err != nil {
		log.Fatal(err)
	}

	app.Icon()
	window := app.NewWindow("GoChat")
	window.SetContent(widget.NewVBox(
		widget.NewLabel("Sign in to GoChat"),
		widget.NewButton("Log in", func() {

		}),
		widget.NewButton("Sign up", func() {
			app.OpenURL(signupURL)
		}),
		widget.NewButton("Cancel", func() {

		}),
	))
}
