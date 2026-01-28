// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alertgroups

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
		return Options{}, cog.MakeBuildErrors("alertgroups.options", builder.errors)
	}

	return *builder.internal, nil
}

// Comma-separated list of values used to filter alert results
func (builder *OptionsBuilder) Labels(labels string) *OptionsBuilder {
	builder.internal.Labels = labels

	return builder
}

// Name of the alertmanager used as a source for alerts
func (builder *OptionsBuilder) Alertmanager(alertmanager string) *OptionsBuilder {
	builder.internal.Alertmanager = alertmanager

	return builder
}

// Expand all alert groups by default
func (builder *OptionsBuilder) ExpandAll(expandAll bool) *OptionsBuilder {
	builder.internal.ExpandAll = expandAll

	return builder
}
