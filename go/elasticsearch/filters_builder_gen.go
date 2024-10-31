// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Filters] = (*FiltersBuilder)(nil)

type FiltersBuilder struct {
	internal *Filters
	errors   map[string]cog.BuildErrors
}

func NewFiltersBuilder() *FiltersBuilder {
	resource := &Filters{}
	builder := &FiltersBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Type = "filters"

	return builder
}

func (builder *FiltersBuilder) Build() (Filters, error) {
	if err := builder.internal.Validate(); err != nil {
		return Filters{}, err
	}

	return *builder.internal, nil
}

func (builder *FiltersBuilder) Id(id string) *FiltersBuilder {
	builder.internal.Id = id

	return builder
}

func (builder *FiltersBuilder) Settings(settings cog.Builder[ElasticsearchFiltersSettings]) *FiltersBuilder {
	settingsResource, err := settings.Build()
	if err != nil {
		builder.errors["settings"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Settings = &settingsResource

	return builder
}

func (builder *FiltersBuilder) applyDefaults() {
}
