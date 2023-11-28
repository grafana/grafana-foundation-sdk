// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[FiltersSettings] = (*FiltersSettingsBuilder)(nil)

type FiltersSettingsBuilder struct {
	internal *FiltersSettings
	errors   map[string]cog.BuildErrors
}

func NewFiltersSettingsBuilder() *FiltersSettingsBuilder {
	resource := &FiltersSettings{}
	builder := &FiltersSettingsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *FiltersSettingsBuilder) Build() (FiltersSettings, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("FiltersSettings", err)...)
	}

	if len(errs) != 0 {
		return FiltersSettings{}, errs
	}

	return *builder.internal, nil
}

func (builder *FiltersSettingsBuilder) Filters(filters []cog.Builder[Filter]) *FiltersSettingsBuilder {
	filtersResources := make([]Filter, 0, len(filters))
	for _, r := range filters {
		filtersResource, err := r.Build()
		if err != nil {
			builder.errors["filters"] = err.(cog.BuildErrors)
			return builder
		}
		filtersResources = append(filtersResources, filtersResource)
	}
	builder.internal.Filters = filtersResources

	return builder
}

func (builder *FiltersSettingsBuilder) applyDefaults() {
}
