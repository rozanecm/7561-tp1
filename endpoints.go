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
	myRouter.HandleFunc("/home", indexHandler)
	myRouter.HandleFunc("/", indexHandler)
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

func indexHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, "Hello, Home!")
}

func jobsHandler(writer http.ResponseWriter, request *http.Request) {
	task, err := createTask("my-queue")
	_, _ = fmt.Fprintf(writer, "Hello, Jobs!\n created task: %s, error: %s", task, err)
}

func aboutHandler(writer http.ResponseWriter, request *http.Request) {
	someDatastoreStuff("home")
	_, _ = fmt.Fprint(writer, "Hello, About!")
}

func aboutLegalsHandler(writer http.ResponseWriter, request *http.Request) {
	_, _ = fmt.Fprint(writer, "Hello, About Legals!")
}

