// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package googlecloudmonitoring

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[PromQLQuery] = (*PromQLQueryBuilder)(nil)

// PromQL sub-query properties.
type PromQLQueryBuilder struct {
	internal *PromQLQuery
	errors   cog.BuildErrors
}

func NewPromQLQueryBuilder() *PromQLQueryBuilder {
	resource := NewPromQLQuery()
	builder := &PromQLQueryBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *PromQLQueryBuilder) Build() (PromQLQuery, error) {
	if err := builder.internal.Validate(); err != nil {
		return PromQLQuery{}, err
	}

	if len(builder.errors) > 0 {
		return PromQLQuery{}, cog.MakeBuildErrors("googlecloudmonitoring.promQLQuery", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *PromQLQueryBuilder) RecordError(path string, err error) *PromQLQueryBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

// GCP project to execute the query against.
func (builder *PromQLQueryBuilder) ProjectName(projectName string) *PromQLQueryBuilder {
	builder.internal.ProjectName = projectName

	return builder
}

// PromQL expression/query to be executed.
func (builder *PromQLQueryBuilder) Expr(expr string) *PromQLQueryBuilder {
	builder.internal.Expr = expr

	return builder
}

// PromQL min step
func (builder *PromQLQueryBuilder) Step(step string) *PromQLQueryBuilder {
	builder.internal.Step = step

	return builder
}
