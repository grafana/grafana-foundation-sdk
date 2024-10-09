// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package gauge

import (
	"encoding/json"

	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

type Options struct {
	ShowThresholdLabels  bool                          `json:"showThresholdLabels"`
	ReduceOptions        common.ReduceDataOptions      `json:"reduceOptions"`
	Text                 *common.VizTextDisplayOptions `json:"text,omitempty"`
	ShowThresholdMarkers bool                          `json:"showThresholdMarkers"`
	Orientation          common.VizOrientation         `json:"orientation"`
}

func (resource Options) Equals(other Options) bool {
	if resource.ShowThresholdLabels != other.ShowThresholdLabels {
		return false
	}
	if !resource.ReduceOptions.Equals(other.ReduceOptions) {
		return false
	}
	if resource.Text == nil && other.Text != nil || resource.Text != nil && other.Text == nil {
		return false
	}

	if resource.Text != nil {
		if !resource.Text.Equals(*other.Text) {
			return false
		}
	}
	if resource.ShowThresholdMarkers != other.ShowThresholdMarkers {
		return false
	}
	if resource.Orientation != other.Orientation {
		return false
	}

	return true
}

func VariantConfig() variants.PanelcfgConfig {
	return variants.PanelcfgConfig{
		Identifier: "gauge",
		OptionsUnmarshaler: func(raw []byte) (any, error) {
			options := &Options{}

			if err := json.Unmarshal(raw, options); err != nil {
				return nil, err
			}

			return options, nil
		},
		GoConverter: func(inputPanel any) string {
			if panel, ok := inputPanel.(*dashboard.Panel); ok {
				return PanelConverter(*panel)
			}

			return PanelConverter(inputPanel.(dashboard.Panel))
		},
	}
}
