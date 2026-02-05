// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[FiltersSettings] = (*FiltersSettingsBuilder)(nil)

type FiltersSettingsBuilder struct {
	internal *FiltersSettings
	errors   cog.BuildErrors
}

func NewFiltersSettingsBuilder() *FiltersSettingsBuilder {
	resource := NewFiltersSettings()
	builder := &FiltersSettingsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *FiltersSettingsBuilder) Build() (FiltersSettings, error) {
	if err := builder.internal.Validate(); err != nil {
		return FiltersSettings{}, err
	}

	if len(builder.errors) > 0 {
		return FiltersSettings{}, cog.MakeBuildErrors("elasticsearch.filtersSettings", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *FiltersSettingsBuilder) Filters(filters []cog.Builder[Filter]) *FiltersSettingsBuilder {
	filtersResources := make([]Filter, 0, len(filters))
	for _, r1 := range filters {
		filtersDepth1, err := r1.Build()
		if err != nil {
			builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
			return builder
		}
		filtersResources = append(filtersResources, filtersDepth1)
	}
	builder.internal.Filters = filtersResources

	return builder
}
