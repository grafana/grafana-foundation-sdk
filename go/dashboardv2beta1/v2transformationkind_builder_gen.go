// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[V2TransformationKind] = (*V2TransformationKindBuilder)(nil)

// Dashboard v2 shape: the transformation ID moved from `kind` to `group`.
type V2TransformationKindBuilder struct {
	internal *V2TransformationKind
	errors   cog.BuildErrors
}

func NewV2TransformationKindBuilder() *V2TransformationKindBuilder {
	resource := NewV2TransformationKind()
	builder := &V2TransformationKindBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "Transformation"

	return builder
}

func (builder *V2TransformationKindBuilder) Build() (V2TransformationKind, error) {
	if err := builder.internal.Validate(); err != nil {
		return V2TransformationKind{}, err
	}

	if len(builder.errors) > 0 {
		return V2TransformationKind{}, cog.MakeBuildErrors("dashboardv2beta1.v2TransformationKind", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *V2TransformationKindBuilder) RecordError(path string, err error) *V2TransformationKindBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

// The group is the transformation ID
func (builder *V2TransformationKindBuilder) Group(group string) *V2TransformationKindBuilder {
	builder.internal.Group = group

	return builder
}

func (builder *V2TransformationKindBuilder) Spec(spec cog.Builder[V2TransformationSpec]) *V2TransformationKindBuilder {
	specResource, err := spec.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec = specResource

	return builder
}
