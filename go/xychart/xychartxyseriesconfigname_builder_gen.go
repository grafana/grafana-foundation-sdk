// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package xychart

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[XychartXYSeriesConfigName] = (*XychartXYSeriesConfigNameBuilder)(nil)

type XychartXYSeriesConfigNameBuilder struct {
	internal *XychartXYSeriesConfigName
	errors   cog.BuildErrors
}

func NewXychartXYSeriesConfigNameBuilder() *XychartXYSeriesConfigNameBuilder {
	resource := NewXychartXYSeriesConfigName()
	builder := &XychartXYSeriesConfigNameBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *XychartXYSeriesConfigNameBuilder) Build() (XychartXYSeriesConfigName, error) {
	if err := builder.internal.Validate(); err != nil {
		return XychartXYSeriesConfigName{}, err
	}

	if len(builder.errors) > 0 {
		return XychartXYSeriesConfigName{}, cog.MakeBuildErrors("xychart.xychartXYSeriesConfigName", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *XychartXYSeriesConfigNameBuilder) RecordError(path string, err error) *XychartXYSeriesConfigNameBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *XychartXYSeriesConfigNameBuilder) Fixed(fixed string) *XychartXYSeriesConfigNameBuilder {
	builder.internal.Fixed = &fixed

	return builder
}
