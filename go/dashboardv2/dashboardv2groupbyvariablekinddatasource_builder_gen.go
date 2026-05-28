// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Dashboardv2GroupByVariableKindDatasource] = (*Dashboardv2GroupByVariableKindDatasourceBuilder)(nil)

type Dashboardv2GroupByVariableKindDatasourceBuilder struct {
	internal *Dashboardv2GroupByVariableKindDatasource
	errors   cog.BuildErrors
}

func NewDashboardv2GroupByVariableKindDatasourceBuilder() *Dashboardv2GroupByVariableKindDatasourceBuilder {
	resource := NewDashboardv2GroupByVariableKindDatasource()
	builder := &Dashboardv2GroupByVariableKindDatasourceBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *Dashboardv2GroupByVariableKindDatasourceBuilder) Build() (Dashboardv2GroupByVariableKindDatasource, error) {
	if err := builder.internal.Validate(); err != nil {
		return Dashboardv2GroupByVariableKindDatasource{}, err
	}

	if len(builder.errors) > 0 {
		return Dashboardv2GroupByVariableKindDatasource{}, cog.MakeBuildErrors("dashboardv2.dashboardv2GroupByVariableKindDatasource", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *Dashboardv2GroupByVariableKindDatasourceBuilder) RecordError(path string, err error) *Dashboardv2GroupByVariableKindDatasourceBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *Dashboardv2GroupByVariableKindDatasourceBuilder) Name(name string) *Dashboardv2GroupByVariableKindDatasourceBuilder {
	builder.internal.Name = &name

	return builder
}
