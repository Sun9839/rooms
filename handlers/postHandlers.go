package handlers

import (
	"net/http"
	"fmt"

	"rooms-server.pet/midlwares"
)

func PostOwnData(w http.ResponseWriter, r *http.Request){
	if (r.Method == "POST") {
		fmt.Println("POSTOwnData")
	}
}

func Registration(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	midlwares.Registration(w, r)
}