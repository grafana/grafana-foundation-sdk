// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package logs

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
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

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Options` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *Options) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "showLabels"
	if fields["showLabels"] != nil {
		if string(fields["showLabels"]) != "null" {
			if err := json.Unmarshal(fields["showLabels"], &resource.ShowLabels); err != nil {
				errs = append(errs, cog.MakeBuildErrors("showLabels", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("showLabels", errors.New("required field is null"))...)

		}
		delete(fields, "showLabels")
	} else {
		errs = append(errs, cog.MakeBuildErrors("showLabels", errors.New("required field is missing from input"))...)
	}
	// Field "showCommonLabels"
	if fields["showCommonLabels"] != nil {
		if string(fields["showCommonLabels"]) != "null" {
			if err := json.Unmarshal(fields["showCommonLabels"], &resource.ShowCommonLabels); err != nil {
				errs = append(errs, cog.MakeBuildErrors("showCommonLabels", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("showCommonLabels", errors.New("required field is null"))...)

		}
		delete(fields, "showCommonLabels")
	} else {
		errs = append(errs, cog.MakeBuildErrors("showCommonLabels", errors.New("required field is missing from input"))...)
	}
	// Field "showTime"
	if fields["showTime"] != nil {
		if string(fields["showTime"]) != "null" {
			if err := json.Unmarshal(fields["showTime"], &resource.ShowTime); err != nil {
				errs = append(errs, cog.MakeBuildErrors("showTime", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("showTime", errors.New("required field is null"))...)

		}
		delete(fields, "showTime")
	} else {
		errs = append(errs, cog.MakeBuildErrors("showTime", errors.New("required field is missing from input"))...)
	}
	// Field "showLogContextToggle"
	if fields["showLogContextToggle"] != nil {
		if string(fields["showLogContextToggle"]) != "null" {
			if err := json.Unmarshal(fields["showLogContextToggle"], &resource.ShowLogContextToggle); err != nil {
				errs = append(errs, cog.MakeBuildErrors("showLogContextToggle", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("showLogContextToggle", errors.New("required field is null"))...)

		}
		delete(fields, "showLogContextToggle")
	} else {
		errs = append(errs, cog.MakeBuildErrors("showLogContextToggle", errors.New("required field is missing from input"))...)
	}
	// Field "wrapLogMessage"
	if fields["wrapLogMessage"] != nil {
		if string(fields["wrapLogMessage"]) != "null" {
			if err := json.Unmarshal(fields["wrapLogMessage"], &resource.WrapLogMessage); err != nil {
				errs = append(errs, cog.MakeBuildErrors("wrapLogMessage", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("wrapLogMessage", errors.New("required field is null"))...)

		}
		delete(fields, "wrapLogMessage")
	} else {
		errs = append(errs, cog.MakeBuildErrors("wrapLogMessage", errors.New("required field is missing from input"))...)
	}
	// Field "prettifyLogMessage"
	if fields["prettifyLogMessage"] != nil {
		if string(fields["prettifyLogMessage"]) != "null" {
			if err := json.Unmarshal(fields["prettifyLogMessage"], &resource.PrettifyLogMessage); err != nil {
				errs = append(errs, cog.MakeBuildErrors("prettifyLogMessage", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("prettifyLogMessage", errors.New("required field is null"))...)

		}
		delete(fields, "prettifyLogMessage")
	} else {
		errs = append(errs, cog.MakeBuildErrors("prettifyLogMessage", errors.New("required field is missing from input"))...)
	}
	// Field "enableLogDetails"
	if fields["enableLogDetails"] != nil {
		if string(fields["enableLogDetails"]) != "null" {
			if err := json.Unmarshal(fields["enableLogDetails"], &resource.EnableLogDetails); err != nil {
				errs = append(errs, cog.MakeBuildErrors("enableLogDetails", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("enableLogDetails", errors.New("required field is null"))...)

		}
		delete(fields, "enableLogDetails")
	} else {
		errs = append(errs, cog.MakeBuildErrors("enableLogDetails", errors.New("required field is missing from input"))...)
	}
	// Field "sortOrder"
	if fields["sortOrder"] != nil {
		if string(fields["sortOrder"]) != "null" {
			if err := json.Unmarshal(fields["sortOrder"], &resource.SortOrder); err != nil {
				errs = append(errs, cog.MakeBuildErrors("sortOrder", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("sortOrder", errors.New("required field is null"))...)

		}
		delete(fields, "sortOrder")
	} else {
		errs = append(errs, cog.MakeBuildErrors("sortOrder", errors.New("required field is missing from input"))...)
	}
	// Field "dedupStrategy"
	if fields["dedupStrategy"] != nil {
		if string(fields["dedupStrategy"]) != "null" {
			if err := json.Unmarshal(fields["dedupStrategy"], &resource.DedupStrategy); err != nil {
				errs = append(errs, cog.MakeBuildErrors("dedupStrategy", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("dedupStrategy", errors.New("required field is null"))...)

		}
		delete(fields, "dedupStrategy")
	} else {
		errs = append(errs, cog.MakeBuildErrors("dedupStrategy", errors.New("required field is missing from input"))...)
	}
	// Field "onClickFilterLabel"
	if fields["onClickFilterLabel"] != nil {
		if string(fields["onClickFilterLabel"]) != "null" {
			if err := json.Unmarshal(fields["onClickFilterLabel"], &resource.OnClickFilterLabel); err != nil {
				errs = append(errs, cog.MakeBuildErrors("onClickFilterLabel", err)...)
			}

		}
		delete(fields, "onClickFilterLabel")

	}
	// Field "onClickFilterOutLabel"
	if fields["onClickFilterOutLabel"] != nil {
		if string(fields["onClickFilterOutLabel"]) != "null" {
			if err := json.Unmarshal(fields["onClickFilterOutLabel"], &resource.OnClickFilterOutLabel); err != nil {
				errs = append(errs, cog.MakeBuildErrors("onClickFilterOutLabel", err)...)
			}

		}
		delete(fields, "onClickFilterOutLabel")

	}
	// Field "isFilterLabelActive"
	if fields["isFilterLabelActive"] != nil {
		if string(fields["isFilterLabelActive"]) != "null" {
			if err := json.Unmarshal(fields["isFilterLabelActive"], &resource.IsFilterLabelActive); err != nil {
				errs = append(errs, cog.MakeBuildErrors("isFilterLabelActive", err)...)
			}

		}
		delete(fields, "isFilterLabelActive")

	}
	// Field "onClickFilterString"
	if fields["onClickFilterString"] != nil {
		if string(fields["onClickFilterString"]) != "null" {
			if err := json.Unmarshal(fields["onClickFilterString"], &resource.OnClickFilterString); err != nil {
				errs = append(errs, cog.MakeBuildErrors("onClickFilterString", err)...)
			}

		}
		delete(fields, "onClickFilterString")

	}
	// Field "onClickFilterOutString"
	if fields["onClickFilterOutString"] != nil {
		if string(fields["onClickFilterOutString"]) != "null" {
			if err := json.Unmarshal(fields["onClickFilterOutString"], &resource.OnClickFilterOutString); err != nil {
				errs = append(errs, cog.MakeBuildErrors("onClickFilterOutString", err)...)
			}

		}
		delete(fields, "onClickFilterOutString")

	}
	// Field "onClickShowField"
	if fields["onClickShowField"] != nil {
		if string(fields["onClickShowField"]) != "null" {
			if err := json.Unmarshal(fields["onClickShowField"], &resource.OnClickShowField); err != nil {
				errs = append(errs, cog.MakeBuildErrors("onClickShowField", err)...)
			}

		}
		delete(fields, "onClickShowField")

	}
	// Field "onClickHideField"
	if fields["onClickHideField"] != nil {
		if string(fields["onClickHideField"]) != "null" {
			if err := json.Unmarshal(fields["onClickHideField"], &resource.OnClickHideField); err != nil {
				errs = append(errs, cog.MakeBuildErrors("onClickHideField", err)...)
			}

		}
		delete(fields, "onClickHideField")

	}
	// Field "displayedFields"
	if fields["displayedFields"] != nil {
		if string(fields["displayedFields"]) != "null" {

			if err := json.Unmarshal(fields["displayedFields"], &resource.DisplayedFields); err != nil {
				errs = append(errs, cog.MakeBuildErrors("displayedFields", err)...)
			}

		}
		delete(fields, "displayedFields")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("Options", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `Options` objects.
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

// Validate checks all the validation constraints that may be defined on `Options` fields for violations and returns them.
func (resource Options) Validate() error {
	return nil
}

// VariantConfig returns the configuration related to logs panels.
// This configuration describes how to unmarshal it, convert it to code, …
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
		StrictOptionsUnmarshaler: func(raw []byte) (any, error) {
			options := &Options{}

			if err := options.UnmarshalJSONStrict(raw); err != nil {
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
