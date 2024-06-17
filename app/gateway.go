package app

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/ChrisMinKhant/megoyougo_framework/filter/filterchain"
	"github.com/ChrisMinKhant/megoyougo_framework/provider/handlerprovider"
	"github.com/sirupsen/logrus"
)

/*
* Gateway is where all the requests to the
* server must pass. The filteration by filter chiain
* on each request is done here. There is also dispatcher
* for requests.
 */

type gateWay struct {
	filterChain *filterchain.FilterChain
}

func NewGateWay() *gateWay {

	filterChain := filterchain.New()
	filterChain.Set()

	return &gateWay{
		filterChain: filterChain,
	}
}

func (gateWay *gateWay) ServeHTTP(response http.ResponseWriter, request *http.Request) {

	logrus.Info("The request reaches to the gateway.")

	/*
	 * The request will be filtered by
	 * the defined filterchain in defined
	 * order.
	 */

	gateWay.filterChain.Invoke(response, request)

	/*
	 * The request is dispatched to
	 * the relative handler according to
	 * the requested path.
	 */

	gateWay.dispatchRequest(response, request)

}

func (gateWay *gateWay) dispatchRequest(response http.ResponseWriter, request *http.Request) {

	logrus.Infof("Dispatching request for [ Path ::: %v , Method ::: %v ]", request.RequestURI, request.Method)

	handlerMap := handlerprovider.GetHandler()

	for endpoint, handler := range handlerMap {

		path, method := gateWay.fetchPathAndMethod(endpoint)

		if request.RequestURI == path && request.Method == method {

			handler(response, request)
			return
		}

	}

	logrus.Errorf("Path not found for ::: %v\n", request.RequestURI)

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusNotFound)
	json.NewEncoder(response).Encode(NewErrorResponse().status("NOT FOUND").message("Path not found.").path(request.RequestURI).timestamp(time.Now().String()))
}

func (gateWay *gateWay) fetchPathAndMethod(endpoint string) (string, string) {

	splitedString := strings.Split(endpoint, "|")
	return splitedString[0], splitedString[1]
}
