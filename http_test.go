package belajar_golang_web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
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

func MultipleParamWithSameValue(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	names := query["name"]

	fmt.Fprint(writer, names)
}

func TestHttpTestWithMultipleParam(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/test-multiple?name=mulyadi&name=okeoke", nil)
	recorder := httptest.NewRecorder()

	MultipleParamWithSameValue(recorder, request)
	response := recorder.Result()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}

func PostMethodNormal(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	firstName := r.PostForm.Get("firstName")
	lastName := r.PostForm.Get("lastName")

	fmt.Fprintf(w, "Hello, %s %s", firstName, lastName)
}

func TestPostMethodNormal(t *testing.T) {
	requestBody := strings.NewReader("firstName=Mulyadi&lastName=Okejuga")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	PostMethodNormal(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

// Sending with JSON Data
type Fullname struct {
	FirstName string
	LastName  string
}

func PostMethod(writer http.ResponseWriter, request *http.Request) {
	var data *Fullname

	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(writer, "Halo-Halo semuanya, Halo %s %s", data.FirstName, data.LastName)

	// firstName := request.PostForm.Get("first_name")
	// lastName := request.PostForm.Get("last_name")

	// fmt.Fprintf(writer, "Hello %s %s !", firstName, lastName)
}

func TestPostMethod(t *testing.T) {
	requestBody := bytes.NewBuffer([]byte(`{"firstName": "Mulyadi", "lastName": "Okeoke"}`))
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", requestBody)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	PostMethod(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}

func TestSetCookies(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/set-cookies?fullname=mulyadi", nil)
	recorder := httptest.NewRecorder()

	SetCookies(recorder, request)

	response := recorder.Result()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	cookies := response.Cookies()

	fmt.Println(string(body))

	for _, cookie := range cookies {
		fmt.Printf("Cookies dengan nama: %s, memiliki value %s\n", cookie.Name, cookie.Value)
	}

}

func TestGetCookies(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)

	cookie := new(http.Cookie)
	cookie.Name = "X-URCN-Name"
	cookie.Value = "Mulyade"	
	request.AddCookie(cookie)

	recorder := httptest.NewRecorder()

	GetCookies(recorder, request)

	response := recorder.Result()
	cookies := request.Cookies()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))

	for _, c := range cookies {
		fmt.Printf("Cookies dengan nama: %s, memiliki value %s\n", c.Name, c.Value)
	}
	
}
