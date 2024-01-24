// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package bargauge

import (
	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

type Options struct {
	DisplayMode   common.BarGaugeDisplayMode    `json:"displayMode"`
	ValueMode     common.BarGaugeValueMode      `json:"valueMode"`
	NamePlacement common.BarGaugeNamePlacement  `json:"namePlacement"`
	ShowUnfilled  bool                          `json:"showUnfilled"`
	Sizing        common.BarGaugeSizing         `json:"sizing"`
	MinVizWidth   uint32                        `json:"minVizWidth"`
	MinVizHeight  uint32                        `json:"minVizHeight"`
	ReduceOptions common.ReduceDataOptions      `json:"reduceOptions"`
	Text          *common.VizTextDisplayOptions `json:"text,omitempty"`
	MaxVizHeight  uint32                        `json:"maxVizHeight"`
	Orientation   common.VizOrientation         `json:"orientation"`
}
