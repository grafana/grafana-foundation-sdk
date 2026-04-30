// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Dashboardv2DataQueryKindDatasource] = (*Dashboardv2DataQueryKindDatasourceBuilder)(nil)

type Dashboardv2DataQueryKindDatasourceBuilder struct {
	internal *Dashboardv2DataQueryKindDatasource
	errors   cog.BuildErrors
}

func NewDashboardv2DataQueryKindDatasourceBuilder() *Dashboardv2DataQueryKindDatasourceBuilder {
	resource := NewDashboardv2DataQueryKindDatasource()
	builder := &Dashboardv2DataQueryKindDatasourceBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *Dashboardv2DataQueryKindDatasourceBuilder) Build() (Dashboardv2DataQueryKindDatasource, error) {
	if err := builder.internal.Validate(); err != nil {
		return Dashboardv2DataQueryKindDatasource{}, err
	}

	if len(builder.errors) > 0 {
		return Dashboardv2DataQueryKindDatasource{}, cog.MakeBuildErrors("dashboardv2.dashboardv2DataQueryKindDatasource", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *Dashboardv2DataQueryKindDatasourceBuilder) RecordError(path string, err error) *Dashboardv2DataQueryKindDatasourceBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *Dashboardv2DataQueryKindDatasourceBuilder) Name(name string) *Dashboardv2DataQueryKindDatasourceBuilder {
	builder.internal.Name = &name

	return builder
}
