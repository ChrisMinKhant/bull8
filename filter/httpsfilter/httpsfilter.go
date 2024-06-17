package httpsfilter

import (
	"net/http"

	"github.com/ChrisMinKhant/megoyougo_framework/filter"
)

type httpsFilter struct {
}

func New() *httpsFilter {
	return &httpsFilter{}
}

func (httpsFilter *httpsFilter) Do(response http.ResponseWriter, request *http.Request) {
	if request.TLS == nil {
		filter.ErrorSigal <- " HTTPS FILTER FAILED "
	}
}
