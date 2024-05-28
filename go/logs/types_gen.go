// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package logs

import (
	"encoding/json"

	cogvariants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
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
}

func VariantConfig() cogvariants.PanelcfgConfig {
	return cogvariants.PanelcfgConfig{
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
