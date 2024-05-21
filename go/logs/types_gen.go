// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package logs

import (
	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

type Options struct {
	ShowLabels           bool                     `json:"showLabels"`
	ShowCommonLabels     bool                     `json:"showCommonLabels"`
	ShowTime             bool                     `json:"showTime"`
	ShowLogContextToggle bool                     `json:"showLogContextToggle"`
	WrapLogMessage       bool                     `json:"wrapLogMessage"`
	PrettifyLogMessage   bool                     `json:"prettifyLogMessage"`
	EnableLogDetails     bool                     `json:"enableLogDetails"`
	SortOrder            common.LogsSortOrder     `json:"sortOrder"`
	DedupStrategy        common.LogsDedupStrategy `json:"dedupStrategy"`
}
