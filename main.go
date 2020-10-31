package main

import (
	"HomeInventoryAPI/controllers"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var versionDbg string = "v0.0.3"

var startupTime time.Time

func main() {
	fmt.Println("Started main")

	setup(context.Background())
	router := mux.NewRouter()
	// router.Use(app.JwtAuthentication)

	header := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "FETCH", "DELETE", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	// Log when an appengine warmup request is used to create the new instance.
	// Warmup steps are taken in setup for consistency with "cold start" instances.
	router.HandleFunc("/_ah/warmup", func(w http.ResponseWriter, r *http.Request) {
		log.Println("warmup done")
	})
	router.HandleFunc("/", indexHandler)

	apiv1 := router.PathPrefix("/api/v1").Subrouter()
	registerRoutesForAPIV1(apiv1)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, handlers.CORS(header, methods, origins)(router)); err != nil {
		log.Fatal(err)
	}
}

func setup(ctx context.Context) error {
	startupTime = time.Now()
	return nil
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	uptime := time.Since(startupTime).Seconds()
	fmt.Fprintf(w, "Hello, World! Uptime: %.2fs\n%s", uptime, versionDbg)
}

func registerRoutesForAPIV1(api *mux.Router) {
	api.HandleFunc("/", indexV1Handler).Methods("GET")

	api.HandleFunc("/inventory/all", controllers.GetAllInventory).Methods("GET")
}

func indexV1Handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/api/v1/" {
		http.NotFound(w, r)
		return
	}
	uptime := time.Since(startupTime).Seconds()
	fmt.Fprintf(w, "API Version 1. Uptime: %.2fs\n%s", uptime, versionDbg)
}
