package web

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed out/*
var staticFiles embed.FS

var staticDir = "out"

func Handler() http.Handler {
	fSys, err := fs.Sub(staticFiles, "out")
	if err != nil {
		panic("Failed to create sub filesystem: " + err.Error())
	}
	return http.FileServer(http.FS(fSys))
}
