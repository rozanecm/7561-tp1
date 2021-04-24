package main

import (
	"cloud.google.com/go/datastore"
	"context"
	"log"
)

type Count struct {
	Count int
}

func someDatastoreStuff() {
	log.Println("in about stuff")
	ctx := context.Background()
	client, _ := datastore.NewClient(ctx, "taller3-tp1-rozanecm")
	defer client.Close()

	const amount = 1
	key := datastore.NameKey("page_visits_counter", "home", nil)
	log.Println("key:", key)
	_, err := client.RunInTransaction(ctx, func(tx *datastore.Transaction) error {
		log.Println("in run in transaction")

		var task Count
		err := tx.Get(key, &task)
		if err != nil {
			log.Println("returning error:", err)
		}
		task.Count += amount
		_, err = tx.Put(key, &task)
		return err

	})
	if err != nil {
		log.Println("some error occurred:", err)
	}
	// [END datastore_transactional_retry]
	_ = err // Check error.
}
