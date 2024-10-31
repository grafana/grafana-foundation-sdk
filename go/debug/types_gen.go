// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package debug

import (
	"encoding/json"
	"errors"
	"fmt"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

type UpdateConfig struct {
	Render        bool `json:"render"`
	DataChanged   bool `json:"dataChanged"`
	SchemaChanged bool `json:"schemaChanged"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `UpdateConfig` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *UpdateConfig) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "render"
	if fields["render"] != nil {
		if string(fields["render"]) != "null" {
			if err := json.Unmarshal(fields["render"], &resource.Render); err != nil {
				errs = append(errs, cog.MakeBuildErrors("render", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("render", errors.New("required field is null"))...)

		}
		delete(fields, "render")
	} else {
		errs = append(errs, cog.MakeBuildErrors("render", errors.New("required field is missing from input"))...)
	}
	// Field "dataChanged"
	if fields["dataChanged"] != nil {
		if string(fields["dataChanged"]) != "null" {
			if err := json.Unmarshal(fields["dataChanged"], &resource.DataChanged); err != nil {
				errs = append(errs, cog.MakeBuildErrors("dataChanged", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("dataChanged", errors.New("required field is null"))...)

		}
		delete(fields, "dataChanged")
	} else {
		errs = append(errs, cog.MakeBuildErrors("dataChanged", errors.New("required field is missing from input"))...)
	}
	// Field "schemaChanged"
	if fields["schemaChanged"] != nil {
		if string(fields["schemaChanged"]) != "null" {
			if err := json.Unmarshal(fields["schemaChanged"], &resource.SchemaChanged); err != nil {
				errs = append(errs, cog.MakeBuildErrors("schemaChanged", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("schemaChanged", errors.New("required field is null"))...)

		}
		delete(fields, "schemaChanged")
	} else {
		errs = append(errs, cog.MakeBuildErrors("schemaChanged", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("UpdateConfig", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `UpdateConfig` objects.
func (resource UpdateConfig) Equals(other UpdateConfig) bool {
	if resource.Render != other.Render {
		return false
	}
	if resource.DataChanged != other.DataChanged {
		return false
	}
	if resource.SchemaChanged != other.SchemaChanged {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `UpdateConfig` fields for violations and returns them.
func (resource UpdateConfig) Validate() error {
	return nil
}

type DebugMode string

const (
	DebugModeRender     DebugMode = "render"
	DebugModeEvents     DebugMode = "events"
	DebugModeCursor     DebugMode = "cursor"
	DebugModeState      DebugMode = "State"
	DebugModeThrowError DebugMode = "ThrowError"
)

type Options struct {
	Mode     DebugMode     `json:"mode"`
	Counters *UpdateConfig `json:"counters,omitempty"`
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
	// Field "mode"
	if fields["mode"] != nil {
		if string(fields["mode"]) != "null" {
			if err := json.Unmarshal(fields["mode"], &resource.Mode); err != nil {
				errs = append(errs, cog.MakeBuildErrors("mode", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("mode", errors.New("required field is null"))...)

		}
		delete(fields, "mode")
	} else {
		errs = append(errs, cog.MakeBuildErrors("mode", errors.New("required field is missing from input"))...)
	}
	// Field "counters"
	if fields["counters"] != nil {
		if string(fields["counters"]) != "null" {

			resource.Counters = &UpdateConfig{}
			if err := resource.Counters.UnmarshalJSONStrict(fields["counters"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("counters", err)...)
			}

		}
		delete(fields, "counters")

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
	if resource.Mode != other.Mode {
		return false
	}
	if resource.Counters == nil && other.Counters != nil || resource.Counters != nil && other.Counters == nil {
		return false
	}

	if resource.Counters != nil {
		if !resource.Counters.Equals(*other.Counters) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `Options` fields for violations and returns them.
func (resource Options) Validate() error {
	var errs cog.BuildErrors
	if resource.Counters != nil {
		if err := resource.Counters.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("counters", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// VariantConfig returns the configuration related to debug panels.
// This configuration describes how to unmarshal it, convert it to code, …
func VariantConfig() variants.PanelcfgConfig {
	return variants.PanelcfgConfig{
		Identifier: "debug",
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
