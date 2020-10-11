package service

import "net/http"

func (c Controller) healthHandler( w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte{})
}
