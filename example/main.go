package main

import(
	"embed"
	"github.com/amalkochuparambilp/akpdeskgo/framework"
)

var assets embed.FS

func main() {
	app := framework.App{
		Assets: assets,
		Port: "8080",
		Path: "/app/index.html",
		Title: "Demo App",
		Width: 800,
		Height: 600,
		Debug: true,
	}
	app.Run()
}