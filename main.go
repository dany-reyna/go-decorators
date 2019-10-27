package main

import (
	"decorators/common"
	"fmt"
	"net/http"
)

func hiHandler(w http.ResponseWriter, r *http.Request) error {
	user := r.FormValue("user")

	if user == "martha" {
		return fmt.Errorf("Why did you say that name!")
	}

	_, err := fmt.Fprintln(w, "Hi, ", user)
	return err
}

var withError = common.HandleHTTPError(hiHandler, "An error occurred on hiHandler")
var withAuth = common.AuthRequired(withError, "Unauthorized!!!")

func main() {
	http.HandleFunc("/hi", withAuth)
	_ = http.ListenAndServe(":8080", nil)
}
