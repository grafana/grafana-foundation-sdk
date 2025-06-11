// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[BucketAggregationWithField] = (*BucketAggregationWithFieldBuilder)(nil)

type BucketAggregationWithFieldBuilder struct {
	internal *BucketAggregationWithField
	errors   cog.BuildErrors
}

func NewBucketAggregationWithFieldBuilder() *BucketAggregationWithFieldBuilder {
	resource := NewBucketAggregationWithField()
	builder := &BucketAggregationWithFieldBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *BucketAggregationWithFieldBuilder) Build() (BucketAggregationWithField, error) {
	if err := builder.internal.Validate(); err != nil {
		return BucketAggregationWithField{}, err
	}

	if len(builder.errors) > 0 {
		return BucketAggregationWithField{}, cog.MakeBuildErrors("elasticsearch.bucketAggregationWithField", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *BucketAggregationWithFieldBuilder) Field(field string) *BucketAggregationWithFieldBuilder {
	builder.internal.Field = &field

	return builder
}

func (builder *BucketAggregationWithFieldBuilder) Id(id string) *BucketAggregationWithFieldBuilder {
	builder.internal.Id = id

	return builder
}

func (builder *BucketAggregationWithFieldBuilder) Type(typeArg BucketAggregationType) *BucketAggregationWithFieldBuilder {
	builder.internal.Type = typeArg

	return builder
}

func (builder *BucketAggregationWithFieldBuilder) Settings(settings any) *BucketAggregationWithFieldBuilder {
	builder.internal.Settings = &settings

	return builder
}
