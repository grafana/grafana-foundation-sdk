// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package logs

import (
	"encoding/json"
	"errors"
	"fmt"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboard"
	dashboardv2beta1 "github.com/grafana/grafana-foundation-sdk/go/dashboardv2beta1"
)

type Options struct {
	ShowLabels              bool                     `json:"showLabels"`
	ShowCommonLabels        bool                     `json:"showCommonLabels"`
	ShowTime                bool                     `json:"showTime"`
	ShowLogContextToggle    bool                     `json:"showLogContextToggle"`
	WrapLogMessage          bool                     `json:"wrapLogMessage"`
	PrettifyLogMessage      bool                     `json:"prettifyLogMessage"`
	EnableLogDetails        bool                     `json:"enableLogDetails"`
	SortOrder               common.LogsSortOrder     `json:"sortOrder"`
	DedupStrategy           common.LogsDedupStrategy `json:"dedupStrategy"`
	EnableInfiniteScrolling *bool                    `json:"enableInfiniteScrolling,omitempty"`
	// Display controls to jump to the last or first log line, and filters by log level.
	ShowControls *bool `json:"showControls,omitempty"`
	// Show a component to manage the displayed fields from the logs.
	ShowFieldSelector *bool `json:"showFieldSelector,omitempty"`
	// Use a predefined coloring scheme to highlight relevant parts of the log lines.
	SyntaxHighlighting *bool               `json:"syntaxHighlighting,omitempty"`
	FontSize           *OptionsFontSize    `json:"fontSize,omitempty"`
	DetailsMode        *OptionsDetailsMode `json:"detailsMode,omitempty"`
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
	// Field "enableInfiniteScrolling"
	if fields["enableInfiniteScrolling"] != nil {
		if string(fields["enableInfiniteScrolling"]) != "null" {
			if err := json.Unmarshal(fields["enableInfiniteScrolling"], &resource.EnableInfiniteScrolling); err != nil {
				errs = append(errs, cog.MakeBuildErrors("enableInfiniteScrolling", err)...)
			}

		}
		delete(fields, "enableInfiniteScrolling")

	}
	// Field "showControls"
	if fields["showControls"] != nil {
		if string(fields["showControls"]) != "null" {
			if err := json.Unmarshal(fields["showControls"], &resource.ShowControls); err != nil {
				errs = append(errs, cog.MakeBuildErrors("showControls", err)...)
			}

		}
		delete(fields, "showControls")

	}
	// Field "showFieldSelector"
	if fields["showFieldSelector"] != nil {
		if string(fields["showFieldSelector"]) != "null" {
			if err := json.Unmarshal(fields["showFieldSelector"], &resource.ShowFieldSelector); err != nil {
				errs = append(errs, cog.MakeBuildErrors("showFieldSelector", err)...)
			}

		}
		delete(fields, "showFieldSelector")

	}
	// Field "syntaxHighlighting"
	if fields["syntaxHighlighting"] != nil {
		if string(fields["syntaxHighlighting"]) != "null" {
			if err := json.Unmarshal(fields["syntaxHighlighting"], &resource.SyntaxHighlighting); err != nil {
				errs = append(errs, cog.MakeBuildErrors("syntaxHighlighting", err)...)
			}

		}
		delete(fields, "syntaxHighlighting")

	}
	// Field "fontSize"
	if fields["fontSize"] != nil {
		if string(fields["fontSize"]) != "null" {
			if err := json.Unmarshal(fields["fontSize"], &resource.FontSize); err != nil {
				errs = append(errs, cog.MakeBuildErrors("fontSize", err)...)
			}

		}
		delete(fields, "fontSize")

	}
	// Field "detailsMode"
	if fields["detailsMode"] != nil {
		if string(fields["detailsMode"]) != "null" {
			if err := json.Unmarshal(fields["detailsMode"], &resource.DetailsMode); err != nil {
				errs = append(errs, cog.MakeBuildErrors("detailsMode", err)...)
			}

		}
		delete(fields, "detailsMode")

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
	if resource.EnableInfiniteScrolling == nil && other.EnableInfiniteScrolling != nil || resource.EnableInfiniteScrolling != nil && other.EnableInfiniteScrolling == nil {
		return false
	}

	if resource.EnableInfiniteScrolling != nil {
		if *resource.EnableInfiniteScrolling != *other.EnableInfiniteScrolling {
			return false
		}
	}
	if resource.ShowControls == nil && other.ShowControls != nil || resource.ShowControls != nil && other.ShowControls == nil {
		return false
	}

	if resource.ShowControls != nil {
		if *resource.ShowControls != *other.ShowControls {
			return false
		}
	}
	if resource.ShowFieldSelector == nil && other.ShowFieldSelector != nil || resource.ShowFieldSelector != nil && other.ShowFieldSelector == nil {
		return false
	}

	if resource.ShowFieldSelector != nil {
		if *resource.ShowFieldSelector != *other.ShowFieldSelector {
			return false
		}
	}
	if resource.SyntaxHighlighting == nil && other.SyntaxHighlighting != nil || resource.SyntaxHighlighting != nil && other.SyntaxHighlighting == nil {
		return false
	}

	if resource.SyntaxHighlighting != nil {
		if *resource.SyntaxHighlighting != *other.SyntaxHighlighting {
			return false
		}
	}
	if resource.FontSize == nil && other.FontSize != nil || resource.FontSize != nil && other.FontSize == nil {
		return false
	}

	if resource.FontSize != nil {
		if *resource.FontSize != *other.FontSize {
			return false
		}
	}
	if resource.DetailsMode == nil && other.DetailsMode != nil || resource.DetailsMode != nil && other.DetailsMode == nil {
		return false
	}

	if resource.DetailsMode != nil {
		if *resource.DetailsMode != *other.DetailsMode {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `Options` fields for violations and returns them.
func (resource Options) Validate() error {
	return nil
}

type OptionsFontSize string

const (
	OptionsFontSizeDefault OptionsFontSize = "default"
	OptionsFontSizeSmall   OptionsFontSize = "small"
)

type OptionsDetailsMode string

const (
	OptionsDetailsModeInline  OptionsDetailsMode = "inline"
	OptionsDetailsModeSidebar OptionsDetailsMode = "sidebar"
)

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
