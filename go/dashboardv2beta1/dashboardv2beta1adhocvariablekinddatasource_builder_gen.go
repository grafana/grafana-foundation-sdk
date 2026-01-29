// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Dashboardv2beta1AdhocVariableKindDatasource] = (*Dashboardv2beta1AdhocVariableKindDatasourceBuilder)(nil)

type Dashboardv2beta1AdhocVariableKindDatasourceBuilder struct {
	internal *Dashboardv2beta1AdhocVariableKindDatasource
	errors   cog.BuildErrors
}

func NewDashboardv2beta1AdhocVariableKindDatasourceBuilder() *Dashboardv2beta1AdhocVariableKindDatasourceBuilder {
	resource := NewDashboardv2beta1AdhocVariableKindDatasource()
	builder := &Dashboardv2beta1AdhocVariableKindDatasourceBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *Dashboardv2beta1AdhocVariableKindDatasourceBuilder) Build() (Dashboardv2beta1AdhocVariableKindDatasource, error) {
	if err := builder.internal.Validate(); err != nil {
		return Dashboardv2beta1AdhocVariableKindDatasource{}, err
	}

	if len(builder.errors) > 0 {
		return Dashboardv2beta1AdhocVariableKindDatasource{}, cog.MakeBuildErrors("dashboardv2beta1.dashboardv2beta1AdhocVariableKindDatasource", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *Dashboardv2beta1AdhocVariableKindDatasourceBuilder) Name(name string) *Dashboardv2beta1AdhocVariableKindDatasourceBuilder {
	builder.internal.Name = &name

	return builder
}
