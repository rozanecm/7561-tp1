package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func handleRequests() {

	//fs := http.FileServer(http.Dir("res"))

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/home", homeHandler)
	myRouter.HandleFunc("/", homeHandler)
	myRouter.HandleFunc("/jobs", jobsHandler)
	myRouter.HandleFunc("/about", aboutHandler)
	myRouter.HandleFunc("/about/legals", aboutLegalsHandler)
	myRouter.HandleFunc("/task_handler", taskHandler)

	myRouter.HandleFunc("/static", staticHandler)
	myRouter.HandleFunc("/hear-tech-wsg-bridge-for-dante.pdf", pdfHandler)
	myRouter.HandleFunc("/Qu-16-User-Guide-AP9031_2.pdf", pdf2Handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Fatal(http.ListenAndServe(":"+port, myRouter))
}

func pdfHandler(writer http.ResponseWriter, request *http.Request) {
	cwd, _ := os.Getwd()
	http.ServeFile(writer, request, filepath.Join(cwd, "res/hear-tech-wsg-bridge-for-dante.pdf"))
}

func pdf2Handler(writer http.ResponseWriter, request *http.Request) {
	cwd, _ := os.Getwd()
	http.ServeFile(writer, request, filepath.Join(cwd, "res/Qu-16-User-Guide-AP9031_2.pdf"))
}

func staticHandler(writer http.ResponseWriter, request *http.Request) {
	cwd, _ := os.Getwd()
	http.ServeFile(writer, request, filepath.Join(cwd, "res/index.html"))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("one log")
	task, err := createTask("home-counter")
	log.Println("one log after creating task")
	_, _ = fmt.Fprintf(w, "Hello, Home!\n created task: %s, error: %s", task, err)
}

func jobsHandler(writer http.ResponseWriter, request *http.Request) {
	task, err := createTask("jobs-counter")
	_, _ = fmt.Fprintf(writer, "Hello, Jobs!\n created task: %s, error: %s", task, err)
}

func aboutHandler(writer http.ResponseWriter, request *http.Request) {
	task, err := createTask("about-counter")
	_, _ = fmt.Fprintf(writer, "Hello, About!\n created task: %s, error: %s", task, err)
}

func aboutLegalsHandler(writer http.ResponseWriter, request *http.Request) {
	task, err := createTask("about-legals-counter")
	_, _ = fmt.Fprintf(writer, "Hello, About Legals!\n created task: %s, error: %s", task, err)
}
