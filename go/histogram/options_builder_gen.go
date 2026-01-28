// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package histogram

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
		return Options{}, cog.MakeBuildErrors("histogram.options", builder.errors)
	}

	return *builder.internal, nil
}

// Size of each bucket
func (builder *OptionsBuilder) BucketSize(bucketSize int32) *OptionsBuilder {
	builder.internal.BucketSize = &bucketSize

	return builder
}

// Offset buckets by this amount
func (builder *OptionsBuilder) BucketOffset(bucketOffset int32) *OptionsBuilder {
	builder.internal.BucketOffset = &bucketOffset

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

// Combines multiple series into a single histogram
func (builder *OptionsBuilder) Combine(combine bool) *OptionsBuilder {
	builder.internal.Combine = &combine

	return builder
}
