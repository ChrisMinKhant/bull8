package filterchain

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/ChrisMinKhant/megoyougo_framework/exception"
	"github.com/ChrisMinKhant/megoyougo_framework/filter"
	"github.com/ChrisMinKhant/megoyougo_framework/util"
	"github.com/sirupsen/logrus"
)

type filterList struct {
	filterNode *filterNode
	exception  exception.Exception
}

func NewFilterList() *filterList {
	return &filterList{
		filterNode: NewFilterNode(nil, nil),
		exception:  exception.GetGeneralExceptionInstance(),
	}
}

func (filterList *filterList) Add(filter filter.Filter) {

	log.Println("New filter is added to the filter list.")

	if filterList.filterNode.nextFilterNode == nil {

		/*
		 * There is no filter existing in the filter list.
		 */

		filterList.filterNode.nextFilterNode = NewFilterNode(filter, nil)
		return

	}

	/*
	 * One or more filter exists in the filter list.
	 */

	tempFilterNode := filterList.filterNode

	for tempFilterNode.nextFilterNode != nil {
		tempFilterNode = tempFilterNode.nextFilterNode
	}

	tempFilterNode.nextFilterNode = NewFilterNode(filter, nil)

}

func (filterList *filterList) Invoke(response http.ResponseWriter, request *http.Request) bool {
	defer filterList.exception.RecoverPanic()

	if filterList.filterNode.nextFilterNode == nil {

		logrus.Warn("No filter was found. No filtration on request may not be safe.")
		return false

	}

	tempFilterNode := filterList.filterNode

	/*
	 * All the filters in the filter list are
	 * invoked through for loop.
	 */

	for tempFilterNode.nextFilterNode != nil {

		tempFilterNode = tempFilterNode.nextFilterNode
		tempFilterNode.filter.Do(response, request)

		// Fetching is there any error signal from added filters.
		// If there is any, the filteration will stop and app will response
		// with error code.
		if fetchedSignal := <-filter.ErrorSigal; fetchedSignal != "" {

			response.Header().Set("Content-Type", "application/json")
			response.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(response).Encode(util.NewErrorResponse().SetStatus("Filteration Failed").SetMessage(fetchedSignal).SetPath(request.RequestURI).SetTimestamp(time.Now().String()))

			logrus.Panicf("Filteration failed with error ::: [ %v ]\n", fetchedSignal)
		}

	}

	return true

}
