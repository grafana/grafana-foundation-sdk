// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package xychart

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[XychartXYSeriesConfigFrame] = (*XychartXYSeriesConfigFrameBuilder)(nil)

type XychartXYSeriesConfigFrameBuilder struct {
	internal *XychartXYSeriesConfigFrame
	errors   map[string]cog.BuildErrors
}

func NewXychartXYSeriesConfigFrameBuilder() *XychartXYSeriesConfigFrameBuilder {
	resource := NewXychartXYSeriesConfigFrame()
	builder := &XychartXYSeriesConfigFrameBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *XychartXYSeriesConfigFrameBuilder) Build() (XychartXYSeriesConfigFrame, error) {
	if err := builder.internal.Validate(); err != nil {
		return XychartXYSeriesConfigFrame{}, err
	}

	return *builder.internal, nil
}

func (builder *XychartXYSeriesConfigFrameBuilder) Matcher(matcher cog.Builder[MatcherConfig]) *XychartXYSeriesConfigFrameBuilder {
	matcherResource, err := matcher.Build()
	if err != nil {
		builder.errors["matcher"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Matcher = matcherResource

	return builder
}
