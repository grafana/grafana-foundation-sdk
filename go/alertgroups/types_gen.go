// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alertgroups

import (
	"encoding/json"
	"errors"
	"fmt"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
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
	// Field "labels"
	if fields["labels"] != nil {
		if string(fields["labels"]) != "null" {
			if err := json.Unmarshal(fields["labels"], &resource.Labels); err != nil {
				errs = append(errs, cog.MakeBuildErrors("labels", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("labels", errors.New("required field is null"))...)

		}
		delete(fields, "labels")
	} else {
		errs = append(errs, cog.MakeBuildErrors("labels", errors.New("required field is missing from input"))...)
	}
	// Field "alertmanager"
	if fields["alertmanager"] != nil {
		if string(fields["alertmanager"]) != "null" {
			if err := json.Unmarshal(fields["alertmanager"], &resource.Alertmanager); err != nil {
				errs = append(errs, cog.MakeBuildErrors("alertmanager", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("alertmanager", errors.New("required field is null"))...)

		}
		delete(fields, "alertmanager")
	} else {
		errs = append(errs, cog.MakeBuildErrors("alertmanager", errors.New("required field is missing from input"))...)
	}
	// Field "expandAll"
	if fields["expandAll"] != nil {
		if string(fields["expandAll"]) != "null" {
			if err := json.Unmarshal(fields["expandAll"], &resource.ExpandAll); err != nil {
				errs = append(errs, cog.MakeBuildErrors("expandAll", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("expandAll", errors.New("required field is null"))...)

		}
		delete(fields, "expandAll")
	} else {
		errs = append(errs, cog.MakeBuildErrors("expandAll", errors.New("required field is missing from input"))...)
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

// Validate checks all the validation constraints that may be defined on `Options` fields for violations and returns them.
func (resource Options) Validate() error {
	return nil
}

// VariantConfig returns the configuration related to alertgroups panels.
// This configuration describes how to unmarshal it, convert it to code, …
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
