// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Sum] = (*SumBuilder)(nil)

type SumBuilder struct {
	internal *Sum
	errors   map[string]cog.BuildErrors
}

func NewSumBuilder() *SumBuilder {
	resource := &Sum{}
	builder := &SumBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Type = "sum"

	return builder
}

func (builder *SumBuilder) Build() (Sum, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("Sum", err)...)
	}

	if len(errs) != 0 {
		return Sum{}, errs
	}

	return *builder.internal, nil
}

func (builder *SumBuilder) Field(field string) *SumBuilder {
	builder.internal.Field = &field

	return builder
}

func (builder *SumBuilder) Id(id string) *SumBuilder {
	builder.internal.Id = id

	return builder
}

func (builder *SumBuilder) Settings(settings struct {
	Script  *InlineScript `json:"script,omitempty"`
	Missing *string       `json:"missing,omitempty"`
}) *SumBuilder {
	builder.internal.Settings = &settings

	return builder
}

func (builder *SumBuilder) Hide(hide bool) *SumBuilder {
	builder.internal.Hide = &hide

	return builder
}

func (builder *SumBuilder) applyDefaults() {
}
