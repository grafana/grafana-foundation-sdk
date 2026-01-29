// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[TabRepeatOptions] = (*TabRepeatOptionsBuilder)(nil)

type TabRepeatOptionsBuilder struct {
	internal *TabRepeatOptions
	errors   cog.BuildErrors
}

func NewTabRepeatOptionsBuilder() *TabRepeatOptionsBuilder {
	resource := NewTabRepeatOptions()
	builder := &TabRepeatOptionsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *TabRepeatOptionsBuilder) Build() (TabRepeatOptions, error) {
	if err := builder.internal.Validate(); err != nil {
		return TabRepeatOptions{}, err
	}

	if len(builder.errors) > 0 {
		return TabRepeatOptions{}, cog.MakeBuildErrors("dashboardv2beta1.tabRepeatOptions", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *TabRepeatOptionsBuilder) Value(value string) *TabRepeatOptionsBuilder {
	builder.internal.Value = value

	return builder
}
