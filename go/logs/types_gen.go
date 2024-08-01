// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package logs

import (
	"encoding/json"

	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

type Options struct {
	ShowLabels           bool                     `json:"showLabels"`
	ShowCommonLabels     bool                     `json:"showCommonLabels"`
	ShowTime             bool                     `json:"showTime"`
	ShowLogContextToggle bool                     `json:"showLogContextToggle"`
	WrapLogMessage       bool                     `json:"wrapLogMessage"`
	PrettifyLogMessage   bool                     `json:"prettifyLogMessage"`
	EnableLogDetails     bool                     `json:"enableLogDetails"`
	SortOrder            common.LogsSortOrder     `json:"sortOrder"`
	DedupStrategy        common.LogsDedupStrategy `json:"dedupStrategy"`
	// TODO: figure out how to define callbacks
	OnClickFilterLabel     any `json:"onClickFilterLabel,omitempty"`
	OnClickFilterOutLabel  any `json:"onClickFilterOutLabel,omitempty"`
	IsFilterLabelActive    any `json:"isFilterLabelActive,omitempty"`
	OnClickFilterString    any `json:"onClickFilterString,omitempty"`
	OnClickFilterOutString any `json:"onClickFilterOutString,omitempty"`
}

func VariantConfig() variants.PanelcfgConfig {
	return variants.PanelcfgConfig{
		Identifier: "logs",
		OptionsUnmarshaler: func(raw []byte) (any, error) {
			options := Options{}

			if err := json.Unmarshal(raw, &options); err != nil {
				return nil, err
			}

			return options, nil
		},
	}
}
