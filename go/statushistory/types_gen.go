// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package statushistory

import (
	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

type Options struct {
	// Set the height of the rows
	RowHeight float32 `json:"rowHeight"`
	// Show values on the columns
	ShowValue common.VisibilityMode    `json:"showValue"`
	Legend    common.VizLegendOptions  `json:"legend"`
	Tooltip   common.VizTooltipOptions `json:"tooltip"`
	Timezone  []common.TimeZone        `json:"timezone,omitempty"`
	// Controls the column width
	ColWidth *float64 `json:"colWidth,omitempty"`
}

type FieldConfig struct {
	LineWidth   *uint32                  `json:"lineWidth,omitempty"`
	HideFrom    *common.HideSeriesConfig `json:"hideFrom,omitempty"`
	FillOpacity *uint32                  `json:"fillOpacity,omitempty"`
}
