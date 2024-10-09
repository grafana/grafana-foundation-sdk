// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package stat

import (
	"encoding/json"

	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

type Options struct {
	GraphMode         common.BigValueGraphMode      `json:"graphMode"`
	ColorMode         common.BigValueColorMode      `json:"colorMode"`
	JustifyMode       common.BigValueJustifyMode    `json:"justifyMode"`
	TextMode          common.BigValueTextMode       `json:"textMode"`
	WideLayout        bool                          `json:"wideLayout"`
	ReduceOptions     common.ReduceDataOptions      `json:"reduceOptions"`
	Text              *common.VizTextDisplayOptions `json:"text,omitempty"`
	ShowPercentChange bool                          `json:"showPercentChange"`
	Orientation       common.VizOrientation         `json:"orientation"`
}

func (resource Options) Equals(other Options) bool {
	if resource.GraphMode != other.GraphMode {
		return false
	}
	if resource.ColorMode != other.ColorMode {
		return false
	}
	if resource.JustifyMode != other.JustifyMode {
		return false
	}
	if resource.TextMode != other.TextMode {
		return false
	}
	if resource.WideLayout != other.WideLayout {
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
	if resource.ShowPercentChange != other.ShowPercentChange {
		return false
	}
	if resource.Orientation != other.Orientation {
		return false
	}

	return true
}

func VariantConfig() variants.PanelcfgConfig {
	return variants.PanelcfgConfig{
		Identifier: "stat",
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
