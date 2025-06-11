// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package xychart

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[XychartXYSeriesConfigY] = (*XychartXYSeriesConfigYBuilder)(nil)

type XychartXYSeriesConfigYBuilder struct {
	internal *XychartXYSeriesConfigY
	errors   cog.BuildErrors
}

func NewXychartXYSeriesConfigYBuilder() *XychartXYSeriesConfigYBuilder {
	resource := NewXychartXYSeriesConfigY()
	builder := &XychartXYSeriesConfigYBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *XychartXYSeriesConfigYBuilder) Build() (XychartXYSeriesConfigY, error) {
	if err := builder.internal.Validate(); err != nil {
		return XychartXYSeriesConfigY{}, err
	}

	if len(builder.errors) > 0 {
		return XychartXYSeriesConfigY{}, cog.MakeBuildErrors("xychart.xychartXYSeriesConfigY", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *XychartXYSeriesConfigYBuilder) Matcher(matcher cog.Builder[MatcherConfig]) *XychartXYSeriesConfigYBuilder {
	matcherResource, err := matcher.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Matcher = matcherResource

	return builder
}
