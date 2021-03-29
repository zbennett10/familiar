package main

import (
	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"

	"dungeon-comrade/api"
)

func main() {

	js := mewn.String("./frontend/dist/my-app/main.js")
	css := mewn.String("./frontend/dist/my-app/styles.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "dungeon-comrade",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})
	app.Bind(api.OnLogin)
	app.Run()

}
