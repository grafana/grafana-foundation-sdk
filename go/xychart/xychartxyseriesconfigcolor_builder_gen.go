// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package xychart

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[XychartXYSeriesConfigColor] = (*XychartXYSeriesConfigColorBuilder)(nil)

type XychartXYSeriesConfigColorBuilder struct {
	internal *XychartXYSeriesConfigColor
	errors   map[string]cog.BuildErrors
}

func NewXychartXYSeriesConfigColorBuilder() *XychartXYSeriesConfigColorBuilder {
	resource := NewXychartXYSeriesConfigColor()
	builder := &XychartXYSeriesConfigColorBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *XychartXYSeriesConfigColorBuilder) Build() (XychartXYSeriesConfigColor, error) {
	if err := builder.internal.Validate(); err != nil {
		return XychartXYSeriesConfigColor{}, err
	}

	return *builder.internal, nil
}

func (builder *XychartXYSeriesConfigColorBuilder) Matcher(matcher cog.Builder[MatcherConfig]) *XychartXYSeriesConfigColorBuilder {
	matcherResource, err := matcher.Build()
	if err != nil {
		builder.errors["matcher"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Matcher = matcherResource

	return builder
}
