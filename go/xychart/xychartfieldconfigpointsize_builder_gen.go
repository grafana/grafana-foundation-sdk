// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package xychart

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[XychartFieldConfigPointSize] = (*XychartFieldConfigPointSizeBuilder)(nil)

type XychartFieldConfigPointSizeBuilder struct {
	internal *XychartFieldConfigPointSize
	errors   cog.BuildErrors
}

func NewXychartFieldConfigPointSizeBuilder() *XychartFieldConfigPointSizeBuilder {
	resource := NewXychartFieldConfigPointSize()
	builder := &XychartFieldConfigPointSizeBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *XychartFieldConfigPointSizeBuilder) Build() (XychartFieldConfigPointSize, error) {
	if err := builder.internal.Validate(); err != nil {
		return XychartFieldConfigPointSize{}, err
	}

	if len(builder.errors) > 0 {
		return XychartFieldConfigPointSize{}, cog.MakeBuildErrors("xychart.xychartFieldConfigPointSize", builder.errors)
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
