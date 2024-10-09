// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package logs

import (
	"encoding/json"
	"reflect"

	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboard"
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
	OnClickFilterLabel     any      `json:"onClickFilterLabel,omitempty"`
	OnClickFilterOutLabel  any      `json:"onClickFilterOutLabel,omitempty"`
	IsFilterLabelActive    any      `json:"isFilterLabelActive,omitempty"`
	OnClickFilterString    any      `json:"onClickFilterString,omitempty"`
	OnClickFilterOutString any      `json:"onClickFilterOutString,omitempty"`
	OnClickShowField       any      `json:"onClickShowField,omitempty"`
	OnClickHideField       any      `json:"onClickHideField,omitempty"`
	DisplayedFields        []string `json:"displayedFields,omitempty"`
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
	if resource.ShowLogContextToggle != other.ShowLogContextToggle {
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
	// is DeepEqual good enough here?
	if !reflect.DeepEqual(resource.OnClickFilterLabel, other.OnClickFilterLabel) {
		return false
	}
	// is DeepEqual good enough here?
	if !reflect.DeepEqual(resource.OnClickFilterOutLabel, other.OnClickFilterOutLabel) {
		return false
	}
	// is DeepEqual good enough here?
	if !reflect.DeepEqual(resource.IsFilterLabelActive, other.IsFilterLabelActive) {
		return false
	}
	// is DeepEqual good enough here?
	if !reflect.DeepEqual(resource.OnClickFilterString, other.OnClickFilterString) {
		return false
	}
	// is DeepEqual good enough here?
	if !reflect.DeepEqual(resource.OnClickFilterOutString, other.OnClickFilterOutString) {
		return false
	}
	// is DeepEqual good enough here?
	if !reflect.DeepEqual(resource.OnClickShowField, other.OnClickShowField) {
		return false
	}
	// is DeepEqual good enough here?
	if !reflect.DeepEqual(resource.OnClickHideField, other.OnClickHideField) {
		return false
	}

	if len(resource.DisplayedFields) != len(other.DisplayedFields) {
		return false
	}

	for i1 := range resource.DisplayedFields {
		if resource.DisplayedFields[i1] != other.DisplayedFields[i1] {
			return false
		}
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
		GoConverter: func(inputPanel any) string {
			if panel, ok := inputPanel.(*dashboard.Panel); ok {
				return PanelConverter(*panel)
			}

			return PanelConverter(inputPanel.(dashboard.Panel))
		},
	}
}
