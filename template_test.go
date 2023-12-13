package belajar_golang_web

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SimpleTemplateHTML(w http.ResponseWriter, r *http.Request)  {
	templateString := `<h1>Hello, {{.}}. Good Miawning!</h1>`

	// t, err := template.New("SIMPLE").Parse(templateString)
	// if err != nil {
	// 	panic(err)
	// }

	// template.Must() is an alternatif if you dont wanna code or include err handling manually
	t := template.Must(template.New("SIMPLE").Parse(templateString))

	t.ExecuteTemplate(w, "SIMPLE", "Mulyadi")
}

func TestSimpleTemplateHTML(t *testing.T) {
	handler := http.HandlerFunc(SimpleTemplateHTML)

	server := http.Server{
		Addr: "localhost:8080",
		Handler:  handler,
	}

	err := server.ListenAndServe()
	panic(err)
}

func TestSimpleTemplateHTMLFake(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	SimpleTemplateHTML(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func FileTemplate(w http.ResponseWriter, r *http.Request)  {
	templateFile := template.Must(template.ParseFiles("templates/index.gohtml"))

	templateFile.ExecuteTemplate(w, "index.gohtml", "Testing")
}

func TestFileTemplate(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	FileTemplate(recorder, request)

	response, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(response))
}

//go:embed templates/*
var templates embed.FS

func FileTemplateEmbed(w http.ResponseWriter, r *http.Request)  {
	templateFile := template.Must(template.ParseFS(templates, "templates/*.gohtml"))

	templateFile.ExecuteTemplate(w, "index.gohtml", "Embed Testing")
}

func TestFileTemplateEmbed(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	FileTemplate(recorder, request)

	response, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(response))
}

func DirectoryTemplate(w http.ResponseWriter, r *http.Request)  {
	t := template.Must(template.ParseGlob("./templates/*.gohtml"))

	t.ExecuteTemplate(w, "index.gohtml", "Hallow wewewe")
}

func TestDirectoryTemplate(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	DirectoryTemplate(recorder, request)

	responseBody, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(responseBody))
}

type Homepage struct {
	Title string
	Name string
}

func DataTemplate(w http.ResponseWriter, r *http.Request)  {
	t := template.Must(template.ParseFS(templates, "templates/*"))

	// When becoming api with multiple connection, CAREFULL on pointing struct directly, 
	// other connection might be able to change the content or structure of the struct. struct is muttable
	// its needed pointer at some case because the data isnt mutually needed for some module, 
	// *BUT ITS RISKY I THINK, PLEASE BE CAREFULL*
	//
	// data := &Homepage{
	// 	Title: "Belajar Golang 123",
	// 	Name: "Mantap",
	// }
	data := Homepage{
		Title: "Belajar Golang 123",
		Name: "Mantap",
	}

	// data selain menggunakan struct, bisa juga menggunakan map. type data yang di butuhkan sebenarnya adalah "interface", Oleh karena itu, parameter nya berbentuk dinamis
	t.ExecuteTemplate(w, "main.gohtml", &data)
}

func TestDataTemplate(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "https://localhost:8080/", nil)
	rec := httptest.NewRecorder()

	DataTemplate(rec, req)

	body, _ := io.ReadAll(rec.Result().Body)

	fmt.Println(string(body))
}