package authfilter

import (
	"log"
	"net/http"
)

type authFilter struct {
}

func New() *authFilter {
	return &authFilter{}
}

func (headerFilter *authFilter) Do(response http.ResponseWriter, request *http.Request) {
	log.Println("Auth filter is invoked.")
}
