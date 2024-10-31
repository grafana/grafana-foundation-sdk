// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package nodegraph

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ArcOption] = (*ArcOptionBuilder)(nil)

type ArcOptionBuilder struct {
	internal *ArcOption
	errors   map[string]cog.BuildErrors
}

func NewArcOptionBuilder() *ArcOptionBuilder {
	resource := &ArcOption{}
	builder := &ArcOptionBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *ArcOptionBuilder) Build() (ArcOption, error) {
	if err := builder.internal.Validate(); err != nil {
		return ArcOption{}, err
	}

	return *builder.internal, nil
}

// Field from which to get the value. Values should be less than 1, representing fraction of a circle.
func (builder *ArcOptionBuilder) Field(field string) *ArcOptionBuilder {
	builder.internal.Field = &field

	return builder
}

// The color of the arc.
func (builder *ArcOptionBuilder) Color(color string) *ArcOptionBuilder {
	builder.internal.Color = &color

	return builder
}

func (builder *ArcOptionBuilder) applyDefaults() {
}
