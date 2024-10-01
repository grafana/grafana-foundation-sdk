// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ElasticsearchDerivativeSettings] = (*ElasticsearchDerivativeSettingsBuilder)(nil)

type ElasticsearchDerivativeSettingsBuilder struct {
	internal *ElasticsearchDerivativeSettings
	errors   map[string]cog.BuildErrors
}

func NewElasticsearchDerivativeSettingsBuilder() *ElasticsearchDerivativeSettingsBuilder {
	resource := &ElasticsearchDerivativeSettings{}
	builder := &ElasticsearchDerivativeSettingsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *ElasticsearchDerivativeSettingsBuilder) Build() (ElasticsearchDerivativeSettings, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("ElasticsearchDerivativeSettings", err)...)
	}

	if len(errs) != 0 {
		return ElasticsearchDerivativeSettings{}, errs
	}

	return *builder.internal, nil
}

func (builder *ElasticsearchDerivativeSettingsBuilder) Unit(unit string) *ElasticsearchDerivativeSettingsBuilder {
	builder.internal.Unit = &unit

	return builder
}

func (builder *ElasticsearchDerivativeSettingsBuilder) applyDefaults() {
}
