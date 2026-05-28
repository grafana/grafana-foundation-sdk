// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Dashboardv2AdhocVariableKindDatasource] = (*Dashboardv2AdhocVariableKindDatasourceBuilder)(nil)

type Dashboardv2AdhocVariableKindDatasourceBuilder struct {
	internal *Dashboardv2AdhocVariableKindDatasource
	errors   cog.BuildErrors
}

func NewDashboardv2AdhocVariableKindDatasourceBuilder() *Dashboardv2AdhocVariableKindDatasourceBuilder {
	resource := NewDashboardv2AdhocVariableKindDatasource()
	builder := &Dashboardv2AdhocVariableKindDatasourceBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *Dashboardv2AdhocVariableKindDatasourceBuilder) Build() (Dashboardv2AdhocVariableKindDatasource, error) {
	if err := builder.internal.Validate(); err != nil {
		return Dashboardv2AdhocVariableKindDatasource{}, err
	}

	if len(builder.errors) > 0 {
		return Dashboardv2AdhocVariableKindDatasource{}, cog.MakeBuildErrors("dashboardv2.dashboardv2AdhocVariableKindDatasource", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *Dashboardv2AdhocVariableKindDatasourceBuilder) RecordError(path string, err error) *Dashboardv2AdhocVariableKindDatasourceBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *Dashboardv2AdhocVariableKindDatasourceBuilder) Name(name string) *Dashboardv2AdhocVariableKindDatasourceBuilder {
	builder.internal.Name = &name

	return builder
}
