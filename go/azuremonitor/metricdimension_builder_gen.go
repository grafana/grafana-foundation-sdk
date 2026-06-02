// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[MetricDimension] = (*MetricDimensionBuilder)(nil)

type MetricDimensionBuilder struct {
	internal *MetricDimension
	errors   cog.BuildErrors
}

func NewMetricDimensionBuilder() *MetricDimensionBuilder {
	resource := NewMetricDimension()
	builder := &MetricDimensionBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *MetricDimensionBuilder) Build() (MetricDimension, error) {
	if err := builder.internal.Validate(); err != nil {
		return MetricDimension{}, err
	}

	if len(builder.errors) > 0 {
		return MetricDimension{}, cog.MakeBuildErrors("azuremonitor.metricDimension", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *MetricDimensionBuilder) RecordError(path string, err error) *MetricDimensionBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

// Name of Dimension to be filtered on.
func (builder *MetricDimensionBuilder) Dimension(dimension string) *MetricDimensionBuilder {
	builder.internal.Dimension = &dimension

	return builder
}

// String denoting the filter operation. Supports 'eq' - equals,'ne' - not equals, 'sw' - starts with. Note that some dimensions may not support all operators.
func (builder *MetricDimensionBuilder) Operator(operator string) *MetricDimensionBuilder {
	builder.internal.Operator = &operator

	return builder
}

// Values to match with the filter.
func (builder *MetricDimensionBuilder) Filters(filters []string) *MetricDimensionBuilder {
	builder.internal.Filters = filters

	return builder
}

// @deprecated filter is deprecated in favour of filters to support multiselect.
func (builder *MetricDimensionBuilder) Filter(filter string) *MetricDimensionBuilder {
	builder.internal.Filter = &filter

	return builder
}
