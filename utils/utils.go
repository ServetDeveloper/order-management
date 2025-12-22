package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func ParseJson(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}

	if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
		return err
	}
	return nil
}

func WriteError(w http.ResponseWriter, status int, err error) {
	err1 := WriteJson(w, status, map[string]string{"err": err.Error()})
	if err1 != nil {
		log.Fatal(err1)
	}
}
