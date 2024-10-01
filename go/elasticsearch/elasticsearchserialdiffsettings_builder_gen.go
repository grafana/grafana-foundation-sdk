// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ElasticsearchSerialDiffSettings] = (*ElasticsearchSerialDiffSettingsBuilder)(nil)

type ElasticsearchSerialDiffSettingsBuilder struct {
	internal *ElasticsearchSerialDiffSettings
	errors   map[string]cog.BuildErrors
}

func NewElasticsearchSerialDiffSettingsBuilder() *ElasticsearchSerialDiffSettingsBuilder {
	resource := &ElasticsearchSerialDiffSettings{}
	builder := &ElasticsearchSerialDiffSettingsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *ElasticsearchSerialDiffSettingsBuilder) Build() (ElasticsearchSerialDiffSettings, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("ElasticsearchSerialDiffSettings", err)...)
	}

	if len(errs) != 0 {
		return ElasticsearchSerialDiffSettings{}, errs
	}

	return *builder.internal, nil
}

func (builder *ElasticsearchSerialDiffSettingsBuilder) Lag(lag string) *ElasticsearchSerialDiffSettingsBuilder {
	builder.internal.Lag = &lag

	return builder
}

func (builder *ElasticsearchSerialDiffSettingsBuilder) applyDefaults() {
}
