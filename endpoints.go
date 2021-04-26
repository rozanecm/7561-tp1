package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/home", homeHandler)
	myRouter.HandleFunc("/", homeHandler)
	myRouter.HandleFunc("/jobs", jobsHandler)
	myRouter.HandleFunc("/about", aboutHandler)
	myRouter.HandleFunc("/about/legals", aboutLegalsHandler)
	myRouter.HandleFunc("/task_handler", taskHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Fatal(http.ListenAndServe(":"+port, myRouter))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	task, err := createTask("home-counter")
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
