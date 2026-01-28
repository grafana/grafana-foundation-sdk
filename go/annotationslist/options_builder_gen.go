// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package annotationslist

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
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
		return Options{}, cog.MakeBuildErrors("annotationslist.options", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *OptionsBuilder) OnlyFromThisDashboard(onlyFromThisDashboard bool) *OptionsBuilder {
	builder.internal.OnlyFromThisDashboard = onlyFromThisDashboard

	return builder
}

func (builder *OptionsBuilder) OnlyInTimeRange(onlyInTimeRange bool) *OptionsBuilder {
	builder.internal.OnlyInTimeRange = onlyInTimeRange

	return builder
}

func (builder *OptionsBuilder) Tags(tags []string) *OptionsBuilder {
	builder.internal.Tags = tags

	return builder
}

func (builder *OptionsBuilder) Limit(limit uint32) *OptionsBuilder {
	builder.internal.Limit = limit

	return builder
}

func (builder *OptionsBuilder) ShowUser(showUser bool) *OptionsBuilder {
	builder.internal.ShowUser = showUser

	return builder
}

func (builder *OptionsBuilder) ShowTime(showTime bool) *OptionsBuilder {
	builder.internal.ShowTime = showTime

	return builder
}

func (builder *OptionsBuilder) ShowTags(showTags bool) *OptionsBuilder {
	builder.internal.ShowTags = showTags

	return builder
}

func (builder *OptionsBuilder) NavigateToPanel(navigateToPanel bool) *OptionsBuilder {
	builder.internal.NavigateToPanel = navigateToPanel

	return builder
}

func (builder *OptionsBuilder) NavigateBefore(navigateBefore string) *OptionsBuilder {
	builder.internal.NavigateBefore = navigateBefore

	return builder
}

func (builder *OptionsBuilder) NavigateAfter(navigateAfter string) *OptionsBuilder {
	builder.internal.NavigateAfter = navigateAfter

	return builder
}
