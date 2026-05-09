package main

import (
	"bookstore/pkg/config"
	"bookstore/pkg/models"
	"bookstore/pkg/routes"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	config.Connect()

	if err := models.MigrateDB(); err != nil {
		log.Fatal("Migration failed:", err)
	}

	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)

	srv := &http.Server{
		Addr:         ":9010",
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Println("Server starting on :9010")
	log.Fatal(srv.ListenAndServe())
}
