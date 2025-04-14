// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Count] = (*CountBuilder)(nil)

type CountBuilder struct {
	internal *Count
	errors   map[string]cog.BuildErrors
}

func NewCountBuilder() *CountBuilder {
	resource := NewCount()
	builder := &CountBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *CountBuilder) Build() (Count, error) {
	if err := builder.internal.Validate(); err != nil {
		return Count{}, err
	}

	return *builder.internal, nil
}

func (builder *CountBuilder) Id(id string) *CountBuilder {
	builder.internal.Id = id

	return builder
}

func (builder *CountBuilder) Hide(hide bool) *CountBuilder {
	builder.internal.Hide = &hide

	return builder
}
