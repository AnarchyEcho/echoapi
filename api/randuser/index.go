package randuser

import (
	"echoapi/api/randuser/data"
	"net/http"
)

func Randusers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(data.Json()))
}
