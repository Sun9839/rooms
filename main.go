package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"rooms-server.pet/handlers"
)



func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.Firsthandler)

	srv := &http.Server{
		Handler: r,
		Addr: "localhost:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}
	fmt.Println("Server is listening..")

	log.Fatal(srv.ListenAndServe())
}