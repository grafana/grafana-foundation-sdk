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
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("Logs", err)...)
	}

	if len(errs) != 0 {
		return Logs{}, errs
	}

	return *builder.internal, nil
}

func (builder *LogsBuilder) Id(id string) *LogsBuilder {
	builder.internal.Id = id

	return builder
}

func (builder *LogsBuilder) Settings(settings struct {
	Limit *string `json:"limit,omitempty"`
}) *LogsBuilder {
	builder.internal.Settings = &settings

	return builder
}

func (builder *LogsBuilder) Hide(hide bool) *LogsBuilder {
	builder.internal.Hide = &hide

	return builder
}

func (builder *LogsBuilder) applyDefaults() {
}
