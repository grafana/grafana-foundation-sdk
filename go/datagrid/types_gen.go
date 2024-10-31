// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package datagrid

import (
	"encoding/json"
	"errors"
	"fmt"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

type Options struct {
	SelectedSeries int32 `json:"selectedSeries"`
}

func (resource *Options) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "selectedSeries"
	if fields["selectedSeries"] != nil {
		if string(fields["selectedSeries"]) != "null" {
			if err := json.Unmarshal(fields["selectedSeries"], &resource.SelectedSeries); err != nil {
				errs = append(errs, cog.MakeBuildErrors("selectedSeries", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("selectedSeries", errors.New("required field is null"))...)

		}
		delete(fields, "selectedSeries")
	} else {
		errs = append(errs, cog.MakeBuildErrors("selectedSeries", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("Options", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

func (resource Options) Equals(other Options) bool {
	if resource.SelectedSeries != other.SelectedSeries {
		return false
	}

	return true
}

// Validate checks any constraint that may be defined for this type
// and returns all violations.
func (resource Options) Validate() error {
	var errs cog.BuildErrors
	if !(resource.SelectedSeries >= 0) {
		errs = append(errs, cog.MakeBuildErrors(
			"selectedSeries",
			errors.New("must be >= 0"),
		)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

func VariantConfig() variants.PanelcfgConfig {
	return variants.PanelcfgConfig{
		Identifier: "datagrid",
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
