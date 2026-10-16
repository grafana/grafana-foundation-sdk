// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[V2TransformationSpec] = (*V2TransformationSpecBuilder)(nil)

// Dashboard v2 shape: no `id`, it is carried by the parent's `group`.
type V2TransformationSpecBuilder struct {
	internal *V2TransformationSpec
	errors   cog.BuildErrors
}

func NewV2TransformationSpecBuilder() *V2TransformationSpecBuilder {
	resource := NewV2TransformationSpec()
	builder := &V2TransformationSpecBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *V2TransformationSpecBuilder) Build() (V2TransformationSpec, error) {
	if err := builder.internal.Validate(); err != nil {
		return V2TransformationSpec{}, err
	}

	if len(builder.errors) > 0 {
		return V2TransformationSpec{}, cog.MakeBuildErrors("dashboardv2beta1.v2TransformationSpec", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *V2TransformationSpecBuilder) RecordError(path string, err error) *V2TransformationSpecBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

// Disabled transformations are skipped
func (builder *V2TransformationSpecBuilder) Disabled(disabled bool) *V2TransformationSpecBuilder {
	builder.internal.Disabled = &disabled

	return builder
}

// Optional frame matcher. When missing it will be applied to all results
func (builder *V2TransformationSpecBuilder) Filter(filter MatcherConfig) *V2TransformationSpecBuilder {
	builder.internal.Filter = &filter

	return builder
}

// Where to pull DataFrames from as input to transformation
func (builder *V2TransformationSpecBuilder) Topic(topic DataTopic) *V2TransformationSpecBuilder {
	builder.internal.Topic = &topic

	return builder
}

// Options to be passed to the transformer
// Valid options depend on the transformer id
func (builder *V2TransformationSpecBuilder) Options(options any) *V2TransformationSpecBuilder {
	builder.internal.Options = options

	return builder
}
