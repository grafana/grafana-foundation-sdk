// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package piechart

import (
	"encoding/json"
	"errors"
	"fmt"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

// Select the pie chart display style.
type PieChartType string

const (
	PieChartTypePie   PieChartType = "pie"
	PieChartTypeDonut PieChartType = "donut"
)

// Select labels to display on the pie chart.
//   - Name - The series or field name.
//   - Percent - The percentage of the whole.
//   - Value - The raw numerical value.
type PieChartLabels string

const (
	PieChartLabelsName    PieChartLabels = "name"
	PieChartLabelsValue   PieChartLabels = "value"
	PieChartLabelsPercent PieChartLabels = "percent"
)

// Select values to display in the legend.
//   - Percent: The percentage of the whole.
//   - Value: The raw numerical value.
type PieChartLegendValues string

const (
	PieChartLegendValuesValue   PieChartLegendValues = "value"
	PieChartLegendValuesPercent PieChartLegendValues = "percent"
)

type PieChartLegendOptions struct {
	Values      []PieChartLegendValues   `json:"values"`
	DisplayMode common.LegendDisplayMode `json:"displayMode"`
	Placement   common.LegendPlacement   `json:"placement"`
	ShowLegend  bool                     `json:"showLegend"`
	AsTable     *bool                    `json:"asTable,omitempty"`
	IsVisible   *bool                    `json:"isVisible,omitempty"`
	SortBy      *string                  `json:"sortBy,omitempty"`
	SortDesc    *bool                    `json:"sortDesc,omitempty"`
	Width       *float64                 `json:"width,omitempty"`
	Calcs       []string                 `json:"calcs"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `PieChartLegendOptions` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *PieChartLegendOptions) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "values"
	if fields["values"] != nil {
		if string(fields["values"]) != "null" {

			if err := json.Unmarshal(fields["values"], &resource.Values); err != nil {
				errs = append(errs, cog.MakeBuildErrors("values", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("values", errors.New("required field is null"))...)

		}
		delete(fields, "values")
	} else {
		errs = append(errs, cog.MakeBuildErrors("values", errors.New("required field is missing from input"))...)
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
	// Field "placement"
	if fields["placement"] != nil {
		if string(fields["placement"]) != "null" {
			if err := json.Unmarshal(fields["placement"], &resource.Placement); err != nil {
				errs = append(errs, cog.MakeBuildErrors("placement", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("placement", errors.New("required field is null"))...)

		}
		delete(fields, "placement")
	} else {
		errs = append(errs, cog.MakeBuildErrors("placement", errors.New("required field is missing from input"))...)
	}
	// Field "showLegend"
	if fields["showLegend"] != nil {
		if string(fields["showLegend"]) != "null" {
			if err := json.Unmarshal(fields["showLegend"], &resource.ShowLegend); err != nil {
				errs = append(errs, cog.MakeBuildErrors("showLegend", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("showLegend", errors.New("required field is null"))...)

		}
		delete(fields, "showLegend")
	} else {
		errs = append(errs, cog.MakeBuildErrors("showLegend", errors.New("required field is missing from input"))...)
	}
	// Field "asTable"
	if fields["asTable"] != nil {
		if string(fields["asTable"]) != "null" {
			if err := json.Unmarshal(fields["asTable"], &resource.AsTable); err != nil {
				errs = append(errs, cog.MakeBuildErrors("asTable", err)...)
			}

		}
		delete(fields, "asTable")

	}
	// Field "isVisible"
	if fields["isVisible"] != nil {
		if string(fields["isVisible"]) != "null" {
			if err := json.Unmarshal(fields["isVisible"], &resource.IsVisible); err != nil {
				errs = append(errs, cog.MakeBuildErrors("isVisible", err)...)
			}

		}
		delete(fields, "isVisible")

	}
	// Field "sortBy"
	if fields["sortBy"] != nil {
		if string(fields["sortBy"]) != "null" {
			if err := json.Unmarshal(fields["sortBy"], &resource.SortBy); err != nil {
				errs = append(errs, cog.MakeBuildErrors("sortBy", err)...)
			}

		}
		delete(fields, "sortBy")

	}
	// Field "sortDesc"
	if fields["sortDesc"] != nil {
		if string(fields["sortDesc"]) != "null" {
			if err := json.Unmarshal(fields["sortDesc"], &resource.SortDesc); err != nil {
				errs = append(errs, cog.MakeBuildErrors("sortDesc", err)...)
			}

		}
		delete(fields, "sortDesc")

	}
	// Field "width"
	if fields["width"] != nil {
		if string(fields["width"]) != "null" {
			if err := json.Unmarshal(fields["width"], &resource.Width); err != nil {
				errs = append(errs, cog.MakeBuildErrors("width", err)...)
			}

		}
		delete(fields, "width")

	}
	// Field "calcs"
	if fields["calcs"] != nil {
		if string(fields["calcs"]) != "null" {

			if err := json.Unmarshal(fields["calcs"], &resource.Calcs); err != nil {
				errs = append(errs, cog.MakeBuildErrors("calcs", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("calcs", errors.New("required field is null"))...)

		}
		delete(fields, "calcs")
	} else {
		errs = append(errs, cog.MakeBuildErrors("calcs", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("PieChartLegendOptions", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `PieChartLegendOptions` objects.
func (resource PieChartLegendOptions) Equals(other PieChartLegendOptions) bool {

	if len(resource.Values) != len(other.Values) {
		return false
	}

	for i1 := range resource.Values {
		if resource.Values[i1] != other.Values[i1] {
			return false
		}
	}
	if resource.DisplayMode != other.DisplayMode {
		return false
	}
	if resource.Placement != other.Placement {
		return false
	}
	if resource.ShowLegend != other.ShowLegend {
		return false
	}
	if resource.AsTable == nil && other.AsTable != nil || resource.AsTable != nil && other.AsTable == nil {
		return false
	}

	if resource.AsTable != nil {
		if *resource.AsTable != *other.AsTable {
			return false
		}
	}
	if resource.IsVisible == nil && other.IsVisible != nil || resource.IsVisible != nil && other.IsVisible == nil {
		return false
	}

	if resource.IsVisible != nil {
		if *resource.IsVisible != *other.IsVisible {
			return false
		}
	}
	if resource.SortBy == nil && other.SortBy != nil || resource.SortBy != nil && other.SortBy == nil {
		return false
	}

	if resource.SortBy != nil {
		if *resource.SortBy != *other.SortBy {
			return false
		}
	}
	if resource.SortDesc == nil && other.SortDesc != nil || resource.SortDesc != nil && other.SortDesc == nil {
		return false
	}

	if resource.SortDesc != nil {
		if *resource.SortDesc != *other.SortDesc {
			return false
		}
	}
	if resource.Width == nil && other.Width != nil || resource.Width != nil && other.Width == nil {
		return false
	}

	if resource.Width != nil {
		if *resource.Width != *other.Width {
			return false
		}
	}

	if len(resource.Calcs) != len(other.Calcs) {
		return false
	}

	for i1 := range resource.Calcs {
		if resource.Calcs[i1] != other.Calcs[i1] {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `PieChartLegendOptions` fields for violations and returns them.
func (resource PieChartLegendOptions) Validate() error {
	return nil
}

type Options struct {
	PieType       PieChartType                  `json:"pieType"`
	DisplayLabels []PieChartLabels              `json:"displayLabels,omitempty"`
	Tooltip       common.VizTooltipOptions      `json:"tooltip"`
	ReduceOptions common.ReduceDataOptions      `json:"reduceOptions"`
	Text          *common.VizTextDisplayOptions `json:"text,omitempty"`
	Legend        PieChartLegendOptions         `json:"legend"`
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
	// Field "pieType"
	if fields["pieType"] != nil {
		if string(fields["pieType"]) != "null" {
			if err := json.Unmarshal(fields["pieType"], &resource.PieType); err != nil {
				errs = append(errs, cog.MakeBuildErrors("pieType", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("pieType", errors.New("required field is null"))...)

		}
		delete(fields, "pieType")
	} else {
		errs = append(errs, cog.MakeBuildErrors("pieType", errors.New("required field is missing from input"))...)
	}
	// Field "displayLabels"
	if fields["displayLabels"] != nil {
		if string(fields["displayLabels"]) != "null" {

			if err := json.Unmarshal(fields["displayLabels"], &resource.DisplayLabels); err != nil {
				errs = append(errs, cog.MakeBuildErrors("displayLabels", err)...)
			}

		}
		delete(fields, "displayLabels")

	}
	// Field "tooltip"
	if fields["tooltip"] != nil {
		if string(fields["tooltip"]) != "null" {

			resource.Tooltip = common.VizTooltipOptions{}
			if err := resource.Tooltip.UnmarshalJSONStrict(fields["tooltip"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("tooltip", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("tooltip", errors.New("required field is null"))...)

		}
		delete(fields, "tooltip")
	} else {
		errs = append(errs, cog.MakeBuildErrors("tooltip", errors.New("required field is missing from input"))...)
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
	// Field "legend"
	if fields["legend"] != nil {
		if string(fields["legend"]) != "null" {

			resource.Legend = PieChartLegendOptions{}
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
	if resource.PieType != other.PieType {
		return false
	}

	if len(resource.DisplayLabels) != len(other.DisplayLabels) {
		return false
	}

	for i1 := range resource.DisplayLabels {
		if resource.DisplayLabels[i1] != other.DisplayLabels[i1] {
			return false
		}
	}
	if !resource.Tooltip.Equals(other.Tooltip) {
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
	if !resource.Legend.Equals(other.Legend) {
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
	if err := resource.Tooltip.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("tooltip", err)...)
	}
	if err := resource.ReduceOptions.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("reduceOptions", err)...)
	}
	if resource.Text != nil {
		if err := resource.Text.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("text", err)...)
		}
	}
	if err := resource.Legend.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("legend", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type FieldConfig = common.HideableFieldConfig

// VariantConfig returns the configuration related to piechart panels.
// This configuration describes how to unmarshal it, convert it to code, …
func VariantConfig() variants.PanelcfgConfig {
	return variants.PanelcfgConfig{
		Identifier: "piechart",
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
		FieldConfigUnmarshaler: func(raw []byte) (any, error) {
			fieldConfig := &FieldConfig{}

			if err := json.Unmarshal(raw, fieldConfig); err != nil {
				return nil, err
			}

			return fieldConfig, nil
		},
		StrictFieldConfigUnmarshaler: func(raw []byte) (any, error) {
			fieldConfig := &FieldConfig{}

			if err := fieldConfig.UnmarshalJSONStrict(raw); err != nil {
				return nil, err
			}

			return fieldConfig, nil
		},
		GoConverter: func(inputPanel any) string {
			if panel, ok := inputPanel.(*dashboard.Panel); ok {
				return PanelConverter(*panel)
			}

			return PanelConverter(inputPanel.(dashboard.Panel))
		},
	}
}
