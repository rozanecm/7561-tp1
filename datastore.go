package main

import (
	"cloud.google.com/go/datastore"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Count struct {
	Count int
	Route string
}

func updateCounter(name, route string) {
	log.Println("in about stuff")
	ctx := context.Background()
	client, _ := datastore.NewClient(ctx, "taller3-tp1-rozanecm")
	defer client.Close()

	const amount = 1
	key := datastore.NameKey("page_visits_counter", name, nil)
	log.Println("key:", key)
	_, err := client.RunInTransaction(ctx, func(tx *datastore.Transaction) error {
		log.Println("in run in transaction")

		var task Count
		err := tx.Get(key, &task)
		if err != nil {
			log.Println("returning error:", err)
		}
		task.Count += amount
		task.Route = route
		log.Println(task.Route)
		_, err = tx.Put(key, &task)
		return err

	})
	if err != nil {
		log.Println("some error occurred:", err)
	}
	// [END datastore_transactional_retry]
	_ = err // Check error.
}

func counterHandler(writer http.ResponseWriter, request *http.Request) {

	counterName := mux.Vars(request)["counter"]

	ctx := context.Background()
	client, _ := datastore.NewClient(ctx, "taller3-tp1-rozanecm")
	defer client.Close()
	// [START datastore_lookup]
	var counter Count
	counterKey := datastore.NameKey("page_visits_counter", counterName, nil)
	err := client.Get(ctx, counterKey, &counter)
	// [END datastore_lookup]
	if err != nil {
		log.Println("Some error occurred retrieving counter:", err)
		fmt.Fprintf(writer, "Some error occurred retrieving counter.")
	} else {
		fmt.Fprintf(writer, strconv.Itoa(counter.Count))
	}
}
