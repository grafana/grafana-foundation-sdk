// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package bargauge

import (
	"encoding/json"

	cogvariants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
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

func VariantConfig() cogvariants.PanelcfgConfig {
	return cogvariants.PanelcfgConfig{
		Identifier: "bargauge",
		OptionsUnmarshaler: func(raw []byte) (any, error) {
			options := Options{}

			if err := json.Unmarshal(raw, &options); err != nil {
				return nil, err
			}

			return options, nil
		},
	}
}
