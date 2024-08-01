// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package gauge

import (
	"encoding/json"

	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

type Options struct {
	ShowThresholdLabels  bool                          `json:"showThresholdLabels"`
	ShowThresholdMarkers bool                          `json:"showThresholdMarkers"`
	Sizing               common.BarGaugeSizing         `json:"sizing"`
	MinVizWidth          uint32                        `json:"minVizWidth"`
	ReduceOptions        common.ReduceDataOptions      `json:"reduceOptions"`
	Text                 *common.VizTextDisplayOptions `json:"text,omitempty"`
	MinVizHeight         uint32                        `json:"minVizHeight"`
	Orientation          common.VizOrientation         `json:"orientation"`
}

func VariantConfig() variants.PanelcfgConfig {
	return variants.PanelcfgConfig{
		Identifier: "gauge",
		OptionsUnmarshaler: func(raw []byte) (any, error) {
			options := Options{}

			if err := json.Unmarshal(raw, &options); err != nil {
				return nil, err
			}

			return options, nil
		},
	}
}
