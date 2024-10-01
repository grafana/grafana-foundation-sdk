// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package bargauge

import (
	"encoding/json"

	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
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

func (resource Options) Equals(other Options) bool {
	if resource.DisplayMode != other.DisplayMode {
		return false
	}
	if resource.ValueMode != other.ValueMode {
		return false
	}
	if resource.ShowUnfilled != other.ShowUnfilled {
		return false
	}
	if resource.MinVizWidth != other.MinVizWidth {
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
	if resource.MinVizHeight != other.MinVizHeight {
		return false
	}
	if resource.Orientation != other.Orientation {
		return false
	}

	return true
}

func VariantConfig() variants.PanelcfgConfig {
	return variants.PanelcfgConfig{
		Identifier: "bargauge",
		OptionsUnmarshaler: func(raw []byte) (any, error) {
			options := &Options{}

			if err := json.Unmarshal(raw, options); err != nil {
				return nil, err
			}

			return options, nil
		},
	}
}
