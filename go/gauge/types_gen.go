// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package gauge

import (
	"encoding/json"

	cogvariants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

type Options struct {
	ShowThresholdLabels  bool                          `json:"showThresholdLabels"`
	ReduceOptions        common.ReduceDataOptions      `json:"reduceOptions"`
	Text                 *common.VizTextDisplayOptions `json:"text,omitempty"`
	ShowThresholdMarkers bool                          `json:"showThresholdMarkers"`
	Orientation          common.VizOrientation         `json:"orientation"`
}

func VariantConfig() cogvariants.PanelcfgConfig {
	return cogvariants.PanelcfgConfig{
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
