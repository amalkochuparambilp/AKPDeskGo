package framework

import (
	"embed"
	"net/http"
	"github.com/webview/webview"
)

type App struct {
	Assets embed.FS
	Port string
	Path string
	Title string
	Width int
	Height int 
	Debug boot
}


func(a *App) Run() {
	go func() {
		http.Handle("/" http.FileServer(http.FS(a.Assets)))
		http.ListenAndServe(":"+a.Port, nil)
	}()
	webview.Open(
		a.Title,
		"http://localhost:"+a.Port+a.Path,
		a.Width,
		a.Height,
		a.Debug,
	)
}