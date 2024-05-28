// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alertgroups

import (
	"encoding/json"

	cogvariants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
)

type Options struct {
	// Comma-separated list of values used to filter alert results
	Labels string `json:"labels"`
	// Name of the alertmanager used as a source for alerts
	Alertmanager string `json:"alertmanager"`
	// Expand all alert groups by default
	ExpandAll bool `json:"expandAll"`
}

func VariantConfig() cogvariants.PanelcfgConfig {
	return cogvariants.PanelcfgConfig{
		Identifier: "alertgroups",
		OptionsUnmarshaler: func(raw []byte) (any, error) {
			options := Options{}

			if err := json.Unmarshal(raw, &options); err != nil {
				return nil, err
			}

			return options, nil
		},
	}
}
