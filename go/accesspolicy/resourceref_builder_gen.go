// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package accesspolicy

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ResourceRef] = (*ResourceRefBuilder)(nil)

type ResourceRefBuilder struct {
	internal *ResourceRef
	errors   map[string]cog.BuildErrors
}

func NewResourceRefBuilder() *ResourceRefBuilder {
	resource := &ResourceRef{}
	builder := &ResourceRefBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *ResourceRefBuilder) Build() (ResourceRef, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("ResourceRef", err)...)
	}

	if len(errs) != 0 {
		return ResourceRef{}, errs
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

func (builder *ResourceRefBuilder) applyDefaults() {
}
