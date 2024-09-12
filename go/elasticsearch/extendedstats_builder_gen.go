// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ExtendedStats] = (*ExtendedStatsBuilder)(nil)

type ExtendedStatsBuilder struct {
	internal *ExtendedStats
	errors   map[string]cog.BuildErrors
}

func NewExtendedStatsBuilder() *ExtendedStatsBuilder {
	resource := &ExtendedStats{}
	builder := &ExtendedStatsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Type = "extended_stats"

	return builder
}

func (builder *ExtendedStatsBuilder) Build() (ExtendedStats, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("ExtendedStats", err)...)
	}

	if len(errs) != 0 {
		return ExtendedStats{}, errs
	}

	return *builder.internal, nil
}

func (builder *ExtendedStatsBuilder) Settings(settings struct {
	Script  *InlineScript `json:"script,omitempty"`
	Missing *string       `json:"missing,omitempty"`
	Sigma   *string       `json:"sigma,omitempty"`
}) *ExtendedStatsBuilder {
	builder.internal.Settings = &settings

	return builder
}

func (builder *ExtendedStatsBuilder) Field(field string) *ExtendedStatsBuilder {
	builder.internal.Field = &field

	return builder
}

func (builder *ExtendedStatsBuilder) Id(id string) *ExtendedStatsBuilder {
	builder.internal.Id = id

	return builder
}

func (builder *ExtendedStatsBuilder) Meta(meta any) *ExtendedStatsBuilder {
	builder.internal.Meta = &meta

	return builder
}

func (builder *ExtendedStatsBuilder) Hide(hide bool) *ExtendedStatsBuilder {
	builder.internal.Hide = &hide

	return builder
}

func (builder *ExtendedStatsBuilder) applyDefaults() {
}
