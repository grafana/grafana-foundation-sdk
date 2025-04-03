// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package xychart

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[XychartXYSeriesConfigName] = (*XychartXYSeriesConfigNameBuilder)(nil)

type XychartXYSeriesConfigNameBuilder struct {
	internal *XychartXYSeriesConfigName
	errors   map[string]cog.BuildErrors
}

func NewXychartXYSeriesConfigNameBuilder() *XychartXYSeriesConfigNameBuilder {
	resource := NewXychartXYSeriesConfigName()
	builder := &XychartXYSeriesConfigNameBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *XychartXYSeriesConfigNameBuilder) Build() (XychartXYSeriesConfigName, error) {
	if err := builder.internal.Validate(); err != nil {
		return XychartXYSeriesConfigName{}, err
	}

	return *builder.internal, nil
}

func (builder *XychartXYSeriesConfigNameBuilder) Fixed(fixed string) *XychartXYSeriesConfigNameBuilder {
	builder.internal.Fixed = &fixed

	return builder
}
