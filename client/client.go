package main

import (
	"log"

	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
)

var root *sciter.Element
var rootSelectorErr error
var w *window.Window
var windowErr error

func init() {
	rect := sciter.NewRect(100, 100, 400, 400)
	w, windowErr := window.New(sciter.SW_RESIZEABLE|sciter.SW_CONTROLS|sciter.SW_MAIN|sciter.SW_ENABLE_DEBUG, rect)
	if windowErr != nil {
		log.Fatal(err)
	}

	root, rootSelectorErr = w.GetRootElement()
	if rootSelectorErr != nil {
		log.Fatalln("Cannot select root element")
		return
	}
}

func main() {
	// log.Printf("handle: %v", w.Handle)
	w.LoadFile("./client.html")
	w.SetTitle("Go Chat")
	w.Show()
	w.Run()
}

func getChild(childName string) {
	childWindow, err := window.New(sciter.SW_CHILD, w.rect)
	if err != nil {
		log.Fatal(err)
	}

	childPath = "./screens/" + childName + "/" + childName + ".html"
	childWindow.LoadFile(childPath)

	return childWindow
}
