// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package xychart

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[XychartXYSeriesConfigFrame] = (*XychartXYSeriesConfigFrameBuilder)(nil)

type XychartXYSeriesConfigFrameBuilder struct {
	internal *XychartXYSeriesConfigFrame
	errors   cog.BuildErrors
}

func NewXychartXYSeriesConfigFrameBuilder() *XychartXYSeriesConfigFrameBuilder {
	resource := NewXychartXYSeriesConfigFrame()
	builder := &XychartXYSeriesConfigFrameBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *XychartXYSeriesConfigFrameBuilder) Build() (XychartXYSeriesConfigFrame, error) {
	if err := builder.internal.Validate(); err != nil {
		return XychartXYSeriesConfigFrame{}, err
	}

	if len(builder.errors) > 0 {
		return XychartXYSeriesConfigFrame{}, cog.MakeBuildErrors("xychart.xychartXYSeriesConfigFrame", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *XychartXYSeriesConfigFrameBuilder) RecordError(path string, err error) *XychartXYSeriesConfigFrameBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *XychartXYSeriesConfigFrameBuilder) Matcher(matcher cog.Builder[MatcherConfig]) *XychartXYSeriesConfigFrameBuilder {
	matcherResource, err := matcher.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Matcher = matcherResource

	return builder
}
