// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package piechart

import (
	"encoding/json"

	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
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

type Options struct {
	PieType       PieChartType                  `json:"pieType"`
	DisplayLabels []PieChartLabels              `json:"displayLabels,omitempty"`
	Tooltip       common.VizTooltipOptions      `json:"tooltip"`
	ReduceOptions common.ReduceDataOptions      `json:"reduceOptions"`
	Text          *common.VizTextDisplayOptions `json:"text,omitempty"`
	Legend        PieChartLegendOptions         `json:"legend"`
	Orientation   common.VizOrientation         `json:"orientation"`
}

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

type FieldConfig = common.HideableFieldConfig

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
		FieldConfigUnmarshaler: func(raw []byte) (any, error) {
			fieldConfig := &FieldConfig{}

			if err := json.Unmarshal(raw, fieldConfig); err != nil {
				return nil, err
			}

			return fieldConfig, nil
		},
	}
}
