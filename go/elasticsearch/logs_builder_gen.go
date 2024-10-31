// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Logs] = (*LogsBuilder)(nil)

type LogsBuilder struct {
	internal *Logs
	errors   map[string]cog.BuildErrors
}

func NewLogsBuilder() *LogsBuilder {
	resource := &Logs{}
	builder := &LogsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Type = "logs"

	return builder
}

func (builder *LogsBuilder) Build() (Logs, error) {
	if err := builder.internal.Validate(); err != nil {
		return Logs{}, err
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
		builder.errors["settings"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Settings = &settingsResource

	return builder
}

func (builder *LogsBuilder) Hide(hide bool) *LogsBuilder {
	builder.internal.Hide = &hide

	return builder
}

func (builder *LogsBuilder) applyDefaults() {
}
