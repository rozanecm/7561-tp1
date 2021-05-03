package main

import (
	"cloud.google.com/go/datastore"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"google.golang.org/api/iterator"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const SecondsCacheThreshold = 5

type Count struct {
	Count int
	Route string
}

type Cache struct {
	Counter   int
	timestamp time.Time
}

var cache map[string]Cache

func init() {
	cache = make(map[string]Cache)
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
	route := strings.Split(counterName, "-counter")[0]

	if cacheExpired(counterName) {
		log.Println("cache expired")
		ctx := context.Background()
		client, _ := datastore.NewClient(ctx, "taller3-tp1-rozanecm")
		defer client.Close()
		query := datastore.NewQuery("page_visits_counter").Filter("Route =", route)
		it := client.Run(ctx, query)
		totalCount := 0
		for {
			var currentCounter Count
			_, err := it.Next(&currentCounter)
			if err == iterator.Done {
				break
			}
			if err != nil {
				log.Fatalf("Error fetching next currentCounter: %v", err)
			}
			totalCount += currentCounter.Count
		}
		log.Println("Total Count", totalCount)
		cache[counterName] = Cache{Counter: totalCount, timestamp: time.Now()}
		fmt.Fprintf(writer, strconv.Itoa(totalCount))
	} else {
		log.Println("returning value from cache")
		fmt.Fprintf(writer, strconv.Itoa(cache[counterName].Counter))
	}
}

func cacheExpired(name string) bool {
	return time.Now().Sub(cache[name].timestamp).Seconds() > SecondsCacheThreshold
}
