// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[BaseBucketAggregation] = (*BaseBucketAggregationBuilder)(nil)

type BaseBucketAggregationBuilder struct {
	internal *BaseBucketAggregation
	errors   map[string]cog.BuildErrors
}

func NewBaseBucketAggregationBuilder() *BaseBucketAggregationBuilder {
	resource := &BaseBucketAggregation{}
	builder := &BaseBucketAggregationBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *BaseBucketAggregationBuilder) Build() (BaseBucketAggregation, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("BaseBucketAggregation", err)...)
	}

	if len(errs) != 0 {
		return BaseBucketAggregation{}, errs
	}

	return *builder.internal, nil
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

func (builder *BaseBucketAggregationBuilder) applyDefaults() {
}
