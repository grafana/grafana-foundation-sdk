// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[AzureResourceGraphQuery] = (*AzureResourceGraphQueryBuilder)(nil)

type AzureResourceGraphQueryBuilder struct {
	internal *AzureResourceGraphQuery
	errors   map[string]cog.BuildErrors
}

func NewAzureResourceGraphQueryBuilder() *AzureResourceGraphQueryBuilder {
	resource := &AzureResourceGraphQuery{}
	builder := &AzureResourceGraphQueryBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *AzureResourceGraphQueryBuilder) Build() (AzureResourceGraphQuery, error) {
	if err := builder.internal.Validate(); err != nil {
		return AzureResourceGraphQuery{}, err
	}

	return *builder.internal, nil
}

// Azure Resource Graph KQL query to be executed.
func (builder *AzureResourceGraphQueryBuilder) Query(query string) *AzureResourceGraphQueryBuilder {
	builder.internal.Query = &query

	return builder
}

// Specifies the format results should be returned as. Defaults to table.
func (builder *AzureResourceGraphQueryBuilder) ResultFormat(resultFormat string) *AzureResourceGraphQueryBuilder {
	builder.internal.ResultFormat = &resultFormat

	return builder
}

func (builder *AzureResourceGraphQueryBuilder) applyDefaults() {
}
