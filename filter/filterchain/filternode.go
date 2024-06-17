package filterchain

import "github.com/ChrisMinKhant/megoyougo_framework/filter"

type filterNode struct {
	filter         filter.Filter
	nextFilterNode *filterNode
}

func NewFilterNode(filter filter.Filter, nextFilterNode *filterNode) *filterNode {
	return &filterNode{
		filter:         filter,
		nextFilterNode: nextFilterNode,
	}
}
