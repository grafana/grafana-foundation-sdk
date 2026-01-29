// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ElementReference] = (*ElementReferenceBuilder)(nil)

type ElementReferenceBuilder struct {
	internal *ElementReference
	errors   cog.BuildErrors
}

func NewElementReferenceBuilder() *ElementReferenceBuilder {
	resource := NewElementReference()
	builder := &ElementReferenceBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "ElementReference"

	return builder
}

func (builder *ElementReferenceBuilder) Build() (ElementReference, error) {
	if err := builder.internal.Validate(); err != nil {
		return ElementReference{}, err
	}

	if len(builder.errors) > 0 {
		return ElementReference{}, cog.MakeBuildErrors("dashboardv2beta1.elementReference", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ElementReferenceBuilder) Name(name string) *ElementReferenceBuilder {
	builder.internal.Name = name

	return builder
}
