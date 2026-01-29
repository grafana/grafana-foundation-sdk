// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[MetricFindValue] = (*MetricFindValueBuilder)(nil)

// Define the MetricFindValue type
type MetricFindValueBuilder struct {
	internal *MetricFindValue
	errors   cog.BuildErrors
}

func NewMetricFindValueBuilder() *MetricFindValueBuilder {
	resource := NewMetricFindValue()
	builder := &MetricFindValueBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *MetricFindValueBuilder) Build() (MetricFindValue, error) {
	if err := builder.internal.Validate(); err != nil {
		return MetricFindValue{}, err
	}

	if len(builder.errors) > 0 {
		return MetricFindValue{}, cog.MakeBuildErrors("dashboardv2beta1.metricFindValue", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *MetricFindValueBuilder) Text(text string) *MetricFindValueBuilder {
	builder.internal.Text = text

	return builder
}

func (builder *MetricFindValueBuilder) Value(value StringOrFloat64) *MetricFindValueBuilder {
	builder.internal.Value = &value

	return builder
}

func (builder *MetricFindValueBuilder) Group(group string) *MetricFindValueBuilder {
	builder.internal.Group = &group

	return builder
}

func (builder *MetricFindValueBuilder) Expandable(expandable bool) *MetricFindValueBuilder {
	builder.internal.Expandable = &expandable

	return builder
}
