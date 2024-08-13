package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"rooms-server.pet/midlwares"
)

func GetUserData(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	midlwares.GetUserData(w, r)
}

func GetRoom(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Room: %v\n", vars["room"])
}