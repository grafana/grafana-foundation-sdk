// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[AzureMetricDimension] = (*AzureMetricDimensionBuilder)(nil)

type AzureMetricDimensionBuilder struct {
	internal *AzureMetricDimension
	errors   map[string]cog.BuildErrors
}

func NewAzureMetricDimensionBuilder() *AzureMetricDimensionBuilder {
	resource := &AzureMetricDimension{}
	builder := &AzureMetricDimensionBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *AzureMetricDimensionBuilder) Build() (AzureMetricDimension, error) {
	if err := builder.internal.Validate(); err != nil {
		return AzureMetricDimension{}, err
	}

	return *builder.internal, nil
}

// Name of Dimension to be filtered on.
func (builder *AzureMetricDimensionBuilder) Dimension(dimension string) *AzureMetricDimensionBuilder {
	builder.internal.Dimension = &dimension

	return builder
}

// String denoting the filter operation. Supports 'eq' - equals,'ne' - not equals, 'sw' - starts with. Note that some dimensions may not support all operators.
func (builder *AzureMetricDimensionBuilder) Operator(operator string) *AzureMetricDimensionBuilder {
	builder.internal.Operator = &operator

	return builder
}

// Values to match with the filter.
func (builder *AzureMetricDimensionBuilder) Filters(filters []string) *AzureMetricDimensionBuilder {
	builder.internal.Filters = filters

	return builder
}

// @deprecated filter is deprecated in favour of filters to support multiselect.
func (builder *AzureMetricDimensionBuilder) Filter(filter string) *AzureMetricDimensionBuilder {
	builder.internal.Filter = &filter

	return builder
}

func (builder *AzureMetricDimensionBuilder) applyDefaults() {
}
