// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alertgroups

import (
	"encoding/json"

	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

type Options struct {
	// Comma-separated list of values used to filter alert results
	Labels string `json:"labels"`
	// Name of the alertmanager used as a source for alerts
	Alertmanager string `json:"alertmanager"`
	// Expand all alert groups by default
	ExpandAll bool `json:"expandAll"`
}

func (resource Options) Equals(other Options) bool {
	if resource.Labels != other.Labels {
		return false
	}
	if resource.Alertmanager != other.Alertmanager {
		return false
	}
	if resource.ExpandAll != other.ExpandAll {
		return false
	}

	return true
}

func VariantConfig() variants.PanelcfgConfig {
	return variants.PanelcfgConfig{
		Identifier: "alertgroups",
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
