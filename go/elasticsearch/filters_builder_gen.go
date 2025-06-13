// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Filters] = (*FiltersBuilder)(nil)

type FiltersBuilder struct {
	internal *Filters
	errors   cog.BuildErrors
}

func NewFiltersBuilder() *FiltersBuilder {
	resource := NewFilters()
	builder := &FiltersBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *FiltersBuilder) Build() (Filters, error) {
	if err := builder.internal.Validate(); err != nil {
		return Filters{}, err
	}

	if len(builder.errors) > 0 {
		return Filters{}, cog.MakeBuildErrors("elasticsearch.filters", builder.errors)
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
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Settings = &settingsResource

	return builder
}
