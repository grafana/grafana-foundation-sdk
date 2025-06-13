// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboard

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[DashboardDashboardTemplating] = (*DashboardDashboardTemplatingBuilder)(nil)

type DashboardDashboardTemplatingBuilder struct {
	internal *DashboardDashboardTemplating
	errors   cog.BuildErrors
}

func NewDashboardDashboardTemplatingBuilder() *DashboardDashboardTemplatingBuilder {
	resource := NewDashboardDashboardTemplating()
	builder := &DashboardDashboardTemplatingBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *DashboardDashboardTemplatingBuilder) Build() (DashboardDashboardTemplating, error) {
	if err := builder.internal.Validate(); err != nil {
		return DashboardDashboardTemplating{}, err
	}

	if len(builder.errors) > 0 {
		return DashboardDashboardTemplating{}, cog.MakeBuildErrors("dashboard.dashboardDashboardTemplating", builder.errors)
	}

	return *builder.internal, nil
}

// List of configured template variables with their saved values along with some other metadata
func (builder *DashboardDashboardTemplatingBuilder) List(list []cog.Builder[VariableModel]) *DashboardDashboardTemplatingBuilder {
	listResources := make([]VariableModel, 0, len(list))
	for _, r1 := range list {
		listDepth1, err := r1.Build()
		if err != nil {
			builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
			return builder
		}
		listResources = append(listResources, listDepth1)
	}
	builder.internal.List = listResources

	return builder
}
