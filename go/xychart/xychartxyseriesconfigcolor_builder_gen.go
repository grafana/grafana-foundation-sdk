// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package xychart

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[XychartXYSeriesConfigColor] = (*XychartXYSeriesConfigColorBuilder)(nil)

type XychartXYSeriesConfigColorBuilder struct {
	internal *XychartXYSeriesConfigColor
	errors   cog.BuildErrors
}

func NewXychartXYSeriesConfigColorBuilder() *XychartXYSeriesConfigColorBuilder {
	resource := NewXychartXYSeriesConfigColor()
	builder := &XychartXYSeriesConfigColorBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *XychartXYSeriesConfigColorBuilder) Build() (XychartXYSeriesConfigColor, error) {
	if err := builder.internal.Validate(); err != nil {
		return XychartXYSeriesConfigColor{}, err
	}

	if len(builder.errors) > 0 {
		return XychartXYSeriesConfigColor{}, cog.MakeBuildErrors("xychart.xychartXYSeriesConfigColor", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *XychartXYSeriesConfigColorBuilder) Matcher(matcher cog.Builder[MatcherConfig]) *XychartXYSeriesConfigColorBuilder {
	matcherResource, err := matcher.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Matcher = matcherResource

	return builder
}
