// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package timeseries

import (
	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

type Options struct {
	Timezone []common.TimeZone        `json:"timezone,omitempty"`
	Legend   common.VizLegendOptions  `json:"legend"`
	Tooltip  common.VizTooltipOptions `json:"tooltip"`
}

type FieldConfig common.GraphFieldConfig
