package common

import (
	"log"
	"net/http"
)

func HandleHTTPError(f errorHandler, errorResponse string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err != nil {
			log.Printf("%v", err)
			http.Error(w, errorResponse, http.StatusInternalServerError)
		}
	}
}
