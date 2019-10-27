package common

import "net/http"

type errorHandler func(http.ResponseWriter, *http.Request) error
