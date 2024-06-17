package headerfilter

import (
	"log"
	"net/http"
)

type headerFilter struct {
}

func New() *headerFilter {
	return &headerFilter{}
}

func (headerFilter *headerFilter) Do(response http.ResponseWriter, request *http.Request) {
	log.Println("Header filter is invoked.")
}
