// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package bargauge

import (
	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

type Options struct {
	DisplayMode   common.BarGaugeDisplayMode    `json:"displayMode"`
	ValueMode     common.BarGaugeValueMode      `json:"valueMode"`
	ShowUnfilled  bool                          `json:"showUnfilled"`
	MinVizWidth   uint32                        `json:"minVizWidth"`
	ReduceOptions common.ReduceDataOptions      `json:"reduceOptions"`
	Text          *common.VizTextDisplayOptions `json:"text,omitempty"`
	MinVizHeight  uint32                        `json:"minVizHeight"`
	Orientation   common.VizOrientation         `json:"orientation"`
}
