// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package annotationslist

import (
	"encoding/json"

	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboard"
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

func (resource Options) Equals(other Options) bool {
	if resource.OnlyFromThisDashboard != other.OnlyFromThisDashboard {
		return false
	}
	if resource.OnlyInTimeRange != other.OnlyInTimeRange {
		return false
	}

	if len(resource.Tags) != len(other.Tags) {
		return false
	}

	for i1 := range resource.Tags {
		if resource.Tags[i1] != other.Tags[i1] {
			return false
		}
	}
	if resource.Limit != other.Limit {
		return false
	}
	if resource.ShowUser != other.ShowUser {
		return false
	}
	if resource.ShowTime != other.ShowTime {
		return false
	}
	if resource.ShowTags != other.ShowTags {
		return false
	}
	if resource.NavigateToPanel != other.NavigateToPanel {
		return false
	}
	if resource.NavigateBefore != other.NavigateBefore {
		return false
	}
	if resource.NavigateAfter != other.NavigateAfter {
		return false
	}

	return true
}

func VariantConfig() variants.PanelcfgConfig {
	return variants.PanelcfgConfig{
		Identifier: "annolist",
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
