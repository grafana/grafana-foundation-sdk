// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package prometheus

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Scopes] = (*ScopesBuilder)(nil)

type ScopesBuilder struct {
	internal *Scopes
	errors   cog.BuildErrors
}

func NewScopesBuilder() *ScopesBuilder {
	resource := NewScopes()
	builder := &ScopesBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ScopesBuilder) Build() (Scopes, error) {
	if err := builder.internal.Validate(); err != nil {
		return Scopes{}, err
	}

	if len(builder.errors) > 0 {
		return Scopes{}, cog.MakeBuildErrors("prometheus.scopes", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ScopesBuilder) RecordError(path string, err error) *ScopesBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *ScopesBuilder) DefaultPath(defaultPath []string) *ScopesBuilder {
	builder.internal.DefaultPath = defaultPath

	return builder
}

func (builder *ScopesBuilder) Filters(filters []cog.Builder[ScopesFilters]) *ScopesBuilder {
	filtersResources := make([]ScopesFilters, 0, len(filters))
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

func (builder *ScopesBuilder) Title(title string) *ScopesBuilder {
	builder.internal.Title = title

	return builder
}
