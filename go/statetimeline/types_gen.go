// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package statetimeline

import (
	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

type Options struct {
	// Show timeline values on chart
	ShowValue common.VisibilityMode `json:"showValue"`
	// Controls the row height
	RowHeight float64 `json:"rowHeight"`
	// Merge equal consecutive values
	MergeValues *bool                    `json:"mergeValues,omitempty"`
	Legend      common.VizLegendOptions  `json:"legend"`
	Tooltip     common.VizTooltipOptions `json:"tooltip"`
	Timezone    []common.TimeZone        `json:"timezone,omitempty"`
	// Controls value alignment on the timelines
	AlignValue *common.TimelineValueAlignment `json:"alignValue,omitempty"`
}

type FieldConfig struct {
	LineWidth   *uint32                  `json:"lineWidth,omitempty"`
	HideFrom    *common.HideSeriesConfig `json:"hideFrom,omitempty"`
	FillOpacity *uint32                  `json:"fillOpacity,omitempty"`
}
