// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ControlSourceRef] = (*ControlSourceRefBuilder)(nil)

type ControlSourceRefBuilder struct {
	internal *ControlSourceRef
	errors   cog.BuildErrors
}

func NewControlSourceRefBuilder() *ControlSourceRefBuilder {
	resource := NewControlSourceRef()
	builder := &ControlSourceRefBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Type = "datasource"

	return builder
}

func (builder *ControlSourceRefBuilder) Build() (ControlSourceRef, error) {
	if err := builder.internal.Validate(); err != nil {
		return ControlSourceRef{}, err
	}

	if len(builder.errors) > 0 {
		return ControlSourceRef{}, cog.MakeBuildErrors("dashboardv2beta1.controlSourceRef", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ControlSourceRefBuilder) RecordError(path string, err error) *ControlSourceRefBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

// The plugin type-id
func (builder *ControlSourceRefBuilder) Group(group string) *ControlSourceRefBuilder {
	builder.internal.Group = group

	return builder
}
