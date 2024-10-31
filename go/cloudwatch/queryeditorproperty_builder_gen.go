// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package cloudwatch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[QueryEditorProperty] = (*QueryEditorPropertyBuilder)(nil)

type QueryEditorPropertyBuilder struct {
	internal *QueryEditorProperty
	errors   map[string]cog.BuildErrors
}

func NewQueryEditorPropertyBuilder() *QueryEditorPropertyBuilder {
	resource := &QueryEditorProperty{}
	builder := &QueryEditorPropertyBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *QueryEditorPropertyBuilder) Build() (QueryEditorProperty, error) {
	if err := builder.internal.Validate(); err != nil {
		return QueryEditorProperty{}, err
	}

	return *builder.internal, nil
}

func (builder *QueryEditorPropertyBuilder) Type(typeArg QueryEditorPropertyType) *QueryEditorPropertyBuilder {
	builder.internal.Type = typeArg

	return builder
}

func (builder *QueryEditorPropertyBuilder) Name(name string) *QueryEditorPropertyBuilder {
	builder.internal.Name = &name

	return builder
}

func (builder *QueryEditorPropertyBuilder) applyDefaults() {
}
