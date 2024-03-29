package service

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

const (
	BYTE = 1 << (10 * iota)
	KILOBYTE
	MEGABYTE
)

func read(r *http.Request, req interface{}) error {
	// Protecting the API from too big uploads
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 10*MEGABYTE))
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &req)
	if err != nil {
		return err
	}
	return nil
}
