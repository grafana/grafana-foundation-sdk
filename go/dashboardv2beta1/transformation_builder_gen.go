// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[TransformationKind] = (*TransformationBuilder)(nil)

type TransformationBuilder struct {
	internal *TransformationKind
	errors   cog.BuildErrors
}

func NewTransformationBuilder() *TransformationBuilder {
	resource := NewTransformationKind()
	builder := &TransformationBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *TransformationBuilder) Build() (TransformationKind, error) {
	if err := builder.internal.Validate(); err != nil {
		return TransformationKind{}, err
	}

	if len(builder.errors) > 0 {
		return TransformationKind{}, cog.MakeBuildErrors("dashboardv2beta1.transformation", builder.errors)
	}

	return *builder.internal, nil
}

// The kind of a TransformationKind is the transformation ID
func (builder *TransformationBuilder) Kind(kind string) *TransformationBuilder {
	builder.internal.Kind = kind

	return builder
}

// Unique identifier of transformer
func (builder *TransformationBuilder) Id(id string) *TransformationBuilder {
	builder.internal.Spec.Id = id

	return builder
}

// Disabled transformations are skipped
func (builder *TransformationBuilder) Disabled(disabled bool) *TransformationBuilder {
	builder.internal.Spec.Disabled = &disabled

	return builder
}

// Optional frame matcher. When missing it will be applied to all results
func (builder *TransformationBuilder) Filter(filter MatcherConfig) *TransformationBuilder {
	builder.internal.Spec.Filter = &filter

	return builder
}

// Where to pull DataFrames from as input to transformation
func (builder *TransformationBuilder) Topic(topic DataTopic) *TransformationBuilder {
	builder.internal.Spec.Topic = &topic

	return builder
}

// Options to be passed to the transformer
// Valid options depend on the transformer id
func (builder *TransformationBuilder) Options(options any) *TransformationBuilder {
	builder.internal.Spec.Options = options

	return builder
}
