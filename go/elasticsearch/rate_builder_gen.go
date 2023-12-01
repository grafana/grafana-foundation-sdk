// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Rate] = (*RateBuilder)(nil)

type RateBuilder struct {
	internal *Rate
	errors   map[string]cog.BuildErrors
}

func NewRateBuilder() *RateBuilder {
	resource := &Rate{}
	builder := &RateBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Type = "rate"

	return builder
}

func (builder *RateBuilder) Build() (Rate, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("Rate", err)...)
	}

	if len(errs) != 0 {
		return Rate{}, errs
	}

	return *builder.internal, nil
}

func (builder *RateBuilder) Field(field string) *RateBuilder {
	builder.internal.Field = &field

	return builder
}

func (builder *RateBuilder) Id(id string) *RateBuilder {
	builder.internal.Id = id

	return builder
}

func (builder *RateBuilder) Settings(settings struct {
	Unit *string `json:"unit,omitempty"`
	Mode *string `json:"mode,omitempty"`
}) *RateBuilder {
	builder.internal.Settings = &settings

	return builder
}

func (builder *RateBuilder) Hide(hide bool) *RateBuilder {
	builder.internal.Hide = &hide

	return builder
}

func (builder *RateBuilder) applyDefaults() {
}
