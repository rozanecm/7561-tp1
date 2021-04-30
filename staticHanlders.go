package main

import (
	"net/http"
	"os"
	"path/filepath"
)

func pdfHandler(writer http.ResponseWriter, request *http.Request) {
	cwd, _ := os.Getwd()
	http.ServeFile(writer, request, filepath.Join(cwd, "static/hear-tech-wsg-bridge-for-dante.pdf"))
}

func argerichHandler(writer http.ResponseWriter, request *http.Request) {
	cwd, _ := os.Getwd()
	http.ServeFile(writer, request, filepath.Join(cwd, "static/argerich.jpeg"))
}

func faviconHandler(writer http.ResponseWriter, request *http.Request) {
	cwd, _ := os.Getwd()
	http.ServeFile(writer, request, filepath.Join(cwd, "static/favicon.ico"))
}

func pdf2Handler(writer http.ResponseWriter, request *http.Request) {
	cwd, _ := os.Getwd()
	http.ServeFile(writer, request, filepath.Join(cwd, "static/Qu-16-User-Guide-AP9031_2.pdf"))
}

func staticHandler(writer http.ResponseWriter, request *http.Request) {
	cwd, _ := os.Getwd()
	http.ServeFile(writer, request, filepath.Join(cwd, "static/index.html"))
}
