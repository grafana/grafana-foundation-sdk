// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package bargauge

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
	DisplayMode   common.BarGaugeDisplayMode    `json:"displayMode"`
	ValueMode     common.BarGaugeValueMode      `json:"valueMode"`
	NamePlacement common.BarGaugeNamePlacement  `json:"namePlacement"`
	ShowUnfilled  bool                          `json:"showUnfilled"`
	Sizing        common.BarGaugeSizing         `json:"sizing"`
	MinVizWidth   uint32                        `json:"minVizWidth"`
	MinVizHeight  uint32                        `json:"minVizHeight"`
	Legend        common.VizLegendOptions       `json:"legend"`
	ReduceOptions common.ReduceDataOptions      `json:"reduceOptions"`
	Text          *common.VizTextDisplayOptions `json:"text,omitempty"`
	MaxVizHeight  uint32                        `json:"maxVizHeight"`
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
	// Field "displayMode"
	if fields["displayMode"] != nil {
		if string(fields["displayMode"]) != "null" {
			if err := json.Unmarshal(fields["displayMode"], &resource.DisplayMode); err != nil {
				errs = append(errs, cog.MakeBuildErrors("displayMode", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("displayMode", errors.New("required field is null"))...)

		}
		delete(fields, "displayMode")
	} else {
		errs = append(errs, cog.MakeBuildErrors("displayMode", errors.New("required field is missing from input"))...)
	}
	// Field "valueMode"
	if fields["valueMode"] != nil {
		if string(fields["valueMode"]) != "null" {
			if err := json.Unmarshal(fields["valueMode"], &resource.ValueMode); err != nil {
				errs = append(errs, cog.MakeBuildErrors("valueMode", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("valueMode", errors.New("required field is null"))...)

		}
		delete(fields, "valueMode")
	} else {
		errs = append(errs, cog.MakeBuildErrors("valueMode", errors.New("required field is missing from input"))...)
	}
	// Field "namePlacement"
	if fields["namePlacement"] != nil {
		if string(fields["namePlacement"]) != "null" {
			if err := json.Unmarshal(fields["namePlacement"], &resource.NamePlacement); err != nil {
				errs = append(errs, cog.MakeBuildErrors("namePlacement", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("namePlacement", errors.New("required field is null"))...)

		}
		delete(fields, "namePlacement")
	} else {
		errs = append(errs, cog.MakeBuildErrors("namePlacement", errors.New("required field is missing from input"))...)
	}
	// Field "showUnfilled"
	if fields["showUnfilled"] != nil {
		if string(fields["showUnfilled"]) != "null" {
			if err := json.Unmarshal(fields["showUnfilled"], &resource.ShowUnfilled); err != nil {
				errs = append(errs, cog.MakeBuildErrors("showUnfilled", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("showUnfilled", errors.New("required field is null"))...)

		}
		delete(fields, "showUnfilled")
	} else {
		errs = append(errs, cog.MakeBuildErrors("showUnfilled", errors.New("required field is missing from input"))...)
	}
	// Field "sizing"
	if fields["sizing"] != nil {
		if string(fields["sizing"]) != "null" {
			if err := json.Unmarshal(fields["sizing"], &resource.Sizing); err != nil {
				errs = append(errs, cog.MakeBuildErrors("sizing", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("sizing", errors.New("required field is null"))...)

		}
		delete(fields, "sizing")
	} else {
		errs = append(errs, cog.MakeBuildErrors("sizing", errors.New("required field is missing from input"))...)
	}
	// Field "minVizWidth"
	if fields["minVizWidth"] != nil {
		if string(fields["minVizWidth"]) != "null" {
			if err := json.Unmarshal(fields["minVizWidth"], &resource.MinVizWidth); err != nil {
				errs = append(errs, cog.MakeBuildErrors("minVizWidth", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("minVizWidth", errors.New("required field is null"))...)

		}
		delete(fields, "minVizWidth")
	} else {
		errs = append(errs, cog.MakeBuildErrors("minVizWidth", errors.New("required field is missing from input"))...)
	}
	// Field "minVizHeight"
	if fields["minVizHeight"] != nil {
		if string(fields["minVizHeight"]) != "null" {
			if err := json.Unmarshal(fields["minVizHeight"], &resource.MinVizHeight); err != nil {
				errs = append(errs, cog.MakeBuildErrors("minVizHeight", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("minVizHeight", errors.New("required field is null"))...)

		}
		delete(fields, "minVizHeight")
	} else {
		errs = append(errs, cog.MakeBuildErrors("minVizHeight", errors.New("required field is missing from input"))...)
	}
	// Field "legend"
	if fields["legend"] != nil {
		if string(fields["legend"]) != "null" {

			resource.Legend = common.VizLegendOptions{}
			if err := resource.Legend.UnmarshalJSONStrict(fields["legend"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("legend", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("legend", errors.New("required field is null"))...)

		}
		delete(fields, "legend")
	} else {
		errs = append(errs, cog.MakeBuildErrors("legend", errors.New("required field is missing from input"))...)
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
	// Field "maxVizHeight"
	if fields["maxVizHeight"] != nil {
		if string(fields["maxVizHeight"]) != "null" {
			if err := json.Unmarshal(fields["maxVizHeight"], &resource.MaxVizHeight); err != nil {
				errs = append(errs, cog.MakeBuildErrors("maxVizHeight", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("maxVizHeight", errors.New("required field is null"))...)

		}
		delete(fields, "maxVizHeight")
	} else {
		errs = append(errs, cog.MakeBuildErrors("maxVizHeight", errors.New("required field is missing from input"))...)
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
	if resource.DisplayMode != other.DisplayMode {
		return false
	}
	if resource.ValueMode != other.ValueMode {
		return false
	}
	if resource.NamePlacement != other.NamePlacement {
		return false
	}
	if resource.ShowUnfilled != other.ShowUnfilled {
		return false
	}
	if resource.Sizing != other.Sizing {
		return false
	}
	if resource.MinVizWidth != other.MinVizWidth {
		return false
	}
	if resource.MinVizHeight != other.MinVizHeight {
		return false
	}
	if !resource.Legend.Equals(other.Legend) {
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
	if resource.MaxVizHeight != other.MaxVizHeight {
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
	if err := resource.Legend.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("legend", err)...)
	}
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

// VariantConfig returns the configuration related to bargauge panels.
// This configuration describes how to unmarshal it, convert it to code, …
func VariantConfig() variants.PanelcfgConfig {
	return variants.PanelcfgConfig{
		Identifier: "bargauge",
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
