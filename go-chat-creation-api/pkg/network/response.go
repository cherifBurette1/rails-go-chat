package network

import (
	"net/http"
	"encoding/json"
)

func Respond(w http.ResponseWriter, response []byte, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func RespondErr(w http.ResponseWriter, err error) {
	errorMsg := "Unable to serve this request: " + err.Error()
	responseBytes, _ := json.Marshal(struct {
		Message string `json:"message"`
	}{Message: errorMsg})

	Respond(w, responseBytes, http.StatusBadRequest)
}