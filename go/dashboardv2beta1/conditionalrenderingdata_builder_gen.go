// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ConditionalRenderingDataKind] = (*ConditionalRenderingDataBuilder)(nil)

type ConditionalRenderingDataBuilder struct {
	internal *ConditionalRenderingDataKind
	errors   cog.BuildErrors
}

func NewConditionalRenderingDataBuilder() *ConditionalRenderingDataBuilder {
	resource := NewConditionalRenderingDataKind()
	builder := &ConditionalRenderingDataBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "ConditionalRenderingData"

	return builder
}

func (builder *ConditionalRenderingDataBuilder) Build() (ConditionalRenderingDataKind, error) {
	if err := builder.internal.Validate(); err != nil {
		return ConditionalRenderingDataKind{}, err
	}

	if len(builder.errors) > 0 {
		return ConditionalRenderingDataKind{}, cog.MakeBuildErrors("dashboardv2beta1.conditionalRenderingData", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ConditionalRenderingDataBuilder) RecordError(path string, err error) *ConditionalRenderingDataBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *ConditionalRenderingDataBuilder) Value(value bool) *ConditionalRenderingDataBuilder {
	builder.internal.Spec.Value = value

	return builder
}
