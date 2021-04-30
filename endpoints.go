package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"os"
)

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/home", templateHandler("home-counter"))
	myRouter.HandleFunc("/", templateHandler("home-counter"))
	myRouter.HandleFunc("/jobs", templateHandler("jobs-counter"))
	myRouter.HandleFunc("/about", templateHandler("about-counter"))
	myRouter.HandleFunc("/about/legals", templateHandler("about-legals-counter"))

	myRouter.HandleFunc("/task_handler", taskHandler)

	myRouter.HandleFunc("/static", staticHandler)
	myRouter.HandleFunc("/hear-tech-wsg-bridge-for-dante.pdf", pdfHandler)
	myRouter.HandleFunc("/Qu-16-User-Guide-AP9031_2.pdf", pdf2Handler)
	myRouter.HandleFunc("/argerich.jpeg", argerichHandler)
	myRouter.HandleFunc("/favicon.ico", faviconHandler)

	myRouter.HandleFunc("/template", templateHandler(""))

	myRouter.HandleFunc("/api/counter/{counter}", counterHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Fatal(http.ListenAndServe(":"+port, myRouter))
}

type CounterRequest struct {
	Counter string
}

func templateHandler(taskName string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Println("taskName:", taskName)
		if taskName != "" {
			log.Println("about to create task:", taskName)
			task, err := createTask(taskName)
			if err != nil {
				log.Println("error creating task:", err)
			} else {
				log.Println("created task:", task)
			}
		}
		//counter := request.URL.Query().Get("counter")
		//log.Println("counter value:", counter)
		tmpl, err := template.ParseFiles("static/index.html")
		if err != nil {
			log.Println("some error ocurred loading teamplte:", err)
		}
		c := CounterRequest{taskName}
		err = tmpl.Execute(writer, c)
		if err != nil {
			log.Println("error occurred Executing template:", err)
		}
	}
}
