// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package statetimeline

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
		return Options{}, cog.MakeBuildErrors("statetimeline.options", builder.errors)
	}

	return *builder.internal, nil
}

// Show timeline values on chart
func (builder *OptionsBuilder) ShowValue(showValue common.VisibilityMode) *OptionsBuilder {
	builder.internal.ShowValue = showValue

	return builder
}

// Controls the row height
func (builder *OptionsBuilder) RowHeight(rowHeight float64) *OptionsBuilder {
	builder.internal.RowHeight = rowHeight

	return builder
}

// Merge equal consecutive values
func (builder *OptionsBuilder) MergeValues(mergeValues bool) *OptionsBuilder {
	builder.internal.MergeValues = &mergeValues

	return builder
}

// Controls value alignment on the timelines
func (builder *OptionsBuilder) AlignValue(alignValue common.TimelineValueAlignment) *OptionsBuilder {
	builder.internal.AlignValue = &alignValue

	return builder
}

func (builder *OptionsBuilder) Legend(legend cog.Builder[common.VizLegendOptions]) *OptionsBuilder {
	legendResource, err := legend.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Legend = legendResource

	return builder
}

func (builder *OptionsBuilder) Tooltip(tooltip cog.Builder[common.VizTooltipOptions]) *OptionsBuilder {
	tooltipResource, err := tooltip.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Tooltip = tooltipResource

	return builder
}

func (builder *OptionsBuilder) Timezone(timezone []common.TimeZone) *OptionsBuilder {
	builder.internal.Timezone = timezone

	return builder
}

// Enables pagination when > 0
func (builder *OptionsBuilder) PerPage(perPage float64) *OptionsBuilder {
	builder.internal.PerPage = &perPage

	return builder
}
