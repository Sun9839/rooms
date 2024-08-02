package handlers

import (
	"net/http"
	"fmt"

	//"rooms-server.pet/midlwares"
)

func PostOwnData(w http.ResponseWriter, r *http.Request){
	if (r.Method == "POST") {
		fmt.Println("POSTOwnData")
	}
}