package handler

import "net/http"

type handler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}
