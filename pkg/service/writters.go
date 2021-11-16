package service

import (
	"encoding/json"
	"net/http"
)

// writeResponse writes a an error to response writer
func writeResponse(w http.ResponseWriter, statusCode int, response string) error {
	w.WriteHeader(statusCode)
	e := json.NewEncoder(w)
	e.SetIndent("", "    ")
	err := e.Encode(response)
	if err != nil {
		return err
	}
	return nil
}
