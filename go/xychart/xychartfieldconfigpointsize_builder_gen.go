// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package xychart

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[XychartFieldConfigPointSize] = (*XychartFieldConfigPointSizeBuilder)(nil)

type XychartFieldConfigPointSizeBuilder struct {
	internal *XychartFieldConfigPointSize
	errors   map[string]cog.BuildErrors
}

func NewXychartFieldConfigPointSizeBuilder() *XychartFieldConfigPointSizeBuilder {
	resource := NewXychartFieldConfigPointSize()
	builder := &XychartFieldConfigPointSizeBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *XychartFieldConfigPointSizeBuilder) Build() (XychartFieldConfigPointSize, error) {
	if err := builder.internal.Validate(); err != nil {
		return XychartFieldConfigPointSize{}, err
	}

	return *builder.internal, nil
}

func (builder *XychartFieldConfigPointSizeBuilder) Fixed(fixed int32) *XychartFieldConfigPointSizeBuilder {
	builder.internal.Fixed = &fixed

	return builder
}

func (builder *XychartFieldConfigPointSizeBuilder) Min(min int32) *XychartFieldConfigPointSizeBuilder {
	builder.internal.Min = &min

	return builder
}

func (builder *XychartFieldConfigPointSizeBuilder) Max(max int32) *XychartFieldConfigPointSizeBuilder {
	builder.internal.Max = &max

	return builder
}
