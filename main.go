package main

import (
	_ "embed"

	"github.com/dezhiShen/wails-study/pkg/tools"
	"github.com/wailsapp/wails"
)

//go:embed frontend/dist/app.js
var js string

//go:embed frontend/dist/app.css
var css string

func main() {
	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "wails-study",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})
	app.Bind(tools.RandomInt)
	app.Run()
}
