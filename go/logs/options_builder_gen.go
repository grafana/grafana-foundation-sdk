// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package logs

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

var _ cog.Builder[Options] = (*OptionsBuilder)(nil)

type OptionsBuilder struct {
	internal *Options
	errors   cog.BuildErrors
}

func NewOptionsBuilder() *OptionsBuilder {
	resource := NewOptions()
	builder := &OptionsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *OptionsBuilder) Build() (Options, error) {
	if err := builder.internal.Validate(); err != nil {
		return Options{}, err
	}

	if len(builder.errors) > 0 {
		return Options{}, cog.MakeBuildErrors("logs.options", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *OptionsBuilder) ShowLabels(showLabels bool) *OptionsBuilder {
	builder.internal.ShowLabels = showLabels

	return builder
}

func (builder *OptionsBuilder) ShowCommonLabels(showCommonLabels bool) *OptionsBuilder {
	builder.internal.ShowCommonLabels = showCommonLabels

	return builder
}

func (builder *OptionsBuilder) ShowTime(showTime bool) *OptionsBuilder {
	builder.internal.ShowTime = showTime

	return builder
}

func (builder *OptionsBuilder) WrapLogMessage(wrapLogMessage bool) *OptionsBuilder {
	builder.internal.WrapLogMessage = wrapLogMessage

	return builder
}

func (builder *OptionsBuilder) PrettifyLogMessage(prettifyLogMessage bool) *OptionsBuilder {
	builder.internal.PrettifyLogMessage = prettifyLogMessage

	return builder
}

func (builder *OptionsBuilder) EnableLogDetails(enableLogDetails bool) *OptionsBuilder {
	builder.internal.EnableLogDetails = enableLogDetails

	return builder
}

func (builder *OptionsBuilder) SortOrder(sortOrder common.LogsSortOrder) *OptionsBuilder {
	builder.internal.SortOrder = sortOrder

	return builder
}

func (builder *OptionsBuilder) DedupStrategy(dedupStrategy common.LogsDedupStrategy) *OptionsBuilder {
	builder.internal.DedupStrategy = dedupStrategy

	return builder
}
