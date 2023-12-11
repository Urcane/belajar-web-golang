package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Hello, Worlds!")
}

func TestHttpTest(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost/hello", nil)
	recorder := httptest.NewRecorder()

	HelloHandler(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}

func SayHelloWithParam(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")

	if name == "" {
		fmt.Fprint(writer, "Hello, Someone")
	} else {
		fmt.Fprintf(writer, "Hello %s, My name is Andr0idn", name)
	}
}

func TestHttpTestWithParam(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/ya-apalah-itu?name=Sofyan", nil)
	recorder := httptest.NewRecorder()

	SayHelloWithParam(recorder, request) // response parsed into recorder

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
