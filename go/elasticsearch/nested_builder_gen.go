// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Nested] = (*NestedBuilder)(nil)

type NestedBuilder struct {
	internal *Nested
	errors   cog.BuildErrors
}

func NewNestedBuilder() *NestedBuilder {
	resource := NewNested()
	builder := &NestedBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *NestedBuilder) Build() (Nested, error) {
	if err := builder.internal.Validate(); err != nil {
		return Nested{}, err
	}

	if len(builder.errors) > 0 {
		return Nested{}, cog.MakeBuildErrors("elasticsearch.nested", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *NestedBuilder) Field(field string) *NestedBuilder {
	builder.internal.Field = &field

	return builder
}

func (builder *NestedBuilder) Id(id string) *NestedBuilder {
	builder.internal.Id = id

	return builder
}

func (builder *NestedBuilder) Settings(settings any) *NestedBuilder {
	builder.internal.Settings = &settings

	return builder
}
