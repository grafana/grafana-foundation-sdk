// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[AzureTracesFilter] = (*AzureTracesFilterBuilder)(nil)

type AzureTracesFilterBuilder struct {
	internal *AzureTracesFilter
	errors   cog.BuildErrors
}

func NewAzureTracesFilterBuilder() *AzureTracesFilterBuilder {
	resource := NewAzureTracesFilter()
	builder := &AzureTracesFilterBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *AzureTracesFilterBuilder) Build() (AzureTracesFilter, error) {
	if err := builder.internal.Validate(); err != nil {
		return AzureTracesFilter{}, err
	}

	if len(builder.errors) > 0 {
		return AzureTracesFilter{}, cog.MakeBuildErrors("azuremonitor.azureTracesFilter", builder.errors)
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
