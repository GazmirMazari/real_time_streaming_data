package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"github.com/gorilla/mux"
)



func main() {
	//set up mux router 

	router := mux.NewRouter()

	//health api end point

	router.HandleFunc("/api/health", func(w*http.ResponseWriter, r http.Request){
		json.NewEncoder(w).Encode(map[string]bool){"ok": true}
	})

	//create server

	srv := &http.Server{
		Handler: router,
		Addr: "127.0.0.8000",
		//enforce timeout
		ReadTimeout: 15* time.Seconds
		WriteTimeout: 15* time.Seconds
	}

	//serve server
	log.Fatal(srv.ListenAndServe())

}

