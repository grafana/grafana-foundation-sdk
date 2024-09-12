// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Percentiles] = (*PercentilesBuilder)(nil)

type PercentilesBuilder struct {
	internal *Percentiles
	errors   map[string]cog.BuildErrors
}

func NewPercentilesBuilder() *PercentilesBuilder {
	resource := &Percentiles{}
	builder := &PercentilesBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Type = "percentiles"

	return builder
}

func (builder *PercentilesBuilder) Build() (Percentiles, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("Percentiles", err)...)
	}

	if len(errs) != 0 {
		return Percentiles{}, errs
	}

	return *builder.internal, nil
}

func (builder *PercentilesBuilder) Field(field string) *PercentilesBuilder {
	builder.internal.Field = &field

	return builder
}

func (builder *PercentilesBuilder) Id(id string) *PercentilesBuilder {
	builder.internal.Id = id

	return builder
}

func (builder *PercentilesBuilder) Settings(settings struct {
	Script   *InlineScript `json:"script,omitempty"`
	Missing  *string       `json:"missing,omitempty"`
	Percents []string      `json:"percents,omitempty"`
}) *PercentilesBuilder {
	builder.internal.Settings = &settings

	return builder
}

func (builder *PercentilesBuilder) Hide(hide bool) *PercentilesBuilder {
	builder.internal.Hide = &hide

	return builder
}

func (builder *PercentilesBuilder) applyDefaults() {
}
