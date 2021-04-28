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
	myRouter.HandleFunc("/home", templateHandler("home", "home-counter"))
	myRouter.HandleFunc("/", templateHandler("home", "home-counter"))
	myRouter.HandleFunc("/jobs", templateHandler("jobs", "jobs-counter"))
	myRouter.HandleFunc("/about", templateHandler("about", "about-counter"))
	myRouter.HandleFunc("/about/legals", templateHandler("about-legals", "about-legals-counter"))

	myRouter.HandleFunc("/task_handler", taskHandler)

	myRouter.HandleFunc("/static", staticHandler)
	myRouter.HandleFunc("/hear-tech-wsg-bridge-for-dante.pdf", pdfHandler)
	myRouter.HandleFunc("/Qu-16-User-Guide-AP9031_2.pdf", pdf2Handler)
	myRouter.HandleFunc("/argerich.jpeg", argerichHandler)
	myRouter.HandleFunc("/favicon.ico", faviconHandler)

	myRouter.HandleFunc("/template", templateHandler("some counter dude", ""))

	myRouter.HandleFunc("/api/counter/{counter}", counterHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Fatal(http.ListenAndServe(":"+port, myRouter))
}

func counterHandler(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	log.Println(params["counter"])
}

type CounterRequest struct {
	Counter string
}

func templateHandler(counter, taskName string) http.HandlerFunc {
	if taskName != "" {
		_, err := createTask(taskName)
		if err != nil {
			log.Println("error creating task:", err)
		}
	}

	return func(writer http.ResponseWriter, request *http.Request) {
		//counter := request.URL.Query().Get("counter")
		//log.Println("counter value:", counter)
		tmpl, err := template.ParseFiles("res/index.html")
		if err != nil {
			log.Println("some error ocurred loading teamplte:", err)
		}
		c := CounterRequest{counter}
		err = tmpl.Execute(writer, c)
		if err != nil {
			log.Println("error occurred Executing template:", err)
		}
	}
}
