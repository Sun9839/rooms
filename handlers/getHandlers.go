package handlers

import (
	"net/http"
	"fmt"

	//"rooms-server.pet/midlwares"
	
	"github.com/gorilla/mux"
)

func GetUserData(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User: %v\n", vars["user"])
}

func GetRoom(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Room: %v\n", vars["room"])
}