package common

import "net/http"

func AuthRequired(f http.HandlerFunc, errorResponse string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.FormValue("user") == "" {
			http.Error(w, errorResponse, http.StatusForbidden)
			return
		}
		f(w, r)
	}
}
