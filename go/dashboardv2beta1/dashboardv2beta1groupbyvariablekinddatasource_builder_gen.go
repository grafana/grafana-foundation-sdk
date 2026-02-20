// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Dashboardv2beta1GroupByVariableKindDatasource] = (*Dashboardv2beta1GroupByVariableKindDatasourceBuilder)(nil)

type Dashboardv2beta1GroupByVariableKindDatasourceBuilder struct {
	internal *Dashboardv2beta1GroupByVariableKindDatasource
	errors   cog.BuildErrors
}

func NewDashboardv2beta1GroupByVariableKindDatasourceBuilder() *Dashboardv2beta1GroupByVariableKindDatasourceBuilder {
	resource := NewDashboardv2beta1GroupByVariableKindDatasource()
	builder := &Dashboardv2beta1GroupByVariableKindDatasourceBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *Dashboardv2beta1GroupByVariableKindDatasourceBuilder) Build() (Dashboardv2beta1GroupByVariableKindDatasource, error) {
	if err := builder.internal.Validate(); err != nil {
		return Dashboardv2beta1GroupByVariableKindDatasource{}, err
	}

	if len(builder.errors) > 0 {
		return Dashboardv2beta1GroupByVariableKindDatasource{}, cog.MakeBuildErrors("dashboardv2beta1.dashboardv2beta1GroupByVariableKindDatasource", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *Dashboardv2beta1GroupByVariableKindDatasourceBuilder) RecordError(path string, err error) *Dashboardv2beta1GroupByVariableKindDatasourceBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *Dashboardv2beta1GroupByVariableKindDatasourceBuilder) Name(name string) *Dashboardv2beta1GroupByVariableKindDatasourceBuilder {
	builder.internal.Name = &name

	return builder
}
