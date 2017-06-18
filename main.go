package main

import (
	"log"
	"net/http"
	"time"

	"github.com/boltdb/bolt"
	"github.com/iheanyi/go-vue-todo/api"
	"github.com/iheanyi/go-vue-todo/db"
)

func main() {
	bdb, err := MustOpenDatabase()
	if err != nil {
		log.Fatal("error opening db: %v", err)
	}
	defer bdb.Close()

	store := db.NewStore(bdb)
	svc := api.NewAPIService(store)
	r := api.NewRouter(svc)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:1415",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func MustOpenDatabase() (*bolt.DB, error) {
	bdb, err := bolt.Open("govue.db", 0600, nil)
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}

	// Initialize creation of the database
	tx, err := bdb.Begin(true)
	if err != nil {
		log.Fatalf("error beginning tx: %v", err)
	}
	defer tx.Rollback()

	_, err = tx.CreateBucketIfNotExists([]byte("todos"))
	if err != nil {
		log.Fatalf("error creating bucket: %v", err)
	}

	if err := tx.Commit(); err != nil {
		log.Fatalf("error commiting tx: %v", err)
	}
	return bdb, err
}
