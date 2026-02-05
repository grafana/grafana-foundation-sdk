// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Count] = (*CountBuilder)(nil)

type CountBuilder struct {
	internal *Count
	errors   cog.BuildErrors
}

func NewCountBuilder() *CountBuilder {
	resource := NewCount()
	builder := &CountBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *CountBuilder) Build() (Count, error) {
	if err := builder.internal.Validate(); err != nil {
		return Count{}, err
	}

	if len(builder.errors) > 0 {
		return Count{}, cog.MakeBuildErrors("elasticsearch.count", builder.errors)
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
