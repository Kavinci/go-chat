package main

import (
	"log"

	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
)

func main() {
	w, err := window.New(sciter.SW_RESIZEABLE|sciter.SW_CONTROLS|sciter.SW_MAIN, nil)
	if err != nil {
		log.Fatal(err)
	}
	// log.Printf("handle: %v", w.Handle)
	w.LoadFile("client.html")
	w.SetTitle("Go Chat")
	w.Show()
	w.Run()
}
