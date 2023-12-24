package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	// Set up mux router
	router := mux.NewRouter()

	// Health API endpoint
	router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	// Create server
	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000", // Corrected the address format
		// Enforce timeout
		ReadTimeout:  15 * time.Second, // Corrected time unit
		WriteTimeout: 15 * time.Second, // Corrected time unit
	}

	// Log that the server is about to start
	log.Printf("Starting server on %s\n", srv.Addr)

	// Serve server
	log.Fatal(srv.ListenAndServe())
}
