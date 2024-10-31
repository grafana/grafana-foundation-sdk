// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ElasticsearchFiltersSettings] = (*ElasticsearchFiltersSettingsBuilder)(nil)

type ElasticsearchFiltersSettingsBuilder struct {
	internal *ElasticsearchFiltersSettings
	errors   map[string]cog.BuildErrors
}

func NewElasticsearchFiltersSettingsBuilder() *ElasticsearchFiltersSettingsBuilder {
	resource := &ElasticsearchFiltersSettings{}
	builder := &ElasticsearchFiltersSettingsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *ElasticsearchFiltersSettingsBuilder) Build() (ElasticsearchFiltersSettings, error) {
	if err := builder.internal.Validate(); err != nil {
		return ElasticsearchFiltersSettings{}, err
	}

	return *builder.internal, nil
}

func (builder *ElasticsearchFiltersSettingsBuilder) Filters(filters []cog.Builder[Filter]) *ElasticsearchFiltersSettingsBuilder {
	filtersResources := make([]Filter, 0, len(filters))
	for _, r1 := range filters {
		filtersDepth1, err := r1.Build()
		if err != nil {
			builder.errors["filters"] = err.(cog.BuildErrors)
			return builder
		}
		filtersResources = append(filtersResources, filtersDepth1)
	}
	builder.internal.Filters = filtersResources

	return builder
}

func (builder *ElasticsearchFiltersSettingsBuilder) applyDefaults() {
}
