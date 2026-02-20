// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package starsv1alpha1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Resource] = (*ResourceBuilder)(nil)

type ResourceBuilder struct {
	internal *Resource
	errors   cog.BuildErrors
}

func NewResourceBuilder() *ResourceBuilder {
	resource := NewResource()
	builder := &ResourceBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ResourceBuilder) Build() (Resource, error) {
	if err := builder.internal.Validate(); err != nil {
		return Resource{}, err
	}

	if len(builder.errors) > 0 {
		return Resource{}, cog.MakeBuildErrors("starsv1alpha1.resource", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ResourceBuilder) RecordError(path string, err error) *ResourceBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *ResourceBuilder) Group(group string) *ResourceBuilder {
	builder.internal.Group = group

	return builder
}

func (builder *ResourceBuilder) Kind(kind string) *ResourceBuilder {
	builder.internal.Kind = kind

	return builder
}

// The set of resources
// +listType=set
func (builder *ResourceBuilder) Names(names []string) *ResourceBuilder {
	builder.internal.Names = names

	return builder
}
