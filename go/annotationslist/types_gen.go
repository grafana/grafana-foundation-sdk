// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package annotationslist

import (
	"encoding/json"

	cogvariants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
)

type Options struct {
	OnlyFromThisDashboard bool     `json:"onlyFromThisDashboard"`
	OnlyInTimeRange       bool     `json:"onlyInTimeRange"`
	Tags                  []string `json:"tags"`
	Limit                 uint32   `json:"limit"`
	ShowUser              bool     `json:"showUser"`
	ShowTime              bool     `json:"showTime"`
	ShowTags              bool     `json:"showTags"`
	NavigateToPanel       bool     `json:"navigateToPanel"`
	NavigateBefore        string   `json:"navigateBefore"`
	NavigateAfter         string   `json:"navigateAfter"`
}

func VariantConfig() cogvariants.PanelcfgConfig {
	return cogvariants.PanelcfgConfig{
		Identifier: "annolist",
		OptionsUnmarshaler: func(raw []byte) (any, error) {
			options := Options{}

			if err := json.Unmarshal(raw, &options); err != nil {
				return nil, err
			}

			return options, nil
		},
	}
}
