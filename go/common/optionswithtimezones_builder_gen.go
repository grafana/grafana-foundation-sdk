// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package common

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[OptionsWithTimezones] = (*OptionsWithTimezonesBuilder)(nil)

// TODO docs
type OptionsWithTimezonesBuilder struct {
	internal *OptionsWithTimezones
	errors   cog.BuildErrors
}

func NewOptionsWithTimezonesBuilder() *OptionsWithTimezonesBuilder {
	resource := NewOptionsWithTimezones()
	builder := &OptionsWithTimezonesBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *OptionsWithTimezonesBuilder) Build() (OptionsWithTimezones, error) {
	if err := builder.internal.Validate(); err != nil {
		return OptionsWithTimezones{}, err
	}

	if len(builder.errors) > 0 {
		return OptionsWithTimezones{}, cog.MakeBuildErrors("common.optionsWithTimezones", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *OptionsWithTimezonesBuilder) Timezone(timezone []TimeZone) *OptionsWithTimezonesBuilder {
	builder.internal.Timezone = timezone

	return builder
}
