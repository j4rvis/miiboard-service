package delivery

import "net/http"

// Health ...
func Health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
