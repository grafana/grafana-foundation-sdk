// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[BuilderQueryEditorProperty] = (*BuilderQueryEditorPropertyBuilder)(nil)

type BuilderQueryEditorPropertyBuilder struct {
	internal *BuilderQueryEditorProperty
	errors   map[string]cog.BuildErrors
}

func NewBuilderQueryEditorPropertyBuilder() *BuilderQueryEditorPropertyBuilder {
	resource := NewBuilderQueryEditorProperty()
	builder := &BuilderQueryEditorPropertyBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *BuilderQueryEditorPropertyBuilder) Build() (BuilderQueryEditorProperty, error) {
	if err := builder.internal.Validate(); err != nil {
		return BuilderQueryEditorProperty{}, err
	}

	return *builder.internal, nil
}

func (builder *BuilderQueryEditorPropertyBuilder) Type(typeArg BuilderQueryEditorPropertyType) *BuilderQueryEditorPropertyBuilder {
	builder.internal.Type = typeArg

	return builder
}

func (builder *BuilderQueryEditorPropertyBuilder) Name(name string) *BuilderQueryEditorPropertyBuilder {
	builder.internal.Name = name

	return builder
}
