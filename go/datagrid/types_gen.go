// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package datagrid

import (
	"encoding/json"

	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
)

type Options struct {
	SelectedSeries int32 `json:"selectedSeries"`
}

func (resource Options) Equals(other Options) bool {
	if resource.SelectedSeries != other.SelectedSeries {
		return false
	}

	return true
}

func VariantConfig() variants.PanelcfgConfig {
	return variants.PanelcfgConfig{
		Identifier: "datagrid",
		OptionsUnmarshaler: func(raw []byte) (any, error) {
			options := &Options{}

			if err := json.Unmarshal(raw, options); err != nil {
				return nil, err
			}

			return options, nil
		},
	}
}
