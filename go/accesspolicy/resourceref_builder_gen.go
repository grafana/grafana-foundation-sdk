// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package accesspolicy

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ResourceRef] = (*ResourceRefBuilder)(nil)

type ResourceRefBuilder struct {
	internal *ResourceRef
	errors   cog.BuildErrors
}

func NewResourceRefBuilder() *ResourceRefBuilder {
	resource := NewResourceRef()
	builder := &ResourceRefBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ResourceRefBuilder) Build() (ResourceRef, error) {
	if err := builder.internal.Validate(); err != nil {
		return ResourceRef{}, err
	}

	if len(builder.errors) > 0 {
		return ResourceRef{}, cog.MakeBuildErrors("accesspolicy.resourceRef", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ResourceRefBuilder) Kind(kind string) *ResourceRefBuilder {
	builder.internal.Kind = kind

	return builder
}

func (builder *ResourceRefBuilder) Name(name string) *ResourceRefBuilder {
	builder.internal.Name = name

	return builder
}
