// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package trend

import (
	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

// Identical to timeseries... except it does not have timezone settings
type Options struct {
	Legend  common.VizLegendOptions  `json:"legend"`
	Tooltip common.VizTooltipOptions `json:"tooltip"`
	// Name of the x field to use (defaults to first number)
	XField *string `json:"xField,omitempty"`
}

type FieldConfig common.GraphFieldConfig
