// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package xychart

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[XychartXYSeriesConfigX] = (*XychartXYSeriesConfigXBuilder)(nil)

type XychartXYSeriesConfigXBuilder struct {
	internal *XychartXYSeriesConfigX
	errors   cog.BuildErrors
}

func NewXychartXYSeriesConfigXBuilder() *XychartXYSeriesConfigXBuilder {
	resource := NewXychartXYSeriesConfigX()
	builder := &XychartXYSeriesConfigXBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *XychartXYSeriesConfigXBuilder) Build() (XychartXYSeriesConfigX, error) {
	if err := builder.internal.Validate(); err != nil {
		return XychartXYSeriesConfigX{}, err
	}

	if len(builder.errors) > 0 {
		return XychartXYSeriesConfigX{}, cog.MakeBuildErrors("xychart.xychartXYSeriesConfigX", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *XychartXYSeriesConfigXBuilder) Matcher(matcher cog.Builder[MatcherConfig]) *XychartXYSeriesConfigXBuilder {
	matcherResource, err := matcher.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Matcher = matcherResource

	return builder
}
