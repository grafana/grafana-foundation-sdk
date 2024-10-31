// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package googlecloudmonitoring

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[PromQLQuery] = (*PromQLQueryBuilder)(nil)

// PromQL sub-query properties.
type PromQLQueryBuilder struct {
	internal *PromQLQuery
	errors   map[string]cog.BuildErrors
}

func NewPromQLQueryBuilder() *PromQLQueryBuilder {
	resource := &PromQLQuery{}
	builder := &PromQLQueryBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *PromQLQueryBuilder) Build() (PromQLQuery, error) {
	if err := builder.internal.Validate(); err != nil {
		return PromQLQuery{}, err
	}

	return *builder.internal, nil
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

func (builder *PromQLQueryBuilder) applyDefaults() {
}
