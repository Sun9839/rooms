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
	r.HandleFunc("/userdata/{user}", handlers.GetUserData)
	r.HandleFunc("/roomdata/{room}", handlers.GetRoom)
	r.HandleFunc("/ownEnter", handlers.PostOwnData)

	srv := &http.Server{
		Handler: r,
		Addr: "localhost:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}
	fmt.Println("Server is listening..")

	log.Fatal(srv.ListenAndServe())
}