// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package googlecloudmonitoring

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Filter] = (*FilterBuilder)(nil)

// Query filter representation.
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
	if err := builder.internal.Validate(); err != nil {
		return Filter{}, err
	}

	return *builder.internal, nil
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

func (builder *FilterBuilder) applyDefaults() {
}
