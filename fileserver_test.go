package belajar_golang_web

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)

//go:embed resources
var resources embed.FS

func TestFileServer(t *testing.T) {
	dir, _ := fs.Sub(resources, "resources")
	fileserver := http.FileServer(http.FS(dir))

	mux := http.NewServeMux()

	mux.Handle("/filesystem/", http.StripPrefix("/filesystem/", fileserver))

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	panic(err)
}