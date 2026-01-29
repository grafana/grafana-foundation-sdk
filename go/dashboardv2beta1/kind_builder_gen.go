// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Kind] = (*KindBuilder)(nil)

// --- Common types ---
type KindBuilder struct {
	internal *Kind
	errors   cog.BuildErrors
}

func NewKindBuilder() *KindBuilder {
	resource := NewKind()
	builder := &KindBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *KindBuilder) Build() (Kind, error) {
	if err := builder.internal.Validate(); err != nil {
		return Kind{}, err
	}

	if len(builder.errors) > 0 {
		return Kind{}, cog.MakeBuildErrors("dashboardv2beta1.kind", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *KindBuilder) Kind(kind string) *KindBuilder {
	builder.internal.Kind = kind

	return builder
}

func (builder *KindBuilder) Spec(spec any) *KindBuilder {
	builder.internal.Spec = spec

	return builder
}

func (builder *KindBuilder) Metadata(metadata any) *KindBuilder {
	builder.internal.Metadata = &metadata

	return builder
}
