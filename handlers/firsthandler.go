package handlers

import (
	"net/http"

	"rooms-server.pet/midlwares"
)

func Firsthandler(w http.ResponseWriter, r *http.Request){
	if (r.Method == "GET") {
		midlwares.FirstMidlware(w)
	}
}