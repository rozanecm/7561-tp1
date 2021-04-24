package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, "Hello, Home!")
}

func jobsHandler(writer http.ResponseWriter, request *http.Request) {
	task, err := createTask("my-queue")
	_, _ = fmt.Fprintf(writer, "Hello, Jobs!\n created task: %s, error: %s", task, err)
}

func aboutHandler(writer http.ResponseWriter, request *http.Request) {
	_, _ = fmt.Fprint(writer, "Hello, About!")
}

func aboutLegalsHandler(writer http.ResponseWriter, request *http.Request) {
	_, _ = fmt.Fprint(writer, "Hello, About Legals!")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/home", indexHandler)
	http.HandleFunc("/jobs", jobsHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/about/legals", aboutLegalsHandler)
	http.HandleFunc("/task_handler", taskHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
