// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package prometheus

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ScopesFilters] = (*ScopesFiltersBuilder)(nil)

type ScopesFiltersBuilder struct {
	internal *ScopesFilters
	errors   cog.BuildErrors
}

func NewScopesFiltersBuilder() *ScopesFiltersBuilder {
	resource := NewScopesFilters()
	builder := &ScopesFiltersBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ScopesFiltersBuilder) Build() (ScopesFilters, error) {
	if err := builder.internal.Validate(); err != nil {
		return ScopesFilters{}, err
	}

	if len(builder.errors) > 0 {
		return ScopesFilters{}, cog.MakeBuildErrors("prometheus.scopesFilters", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ScopesFiltersBuilder) RecordError(path string, err error) *ScopesFiltersBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *ScopesFiltersBuilder) Key(key string) *ScopesFiltersBuilder {
	builder.internal.Key = key

	return builder
}

func (builder *ScopesFiltersBuilder) Operator(operator string) *ScopesFiltersBuilder {
	builder.internal.Operator = operator

	return builder
}

func (builder *ScopesFiltersBuilder) Value(value string) *ScopesFiltersBuilder {
	builder.internal.Value = value

	return builder
}

func (builder *ScopesFiltersBuilder) Values(values []string) *ScopesFiltersBuilder {
	builder.internal.Values = values

	return builder
}
