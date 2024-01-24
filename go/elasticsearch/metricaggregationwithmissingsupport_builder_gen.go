// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[MetricAggregationWithMissingSupport] = (*MetricAggregationWithMissingSupportBuilder)(nil)

type MetricAggregationWithMissingSupportBuilder struct {
	internal *MetricAggregationWithMissingSupport
	errors   map[string]cog.BuildErrors
}

func NewMetricAggregationWithMissingSupportBuilder() *MetricAggregationWithMissingSupportBuilder {
	resource := &MetricAggregationWithMissingSupport{}
	builder := &MetricAggregationWithMissingSupportBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *MetricAggregationWithMissingSupportBuilder) Build() (MetricAggregationWithMissingSupport, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("MetricAggregationWithMissingSupport", err)...)
	}

	if len(errs) != 0 {
		return MetricAggregationWithMissingSupport{}, errs
	}

	return *builder.internal, nil
}

func (builder *MetricAggregationWithMissingSupportBuilder) Settings(settings struct {
	Missing *string `json:"missing,omitempty"`
}) *MetricAggregationWithMissingSupportBuilder {
	builder.internal.Settings = &settings

	return builder
}

func (builder *MetricAggregationWithMissingSupportBuilder) Type(typeArg cog.Builder[MetricAggregationType]) *MetricAggregationWithMissingSupportBuilder {
	typeArgResource, err := typeArg.Build()
	if err != nil {
		builder.errors["type"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Type = typeArgResource

	return builder
}

func (builder *MetricAggregationWithMissingSupportBuilder) Id(id string) *MetricAggregationWithMissingSupportBuilder {
	builder.internal.Id = id

	return builder
}

func (builder *MetricAggregationWithMissingSupportBuilder) Hide(hide bool) *MetricAggregationWithMissingSupportBuilder {
	builder.internal.Hide = &hide

	return builder
}

func (builder *MetricAggregationWithMissingSupportBuilder) applyDefaults() {
}
