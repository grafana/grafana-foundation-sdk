// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package logsnew

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboard"
	dashboardv2beta1 "github.com/grafana/grafana-foundation-sdk/go/dashboardv2beta1"
)

type Options struct {
	ShowTime                bool                     `json:"showTime"`
	WrapLogMessage          bool                     `json:"wrapLogMessage"`
	EnableLogDetails        bool                     `json:"enableLogDetails"`
	SortOrder               common.LogsSortOrder     `json:"sortOrder"`
	DedupStrategy           common.LogsDedupStrategy `json:"dedupStrategy"`
	EnableInfiniteScrolling *bool                    `json:"enableInfiniteScrolling,omitempty"`
	OnNewLogsReceived       any                      `json:"onNewLogsReceived,omitempty"`
}

// NewOptions creates a new Options object.
func NewOptions() *Options {
	return &Options{}
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
	// Field "enableInfiniteScrolling"
	if fields["enableInfiniteScrolling"] != nil {
		if string(fields["enableInfiniteScrolling"]) != "null" {
			if err := json.Unmarshal(fields["enableInfiniteScrolling"], &resource.EnableInfiniteScrolling); err != nil {
				errs = append(errs, cog.MakeBuildErrors("enableInfiniteScrolling", err)...)
			}

		}
		delete(fields, "enableInfiniteScrolling")

	}
	// Field "onNewLogsReceived"
	if fields["onNewLogsReceived"] != nil {
		if string(fields["onNewLogsReceived"]) != "null" {
			if err := json.Unmarshal(fields["onNewLogsReceived"], &resource.OnNewLogsReceived); err != nil {
				errs = append(errs, cog.MakeBuildErrors("onNewLogsReceived", err)...)
			}

		}
		delete(fields, "onNewLogsReceived")

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
	if resource.ShowTime != other.ShowTime {
		return false
	}
	if resource.WrapLogMessage != other.WrapLogMessage {
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
	if resource.EnableInfiniteScrolling == nil && other.EnableInfiniteScrolling != nil || resource.EnableInfiniteScrolling != nil && other.EnableInfiniteScrolling == nil {
		return false
	}

	if resource.EnableInfiniteScrolling != nil {
		if *resource.EnableInfiniteScrolling != *other.EnableInfiniteScrolling {
			return false
		}
	}
	// is DeepEqual good enough here?
	if !reflect.DeepEqual(resource.OnNewLogsReceived, other.OnNewLogsReceived) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `Options` fields for violations and returns them.
func (resource Options) Validate() error {
	return nil
}

// VariantConfig returns the configuration related to logsnew panels.
// This configuration describes how to unmarshal it, convert it to code, …
func VariantConfig() variants.PanelcfgConfig {
	return variants.PanelcfgConfig{
		Identifier: "logsnew",
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
			if panel, ok := inputPanel.(dashboard.Panel); ok {
				return PanelConverter(panel)
			}
			if panel, ok := inputPanel.(*dashboardv2beta1.VizConfigKind); ok {
				return VisualizationConverter(*panel)
			}

			return VisualizationConverter(inputPanel.(dashboardv2beta1.VizConfigKind))
		},
	}
}
