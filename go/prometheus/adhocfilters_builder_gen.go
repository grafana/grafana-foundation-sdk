// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package prometheus

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[AdhocFilters] = (*AdhocFiltersBuilder)(nil)

type AdhocFiltersBuilder struct {
	internal *AdhocFilters
	errors   cog.BuildErrors
}

func NewAdhocFiltersBuilder() *AdhocFiltersBuilder {
	resource := NewAdhocFilters()
	builder := &AdhocFiltersBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *AdhocFiltersBuilder) Build() (AdhocFilters, error) {
	if err := builder.internal.Validate(); err != nil {
		return AdhocFilters{}, err
	}

	if len(builder.errors) > 0 {
		return AdhocFilters{}, cog.MakeBuildErrors("prometheus.adhocFilters", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *AdhocFiltersBuilder) RecordError(path string, err error) *AdhocFiltersBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *AdhocFiltersBuilder) Key(key string) *AdhocFiltersBuilder {
	builder.internal.Key = key

	return builder
}

func (builder *AdhocFiltersBuilder) Operator(operator string) *AdhocFiltersBuilder {
	builder.internal.Operator = operator

	return builder
}

func (builder *AdhocFiltersBuilder) Value(value string) *AdhocFiltersBuilder {
	builder.internal.Value = value

	return builder
}

func (builder *AdhocFiltersBuilder) Values(values []string) *AdhocFiltersBuilder {
	builder.internal.Values = values

	return builder
}
