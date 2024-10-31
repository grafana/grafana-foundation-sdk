// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package stat

import (
	"encoding/json"
	"errors"
	"fmt"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

type Options struct {
	GraphMode     common.BigValueGraphMode      `json:"graphMode"`
	ColorMode     common.BigValueColorMode      `json:"colorMode"`
	JustifyMode   common.BigValueJustifyMode    `json:"justifyMode"`
	TextMode      common.BigValueTextMode       `json:"textMode"`
	ReduceOptions common.ReduceDataOptions      `json:"reduceOptions"`
	Text          *common.VizTextDisplayOptions `json:"text,omitempty"`
	WideLayout    bool                          `json:"wideLayout"`
	Orientation   common.VizOrientation         `json:"orientation"`
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
	// Field "graphMode"
	if fields["graphMode"] != nil {
		if string(fields["graphMode"]) != "null" {
			if err := json.Unmarshal(fields["graphMode"], &resource.GraphMode); err != nil {
				errs = append(errs, cog.MakeBuildErrors("graphMode", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("graphMode", errors.New("required field is null"))...)

		}
		delete(fields, "graphMode")
	} else {
		errs = append(errs, cog.MakeBuildErrors("graphMode", errors.New("required field is missing from input"))...)
	}
	// Field "colorMode"
	if fields["colorMode"] != nil {
		if string(fields["colorMode"]) != "null" {
			if err := json.Unmarshal(fields["colorMode"], &resource.ColorMode); err != nil {
				errs = append(errs, cog.MakeBuildErrors("colorMode", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("colorMode", errors.New("required field is null"))...)

		}
		delete(fields, "colorMode")
	} else {
		errs = append(errs, cog.MakeBuildErrors("colorMode", errors.New("required field is missing from input"))...)
	}
	// Field "justifyMode"
	if fields["justifyMode"] != nil {
		if string(fields["justifyMode"]) != "null" {
			if err := json.Unmarshal(fields["justifyMode"], &resource.JustifyMode); err != nil {
				errs = append(errs, cog.MakeBuildErrors("justifyMode", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("justifyMode", errors.New("required field is null"))...)

		}
		delete(fields, "justifyMode")
	} else {
		errs = append(errs, cog.MakeBuildErrors("justifyMode", errors.New("required field is missing from input"))...)
	}
	// Field "textMode"
	if fields["textMode"] != nil {
		if string(fields["textMode"]) != "null" {
			if err := json.Unmarshal(fields["textMode"], &resource.TextMode); err != nil {
				errs = append(errs, cog.MakeBuildErrors("textMode", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("textMode", errors.New("required field is null"))...)

		}
		delete(fields, "textMode")
	} else {
		errs = append(errs, cog.MakeBuildErrors("textMode", errors.New("required field is missing from input"))...)
	}
	// Field "reduceOptions"
	if fields["reduceOptions"] != nil {
		if string(fields["reduceOptions"]) != "null" {

			resource.ReduceOptions = common.ReduceDataOptions{}
			if err := resource.ReduceOptions.UnmarshalJSONStrict(fields["reduceOptions"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("reduceOptions", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("reduceOptions", errors.New("required field is null"))...)

		}
		delete(fields, "reduceOptions")
	} else {
		errs = append(errs, cog.MakeBuildErrors("reduceOptions", errors.New("required field is missing from input"))...)
	}
	// Field "text"
	if fields["text"] != nil {
		if string(fields["text"]) != "null" {

			resource.Text = &common.VizTextDisplayOptions{}
			if err := resource.Text.UnmarshalJSONStrict(fields["text"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("text", err)...)
			}

		}
		delete(fields, "text")

	}
	// Field "wideLayout"
	if fields["wideLayout"] != nil {
		if string(fields["wideLayout"]) != "null" {
			if err := json.Unmarshal(fields["wideLayout"], &resource.WideLayout); err != nil {
				errs = append(errs, cog.MakeBuildErrors("wideLayout", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("wideLayout", errors.New("required field is null"))...)

		}
		delete(fields, "wideLayout")
	} else {
		errs = append(errs, cog.MakeBuildErrors("wideLayout", errors.New("required field is missing from input"))...)
	}
	// Field "orientation"
	if fields["orientation"] != nil {
		if string(fields["orientation"]) != "null" {
			if err := json.Unmarshal(fields["orientation"], &resource.Orientation); err != nil {
				errs = append(errs, cog.MakeBuildErrors("orientation", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("orientation", errors.New("required field is null"))...)

		}
		delete(fields, "orientation")
	} else {
		errs = append(errs, cog.MakeBuildErrors("orientation", errors.New("required field is missing from input"))...)
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
	if resource.GraphMode != other.GraphMode {
		return false
	}
	if resource.ColorMode != other.ColorMode {
		return false
	}
	if resource.JustifyMode != other.JustifyMode {
		return false
	}
	if resource.TextMode != other.TextMode {
		return false
	}
	if !resource.ReduceOptions.Equals(other.ReduceOptions) {
		return false
	}
	if resource.Text == nil && other.Text != nil || resource.Text != nil && other.Text == nil {
		return false
	}

	if resource.Text != nil {
		if !resource.Text.Equals(*other.Text) {
			return false
		}
	}
	if resource.WideLayout != other.WideLayout {
		return false
	}
	if resource.Orientation != other.Orientation {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `Options` fields for violations and returns them.
func (resource Options) Validate() error {
	var errs cog.BuildErrors
	if err := resource.ReduceOptions.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("reduceOptions", err)...)
	}
	if resource.Text != nil {
		if err := resource.Text.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("text", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// VariantConfig returns the configuration related to stat panels.
// This configuration describes how to unmarshal it, convert it to code, …
func VariantConfig() variants.PanelcfgConfig {
	return variants.PanelcfgConfig{
		Identifier: "stat",
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
