// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package annotationslist

import (
	"encoding/json"
	"errors"
	"fmt"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
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
	// Field "onlyFromThisDashboard"
	if fields["onlyFromThisDashboard"] != nil {
		if string(fields["onlyFromThisDashboard"]) != "null" {
			if err := json.Unmarshal(fields["onlyFromThisDashboard"], &resource.OnlyFromThisDashboard); err != nil {
				errs = append(errs, cog.MakeBuildErrors("onlyFromThisDashboard", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("onlyFromThisDashboard", errors.New("required field is null"))...)

		}
		delete(fields, "onlyFromThisDashboard")
	} else {
		errs = append(errs, cog.MakeBuildErrors("onlyFromThisDashboard", errors.New("required field is missing from input"))...)
	}
	// Field "onlyInTimeRange"
	if fields["onlyInTimeRange"] != nil {
		if string(fields["onlyInTimeRange"]) != "null" {
			if err := json.Unmarshal(fields["onlyInTimeRange"], &resource.OnlyInTimeRange); err != nil {
				errs = append(errs, cog.MakeBuildErrors("onlyInTimeRange", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("onlyInTimeRange", errors.New("required field is null"))...)

		}
		delete(fields, "onlyInTimeRange")
	} else {
		errs = append(errs, cog.MakeBuildErrors("onlyInTimeRange", errors.New("required field is missing from input"))...)
	}
	// Field "tags"
	if fields["tags"] != nil {
		if string(fields["tags"]) != "null" {

			if err := json.Unmarshal(fields["tags"], &resource.Tags); err != nil {
				errs = append(errs, cog.MakeBuildErrors("tags", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("tags", errors.New("required field is null"))...)

		}
		delete(fields, "tags")
	} else {
		errs = append(errs, cog.MakeBuildErrors("tags", errors.New("required field is missing from input"))...)
	}
	// Field "limit"
	if fields["limit"] != nil {
		if string(fields["limit"]) != "null" {
			if err := json.Unmarshal(fields["limit"], &resource.Limit); err != nil {
				errs = append(errs, cog.MakeBuildErrors("limit", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("limit", errors.New("required field is null"))...)

		}
		delete(fields, "limit")
	} else {
		errs = append(errs, cog.MakeBuildErrors("limit", errors.New("required field is missing from input"))...)
	}
	// Field "showUser"
	if fields["showUser"] != nil {
		if string(fields["showUser"]) != "null" {
			if err := json.Unmarshal(fields["showUser"], &resource.ShowUser); err != nil {
				errs = append(errs, cog.MakeBuildErrors("showUser", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("showUser", errors.New("required field is null"))...)

		}
		delete(fields, "showUser")
	} else {
		errs = append(errs, cog.MakeBuildErrors("showUser", errors.New("required field is missing from input"))...)
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
	// Field "showTags"
	if fields["showTags"] != nil {
		if string(fields["showTags"]) != "null" {
			if err := json.Unmarshal(fields["showTags"], &resource.ShowTags); err != nil {
				errs = append(errs, cog.MakeBuildErrors("showTags", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("showTags", errors.New("required field is null"))...)

		}
		delete(fields, "showTags")
	} else {
		errs = append(errs, cog.MakeBuildErrors("showTags", errors.New("required field is missing from input"))...)
	}
	// Field "navigateToPanel"
	if fields["navigateToPanel"] != nil {
		if string(fields["navigateToPanel"]) != "null" {
			if err := json.Unmarshal(fields["navigateToPanel"], &resource.NavigateToPanel); err != nil {
				errs = append(errs, cog.MakeBuildErrors("navigateToPanel", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("navigateToPanel", errors.New("required field is null"))...)

		}
		delete(fields, "navigateToPanel")
	} else {
		errs = append(errs, cog.MakeBuildErrors("navigateToPanel", errors.New("required field is missing from input"))...)
	}
	// Field "navigateBefore"
	if fields["navigateBefore"] != nil {
		if string(fields["navigateBefore"]) != "null" {
			if err := json.Unmarshal(fields["navigateBefore"], &resource.NavigateBefore); err != nil {
				errs = append(errs, cog.MakeBuildErrors("navigateBefore", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("navigateBefore", errors.New("required field is null"))...)

		}
		delete(fields, "navigateBefore")
	} else {
		errs = append(errs, cog.MakeBuildErrors("navigateBefore", errors.New("required field is missing from input"))...)
	}
	// Field "navigateAfter"
	if fields["navigateAfter"] != nil {
		if string(fields["navigateAfter"]) != "null" {
			if err := json.Unmarshal(fields["navigateAfter"], &resource.NavigateAfter); err != nil {
				errs = append(errs, cog.MakeBuildErrors("navigateAfter", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("navigateAfter", errors.New("required field is null"))...)

		}
		delete(fields, "navigateAfter")
	} else {
		errs = append(errs, cog.MakeBuildErrors("navigateAfter", errors.New("required field is missing from input"))...)
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

// Validate checks all the validation constraints that may be defined on `Options` fields for violations and returns them.
func (resource Options) Validate() error {
	return nil
}

// VariantConfig returns the configuration related to annolist panels.
// This configuration describes how to unmarshal it, convert it to code, …
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
