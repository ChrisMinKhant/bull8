package filterchain

import (
	"log"
	"net/http"

	"github.com/ChrisMinKhant/megoyougo_framework/exception"
	"github.com/ChrisMinKhant/megoyougo_framework/filter"
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

func (filterList *filterList) Invoke(response http.ResponseWriter, request *http.Request) {

	if filterList.filterNode.nextFilterNode == nil {

		logrus.Warn("No filter was found. No filtration on request may not be safe.")
		return

	}

	tempFilterNode := filterList.filterNode

	/*
	 * All the filters in the filter list are
	 * invoked through for loop.
	 */

	for tempFilterNode.nextFilterNode != nil {

		tempFilterNode = tempFilterNode.nextFilterNode
		tempFilterNode.filter.Do(response, request)

	}

}
