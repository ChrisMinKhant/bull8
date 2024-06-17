package filter

import "net/http"

var ErrorSigal = make(chan string, 1)

type Filter interface {
	Do(response http.ResponseWriter, request *http.Request)
}
