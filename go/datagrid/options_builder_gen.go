// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package datagrid

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
		return Options{}, cog.MakeBuildErrors("datagrid.options", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *OptionsBuilder) SelectedSeries(selectedSeries int32) *OptionsBuilder {
	builder.internal.SelectedSeries = selectedSeries

	return builder
}
