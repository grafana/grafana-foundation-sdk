// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package xychart

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[XychartXYSeriesConfigSize] = (*XychartXYSeriesConfigSizeBuilder)(nil)

type XychartXYSeriesConfigSizeBuilder struct {
	internal *XychartXYSeriesConfigSize
	errors   map[string]cog.BuildErrors
}

func NewXychartXYSeriesConfigSizeBuilder() *XychartXYSeriesConfigSizeBuilder {
	resource := NewXychartXYSeriesConfigSize()
	builder := &XychartXYSeriesConfigSizeBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *XychartXYSeriesConfigSizeBuilder) Build() (XychartXYSeriesConfigSize, error) {
	if err := builder.internal.Validate(); err != nil {
		return XychartXYSeriesConfigSize{}, err
	}

	return *builder.internal, nil
}

func (builder *XychartXYSeriesConfigSizeBuilder) Matcher(matcher cog.Builder[MatcherConfig]) *XychartXYSeriesConfigSizeBuilder {
	matcherResource, err := matcher.Build()
	if err != nil {
		builder.errors["matcher"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Matcher = matcherResource

	return builder
}
