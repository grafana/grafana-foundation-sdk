// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ResourceGraphQuery] = (*ResourceGraphQueryBuilder)(nil)

type ResourceGraphQueryBuilder struct {
	internal *ResourceGraphQuery
	errors   cog.BuildErrors
}

func NewResourceGraphQueryBuilder() *ResourceGraphQueryBuilder {
	resource := NewResourceGraphQuery()
	builder := &ResourceGraphQueryBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ResourceGraphQueryBuilder) Build() (ResourceGraphQuery, error) {
	if err := builder.internal.Validate(); err != nil {
		return ResourceGraphQuery{}, err
	}

	if len(builder.errors) > 0 {
		return ResourceGraphQuery{}, cog.MakeBuildErrors("azuremonitor.resourceGraphQuery", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ResourceGraphQueryBuilder) RecordError(path string, err error) *ResourceGraphQueryBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

// Azure Resource Graph KQL query to be executed.
func (builder *ResourceGraphQueryBuilder) Query(query string) *ResourceGraphQueryBuilder {
	builder.internal.Query = &query

	return builder
}

// Specifies the format results should be returned as. Defaults to table.
func (builder *ResourceGraphQueryBuilder) ResultFormat(resultFormat string) *ResourceGraphQueryBuilder {
	builder.internal.ResultFormat = &resultFormat

	return builder
}
