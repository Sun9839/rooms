package midlwares

import (
	"net/http"
)

type RegisterRequest struct {
    login     string `json:"login"`
    Password string `json:"password"`
	name string `json:"name"`
	surname string `json:"surname"`
	
}

func Registration(w http.ResponseWriter, r *http.Request) {
	
}