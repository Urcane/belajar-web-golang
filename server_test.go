package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestServerWithHandler(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "Hello Worlds")
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestServerWithServerMux(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "Halo this is root dir")
	})

	mux.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "hallo")
	})

	mux.HandleFunc("/images/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "Images")
	})

	mux.HandleFunc("/images/thumbnails/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "Thumbnails")
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestServerRequest() {
	mux := http.NewServeMux()

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
