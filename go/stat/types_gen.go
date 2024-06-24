// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package stat

import (
	"encoding/json"

	cogvariants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

type Options struct {
	GraphMode              common.BigValueGraphMode      `json:"graphMode"`
	ColorMode              common.BigValueColorMode      `json:"colorMode"`
	JustifyMode            common.BigValueJustifyMode    `json:"justifyMode"`
	TextMode               common.BigValueTextMode       `json:"textMode"`
	WideLayout             bool                          `json:"wideLayout"`
	ShowPercentChange      bool                          `json:"showPercentChange"`
	ReduceOptions          common.ReduceDataOptions      `json:"reduceOptions"`
	Text                   *common.VizTextDisplayOptions `json:"text,omitempty"`
	PercentChangeColorMode common.PercentChangeColorMode `json:"percentChangeColorMode"`
	Orientation            common.VizOrientation         `json:"orientation"`
}

func VariantConfig() cogvariants.PanelcfgConfig {
	return cogvariants.PanelcfgConfig{
		Identifier: "stat",
		OptionsUnmarshaler: func(raw []byte) (any, error) {
			options := Options{}

			if err := json.Unmarshal(raw, &options); err != nil {
				return nil, err
			}

			return options, nil
		},
	}
}
