package belajar_golang_web

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

//go:embed templates/*
var templates2 embed.FS

var t = template.Must(template.ParseFS(templates2, "templates/*"))

func GetUploadFormHandler(w http.ResponseWriter, _ *http.Request)  {
	t.ExecuteTemplate(w, "upload_file", nil)
}

func PostUploadFormHandler(w http.ResponseWriter, r *http.Request)  {
	file, header, err := r.FormFile("uploaded-file")
	if err != nil {
		panic(err)
	}

	outputFilePath := HandleUploadFile(file, header, "./resources/")

	name := r.FormValue("name")

	t.ExecuteTemplate(w, "upload_file", map[string]interface{}{
		"Name": name,
		"FileSrc": outputFilePath,
	})
}

func HandleUploadFile(File multipart.File, Header *multipart.FileHeader, FilepathParent string) (string) {
	filepath := FilepathParent + Header.Filename

	// process to create the file 
	outputFile, err := os.Create(filepath)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(outputFile, File)

	if err != nil {
		panic(err)
	}

	return filepath
}

func TestUploadForm(t *testing.T) {
	mux := http.ServeMux{}

	mux.HandleFunc("/upload-form", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			PostUploadFormHandler(w,r)
			return
		}
		GetUploadFormHandler(w,r)
	})
	mux.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./resources"))))

	server := http.Server{
		Addr: "localhost:8080",
		Handler: &mux,
	}

	err := server.ListenAndServe()
	panic(err)
}


//go:embed resources/RichardOwen.jpeg
var richard []byte

func TestUploadFile(t *testing.T) {
	body := new(bytes.Buffer)

	writer := multipart.NewWriter(body)
	writer.WriteField("name", "Mulyadi")
	file, _ := writer.CreateFormFile("uploaded-file", "contohupload.jpeg")
	file.Write(richard)
	writer.Close()

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/upload-form", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	recorder := httptest.NewRecorder()

	PostUploadFormHandler(recorder, request)

	bodyResponse, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(bodyResponse))
}

