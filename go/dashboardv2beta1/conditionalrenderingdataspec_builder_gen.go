// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ConditionalRenderingDataSpec] = (*ConditionalRenderingDataSpecBuilder)(nil)

type ConditionalRenderingDataSpecBuilder struct {
	internal *ConditionalRenderingDataSpec
	errors   cog.BuildErrors
}

func NewConditionalRenderingDataSpecBuilder() *ConditionalRenderingDataSpecBuilder {
	resource := NewConditionalRenderingDataSpec()
	builder := &ConditionalRenderingDataSpecBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ConditionalRenderingDataSpecBuilder) Build() (ConditionalRenderingDataSpec, error) {
	if err := builder.internal.Validate(); err != nil {
		return ConditionalRenderingDataSpec{}, err
	}

	if len(builder.errors) > 0 {
		return ConditionalRenderingDataSpec{}, cog.MakeBuildErrors("dashboardv2beta1.conditionalRenderingDataSpec", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ConditionalRenderingDataSpecBuilder) RecordError(path string, err error) *ConditionalRenderingDataSpecBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *ConditionalRenderingDataSpecBuilder) Value(value bool) *ConditionalRenderingDataSpecBuilder {
	builder.internal.Value = value

	return builder
}
