// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[AzureTracesFilter] = (*AzureTracesFilterBuilder)(nil)

type AzureTracesFilterBuilder struct {
	internal *AzureTracesFilter
	errors   map[string]cog.BuildErrors
}

func NewAzureTracesFilterBuilder() *AzureTracesFilterBuilder {
	resource := &AzureTracesFilter{}
	builder := &AzureTracesFilterBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *AzureTracesFilterBuilder) Build() (AzureTracesFilter, error) {
	if err := builder.internal.Validate(); err != nil {
		return AzureTracesFilter{}, err
	}

	return *builder.internal, nil
}

// Property name, auto-populated based on available traces.
func (builder *AzureTracesFilterBuilder) Property(property string) *AzureTracesFilterBuilder {
	builder.internal.Property = property

	return builder
}

// Comparison operator to use. Either equals or not equals.
func (builder *AzureTracesFilterBuilder) Operation(operation string) *AzureTracesFilterBuilder {
	builder.internal.Operation = operation

	return builder
}

// Values to filter by.
func (builder *AzureTracesFilterBuilder) Filters(filters []string) *AzureTracesFilterBuilder {
	builder.internal.Filters = filters

	return builder
}

func (builder *AzureTracesFilterBuilder) applyDefaults() {
}
