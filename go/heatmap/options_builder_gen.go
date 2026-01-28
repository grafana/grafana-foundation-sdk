// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package heatmap

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

var _ cog.Builder[Options] = (*OptionsBuilder)(nil)

type OptionsBuilder struct {
	internal *Options
	errors   cog.BuildErrors
}

func NewOptionsBuilder() *OptionsBuilder {
	resource := NewOptions()
	builder := &OptionsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *OptionsBuilder) Build() (Options, error) {
	if err := builder.internal.Validate(); err != nil {
		return Options{}, err
	}

	if len(builder.errors) > 0 {
		return Options{}, cog.MakeBuildErrors("heatmap.options", builder.errors)
	}

	return *builder.internal, nil
}

// Controls if the heatmap should be calculated from data
func (builder *OptionsBuilder) Calculate(calculate bool) *OptionsBuilder {
	builder.internal.Calculate = &calculate

	return builder
}

// Calculation options for the heatmap
func (builder *OptionsBuilder) Calculation(calculation cog.Builder[common.HeatmapCalculationOptions]) *OptionsBuilder {
	calculationResource, err := calculation.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Calculation = &calculationResource

	return builder
}

// Controls the color options
func (builder *OptionsBuilder) Color(color cog.Builder[HeatmapColorOptions]) *OptionsBuilder {
	colorResource, err := color.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Color = colorResource

	return builder
}

// Filters values between a given range
func (builder *OptionsBuilder) FilterValues(filterValues cog.Builder[FilterValueRange]) *OptionsBuilder {
	filterValuesResource, err := filterValues.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.FilterValues = &filterValuesResource

	return builder
}

// Controls tick alignment and value name when not calculating from data
func (builder *OptionsBuilder) RowsFrame(rowsFrame cog.Builder[RowsHeatmapOptions]) *OptionsBuilder {
	rowsFrameResource, err := rowsFrame.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.RowsFrame = &rowsFrameResource

	return builder
}

//	| *{
//		layout: ui.HeatmapCellLayout & "auto" // TODO: fix after remove when https://github.com/grafana/cuetsy/issues/74 is fixed
//	}
//
// Controls the display of the value in the cell
func (builder *OptionsBuilder) ShowValue(showValue common.VisibilityMode) *OptionsBuilder {
	builder.internal.ShowValue = showValue

	return builder
}

// Controls gap between cells
func (builder *OptionsBuilder) CellGap(cellGap uint8) *OptionsBuilder {
	builder.internal.CellGap = &cellGap

	return builder
}

// Controls cell radius
func (builder *OptionsBuilder) CellRadius(cellRadius float32) *OptionsBuilder {
	builder.internal.CellRadius = &cellRadius

	return builder
}

// Controls cell value unit
func (builder *OptionsBuilder) CellValues(cellValues cog.Builder[CellValues]) *OptionsBuilder {
	cellValuesResource, err := cellValues.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.CellValues = &cellValuesResource

	return builder
}

// Controls yAxis placement
func (builder *OptionsBuilder) YAxis(yAxis cog.Builder[YAxisConfig]) *OptionsBuilder {
	yAxisResource, err := yAxis.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.YAxis = yAxisResource

	return builder
}

//	| *{
//		axisPlacement: ui.AxisPlacement & "left" // TODO: fix after remove when https://github.com/grafana/cuetsy/issues/74 is fixed
//	}
//
// Controls legend options
func (builder *OptionsBuilder) Legend(legend cog.Builder[HeatmapLegend]) *OptionsBuilder {
	legendResource, err := legend.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Legend = legendResource

	return builder
}

// Controls tooltip options
func (builder *OptionsBuilder) Tooltip(tooltip cog.Builder[HeatmapTooltip]) *OptionsBuilder {
	tooltipResource, err := tooltip.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Tooltip = tooltipResource

	return builder
}

// Controls exemplar options
func (builder *OptionsBuilder) Exemplars(exemplars cog.Builder[ExemplarConfig]) *OptionsBuilder {
	exemplarsResource, err := exemplars.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Exemplars = exemplarsResource

	return builder
}

// Controls which axis to allow selection on
func (builder *OptionsBuilder) SelectionMode(selectionMode HeatmapSelectionMode) *OptionsBuilder {
	builder.internal.SelectionMode = &selectionMode

	return builder
}
