// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[TracesFilter] = (*TracesFilterBuilder)(nil)

type TracesFilterBuilder struct {
	internal *TracesFilter
	errors   cog.BuildErrors
}

func NewTracesFilterBuilder() *TracesFilterBuilder {
	resource := NewTracesFilter()
	builder := &TracesFilterBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *TracesFilterBuilder) Build() (TracesFilter, error) {
	if err := builder.internal.Validate(); err != nil {
		return TracesFilter{}, err
	}

	if len(builder.errors) > 0 {
		return TracesFilter{}, cog.MakeBuildErrors("azuremonitor.tracesFilter", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *TracesFilterBuilder) RecordError(path string, err error) *TracesFilterBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

// Property name, auto-populated based on available traces.
func (builder *TracesFilterBuilder) Property(property string) *TracesFilterBuilder {
	builder.internal.Property = property

	return builder
}

// Comparison operator to use. Either equals or not equals.
func (builder *TracesFilterBuilder) Operation(operation string) *TracesFilterBuilder {
	builder.internal.Operation = operation

	return builder
}

// Values to filter by.
func (builder *TracesFilterBuilder) Filters(filters []string) *TracesFilterBuilder {
	builder.internal.Filters = filters

	return builder
}
