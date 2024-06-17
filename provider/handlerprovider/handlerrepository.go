package handlerprovider

import "net/http"

var handlerRepository = make(map[string]func(response http.ResponseWriter, request *http.Request))

func BindHandler(endpoint string, handler func(response http.ResponseWriter, request *http.Request)) {
	handlerRepository[endpoint] = handler
}

func GetHandler() map[string]func(response http.ResponseWriter, request *http.Request) {
	return handlerRepository
}
