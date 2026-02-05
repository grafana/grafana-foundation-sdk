// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Logs] = (*LogsBuilder)(nil)

type LogsBuilder struct {
	internal *Logs
	errors   cog.BuildErrors
}

func NewLogsBuilder() *LogsBuilder {
	resource := NewLogs()
	builder := &LogsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *LogsBuilder) Build() (Logs, error) {
	if err := builder.internal.Validate(); err != nil {
		return Logs{}, err
	}

	if len(builder.errors) > 0 {
		return Logs{}, cog.MakeBuildErrors("elasticsearch.logs", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *LogsBuilder) Id(id string) *LogsBuilder {
	builder.internal.Id = id

	return builder
}

func (builder *LogsBuilder) Settings(settings cog.Builder[ElasticsearchLogsSettings]) *LogsBuilder {
	settingsResource, err := settings.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Settings = &settingsResource

	return builder
}

func (builder *LogsBuilder) Hide(hide bool) *LogsBuilder {
	builder.internal.Hide = &hide

	return builder
}
