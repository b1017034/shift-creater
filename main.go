package main

import (
	_ "embed"
	"shift-creater/backend/controller"

	"github.com/wailsapp/wails"
)

func basic() string {
	return "Hello World!"
}

//go:embed frontend/dist/app.js
var js string

//go:embed frontend/dist/app.css
var css string

func main() {

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "shift-creater",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})

	staffController := controller.NewStaffController()

	app.Bind(staffController)
	app.Bind(basic)
	app.Run()
}