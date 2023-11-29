// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Nested] = (*NestedBuilder)(nil)

type NestedBuilder struct {
	internal *Nested
	errors   map[string]cog.BuildErrors
}

func NewNestedBuilder() *NestedBuilder {
	resource := &Nested{}
	builder := &NestedBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Type = "nested"

	return builder
}

func (builder *NestedBuilder) Build() (Nested, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("Nested", err)...)
	}

	if len(errs) != 0 {
		return Nested{}, errs
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

func (builder *NestedBuilder) applyDefaults() {
}
