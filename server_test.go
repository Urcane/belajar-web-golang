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

func TestServerRequest(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, request.Method)
		fmt.Fprintln(writer, request.RequestURI)
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

func SetCookies(writer http.ResponseWriter, request *http.Request)  {
	if request.URL.Query().Get("fullname") == "" {
		fmt.Fprint(writer, "No Fullname has been set")
		return
	}
	
	cookies := new(http.Cookie)
	
	cookies.Name = "X-URCN-Name"
	cookies.Value = request.URL.Query().Get("fullname")
	cookies.Path = "/"

	http.SetCookie(writer, cookies)
	fmt.Fprint(writer, "Succeed to procceed cookies")
}

func GetCookies(writer http.ResponseWriter, request *http.Request)  {
	cookies, err := request.Cookie("X-URCN-Name")
	if err != nil {
		fmt.Fprint(writer, "No Cookies")
		return
	}
	name := cookies.Value
	fmt.Fprintf(writer, "Hello, %s", name)
}

func TestCookies(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/set-cookies", SetCookies)
	mux.HandleFunc("/get-cookies", GetCookies)

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	panic(err)
}
