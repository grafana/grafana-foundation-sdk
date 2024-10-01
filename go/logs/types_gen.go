// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package logs

import (
	"encoding/json"

	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

type Options struct {
	ShowLabels         bool                     `json:"showLabels"`
	ShowCommonLabels   bool                     `json:"showCommonLabels"`
	ShowTime           bool                     `json:"showTime"`
	WrapLogMessage     bool                     `json:"wrapLogMessage"`
	PrettifyLogMessage bool                     `json:"prettifyLogMessage"`
	EnableLogDetails   bool                     `json:"enableLogDetails"`
	SortOrder          common.LogsSortOrder     `json:"sortOrder"`
	DedupStrategy      common.LogsDedupStrategy `json:"dedupStrategy"`
}

func (resource Options) Equals(other Options) bool {
	if resource.ShowLabels != other.ShowLabels {
		return false
	}
	if resource.ShowCommonLabels != other.ShowCommonLabels {
		return false
	}
	if resource.ShowTime != other.ShowTime {
		return false
	}
	if resource.WrapLogMessage != other.WrapLogMessage {
		return false
	}
	if resource.PrettifyLogMessage != other.PrettifyLogMessage {
		return false
	}
	if resource.EnableLogDetails != other.EnableLogDetails {
		return false
	}
	if resource.SortOrder != other.SortOrder {
		return false
	}
	if resource.DedupStrategy != other.DedupStrategy {
		return false
	}

	return true
}

func VariantConfig() variants.PanelcfgConfig {
	return variants.PanelcfgConfig{
		Identifier: "logs",
		OptionsUnmarshaler: func(raw []byte) (any, error) {
			options := &Options{}

			if err := json.Unmarshal(raw, options); err != nil {
				return nil, err
			}

			return options, nil
		},
	}
}
