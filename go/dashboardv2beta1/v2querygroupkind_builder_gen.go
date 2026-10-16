// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[V2QueryGroupKind] = (*V2QueryGroupKindBuilder)(nil)

type V2QueryGroupKindBuilder struct {
	internal *V2QueryGroupKind
	errors   cog.BuildErrors
}

func NewV2QueryGroupKindBuilder() *V2QueryGroupKindBuilder {
	resource := NewV2QueryGroupKind()
	builder := &V2QueryGroupKindBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "QueryGroup"

	return builder
}

func (builder *V2QueryGroupKindBuilder) Build() (V2QueryGroupKind, error) {
	if err := builder.internal.Validate(); err != nil {
		return V2QueryGroupKind{}, err
	}

	if len(builder.errors) > 0 {
		return V2QueryGroupKind{}, cog.MakeBuildErrors("dashboardv2beta1.v2QueryGroupKind", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *V2QueryGroupKindBuilder) RecordError(path string, err error) *V2QueryGroupKindBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *V2QueryGroupKindBuilder) Spec(spec cog.Builder[V2QueryGroupSpec]) *V2QueryGroupKindBuilder {
	specResource, err := spec.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec = specResource

	return builder
}
