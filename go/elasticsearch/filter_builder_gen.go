// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Filter] = (*FilterBuilder)(nil)

type FilterBuilder struct {
	internal *Filter
	errors   map[string]cog.BuildErrors
}

func NewFilterBuilder() *FilterBuilder {
	resource := &Filter{}
	builder := &FilterBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *FilterBuilder) Build() (Filter, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("Filter", err)...)
	}

	if len(errs) != 0 {
		return Filter{}, errs
	}

	return *builder.internal, nil
}

func (builder *FilterBuilder) Query(query string) *FilterBuilder {
	builder.internal.Query = query

	return builder
}

func (builder *FilterBuilder) Label(label string) *FilterBuilder {
	builder.internal.Label = label

	return builder
}

func (builder *FilterBuilder) applyDefaults() {
}
