// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package googlecloudmonitoring

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Filter] = (*FilterBuilder)(nil)

// Query filter representation.
type FilterBuilder struct {
	internal *Filter
	errors   cog.BuildErrors
}

func NewFilterBuilder() *FilterBuilder {
	resource := NewFilter()
	builder := &FilterBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *FilterBuilder) Build() (Filter, error) {
	if err := builder.internal.Validate(); err != nil {
		return Filter{}, err
	}

	if len(builder.errors) > 0 {
		return Filter{}, cog.MakeBuildErrors("googlecloudmonitoring.filter", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *FilterBuilder) RecordError(path string, err error) *FilterBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

// Filter key.
func (builder *FilterBuilder) Key(key string) *FilterBuilder {
	builder.internal.Key = key

	return builder
}

// Filter operator.
func (builder *FilterBuilder) Operator(operator string) *FilterBuilder {
	builder.internal.Operator = operator

	return builder
}

// Filter value.
func (builder *FilterBuilder) Value(value string) *FilterBuilder {
	builder.internal.Value = value

	return builder
}

// Filter condition.
func (builder *FilterBuilder) Condition(condition string) *FilterBuilder {
	builder.internal.Condition = &condition

	return builder
}
