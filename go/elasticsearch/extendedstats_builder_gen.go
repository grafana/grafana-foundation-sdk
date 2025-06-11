// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ExtendedStats] = (*ExtendedStatsBuilder)(nil)

type ExtendedStatsBuilder struct {
	internal *ExtendedStats
	errors   cog.BuildErrors
}

func NewExtendedStatsBuilder() *ExtendedStatsBuilder {
	resource := NewExtendedStats()
	builder := &ExtendedStatsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ExtendedStatsBuilder) Build() (ExtendedStats, error) {
	if err := builder.internal.Validate(); err != nil {
		return ExtendedStats{}, err
	}

	if len(builder.errors) > 0 {
		return ExtendedStats{}, cog.MakeBuildErrors("elasticsearch.extendedStats", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ExtendedStatsBuilder) Settings(settings cog.Builder[ElasticsearchExtendedStatsSettings]) *ExtendedStatsBuilder {
	settingsResource, err := settings.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Settings = &settingsResource

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
