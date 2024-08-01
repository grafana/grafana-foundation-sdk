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

type Options struct {
	PieType       PieChartType                  `json:"pieType"`
	DisplayLabels []PieChartLabels              `json:"displayLabels"`
	Tooltip       common.VizTooltipOptions      `json:"tooltip"`
	ReduceOptions common.ReduceDataOptions      `json:"reduceOptions"`
	Text          *common.VizTextDisplayOptions `json:"text,omitempty"`
	Legend        PieChartLegendOptions         `json:"legend"`
	Orientation   common.VizOrientation         `json:"orientation"`
}

type FieldConfig = common.HideableFieldConfig

func VariantConfig() variants.PanelcfgConfig {
	return variants.PanelcfgConfig{
		Identifier: "piechart",
		OptionsUnmarshaler: func(raw []byte) (any, error) {
			options := Options{}

			if err := json.Unmarshal(raw, &options); err != nil {
				return nil, err
			}

			return options, nil
		},
		FieldConfigUnmarshaler: func(raw []byte) (any, error) {
			fieldConfig := FieldConfig{}

			if err := json.Unmarshal(raw, &fieldConfig); err != nil {
				return nil, err
			}

			return fieldConfig, nil
		},
	}
}
