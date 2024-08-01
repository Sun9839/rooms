package midlwares

import (
	"net/http"
	"encoding/json"
)

func FirstMidlware(w http.ResponseWriter){
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}