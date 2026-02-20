// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Dashboardv2beta1DataQueryKindDatasource] = (*Dashboardv2beta1DataQueryKindDatasourceBuilder)(nil)

type Dashboardv2beta1DataQueryKindDatasourceBuilder struct {
	internal *Dashboardv2beta1DataQueryKindDatasource
	errors   cog.BuildErrors
}

func NewDashboardv2beta1DataQueryKindDatasourceBuilder() *Dashboardv2beta1DataQueryKindDatasourceBuilder {
	resource := NewDashboardv2beta1DataQueryKindDatasource()
	builder := &Dashboardv2beta1DataQueryKindDatasourceBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *Dashboardv2beta1DataQueryKindDatasourceBuilder) Build() (Dashboardv2beta1DataQueryKindDatasource, error) {
	if err := builder.internal.Validate(); err != nil {
		return Dashboardv2beta1DataQueryKindDatasource{}, err
	}

	if len(builder.errors) > 0 {
		return Dashboardv2beta1DataQueryKindDatasource{}, cog.MakeBuildErrors("dashboardv2beta1.dashboardv2beta1DataQueryKindDatasource", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *Dashboardv2beta1DataQueryKindDatasourceBuilder) RecordError(path string, err error) *Dashboardv2beta1DataQueryKindDatasourceBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *Dashboardv2beta1DataQueryKindDatasourceBuilder) Name(name string) *Dashboardv2beta1DataQueryKindDatasourceBuilder {
	builder.internal.Name = &name

	return builder
}
