// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardlist

import (
	"encoding/json"

	cogvariants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
)

func VariantConfig() cogvariants.PanelcfgConfig {
	return cogvariants.PanelcfgConfig{
		Identifier: "dashlist",
		OptionsUnmarshaler: func(raw []byte) (any, error) {
			options := Options{}

			if err := json.Unmarshal(raw, &options); err != nil {
				return nil, err
			}

			return options, nil
		},
	}
}