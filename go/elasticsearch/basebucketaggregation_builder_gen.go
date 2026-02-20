// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[BaseBucketAggregation] = (*BaseBucketAggregationBuilder)(nil)

type BaseBucketAggregationBuilder struct {
	internal *BaseBucketAggregation
	errors   cog.BuildErrors
}

func NewBaseBucketAggregationBuilder() *BaseBucketAggregationBuilder {
	resource := NewBaseBucketAggregation()
	builder := &BaseBucketAggregationBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *BaseBucketAggregationBuilder) Build() (BaseBucketAggregation, error) {
	if err := builder.internal.Validate(); err != nil {
		return BaseBucketAggregation{}, err
	}

	if len(builder.errors) > 0 {
		return BaseBucketAggregation{}, cog.MakeBuildErrors("elasticsearch.baseBucketAggregation", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *BaseBucketAggregationBuilder) RecordError(path string, err error) *BaseBucketAggregationBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *BaseBucketAggregationBuilder) Id(id string) *BaseBucketAggregationBuilder {
	builder.internal.Id = id

	return builder
}

func (builder *BaseBucketAggregationBuilder) Type(typeArg BucketAggregationType) *BaseBucketAggregationBuilder {
	builder.internal.Type = typeArg

	return builder
}

func (builder *BaseBucketAggregationBuilder) Settings(settings any) *BaseBucketAggregationBuilder {
	builder.internal.Settings = &settings

	return builder
}
