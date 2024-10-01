// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ElasticsearchLogsSettings] = (*ElasticsearchLogsSettingsBuilder)(nil)

type ElasticsearchLogsSettingsBuilder struct {
	internal *ElasticsearchLogsSettings
	errors   map[string]cog.BuildErrors
}

func NewElasticsearchLogsSettingsBuilder() *ElasticsearchLogsSettingsBuilder {
	resource := &ElasticsearchLogsSettings{}
	builder := &ElasticsearchLogsSettingsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *ElasticsearchLogsSettingsBuilder) Build() (ElasticsearchLogsSettings, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("ElasticsearchLogsSettings", err)...)
	}

	if len(errs) != 0 {
		return ElasticsearchLogsSettings{}, errs
	}

	return *builder.internal, nil
}

func (builder *ElasticsearchLogsSettingsBuilder) Limit(limit string) *ElasticsearchLogsSettingsBuilder {
	builder.internal.Limit = &limit

	return builder
}

func (builder *ElasticsearchLogsSettingsBuilder) applyDefaults() {
}
