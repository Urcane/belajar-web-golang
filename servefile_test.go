package belajar_golang_web

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

func ServeFileHandle(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("name") != "" {
		http.ServeFile(writer, request, "./resources/success.html")
		return
	}
	http.ServeFile(writer, request, "./resources/index.html")
}

func TestServeFile(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", ServeFileHandle)

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	panic(err)
}

//go:embed resources/index.html
var resources1 string

//go:embed resources/success.html
var resources2 string

func ServeFileEmbedHandle(writer http.ResponseWriter, request *http.Request)  {
	if request.URL.Query().Get("name") != "" {
		fmt.Fprintf(writer, resources2)
		return
	}
	fmt.Fprintf(writer, resources1)
}

func TestServeFileEmbed(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", ServeFileEmbedHandle)

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	panic(err)
}