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

func (builder *OptionsBuilder) ShowLogContextToggle(showLogContextToggle bool) *OptionsBuilder {
	builder.internal.ShowLogContextToggle = showLogContextToggle

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

// TODO: figure out how to define callbacks
func (builder *OptionsBuilder) OnClickFilterLabel(onClickFilterLabel any) *OptionsBuilder {
	builder.internal.OnClickFilterLabel = &onClickFilterLabel

	return builder
}

func (builder *OptionsBuilder) OnClickFilterOutLabel(onClickFilterOutLabel any) *OptionsBuilder {
	builder.internal.OnClickFilterOutLabel = &onClickFilterOutLabel

	return builder
}

func (builder *OptionsBuilder) IsFilterLabelActive(isFilterLabelActive any) *OptionsBuilder {
	builder.internal.IsFilterLabelActive = &isFilterLabelActive

	return builder
}

func (builder *OptionsBuilder) OnClickFilterString(onClickFilterString any) *OptionsBuilder {
	builder.internal.OnClickFilterString = &onClickFilterString

	return builder
}

func (builder *OptionsBuilder) OnClickFilterOutString(onClickFilterOutString any) *OptionsBuilder {
	builder.internal.OnClickFilterOutString = &onClickFilterOutString

	return builder
}

func (builder *OptionsBuilder) OnClickShowField(onClickShowField any) *OptionsBuilder {
	builder.internal.OnClickShowField = &onClickShowField

	return builder
}

func (builder *OptionsBuilder) OnClickHideField(onClickHideField any) *OptionsBuilder {
	builder.internal.OnClickHideField = &onClickHideField

	return builder
}

func (builder *OptionsBuilder) DisplayedFields(displayedFields []string) *OptionsBuilder {
	builder.internal.DisplayedFields = displayedFields

	return builder
}
