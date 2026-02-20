// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboard

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[DashboardActionStyle] = (*DashboardActionStyleBuilder)(nil)

type DashboardActionStyleBuilder struct {
	internal *DashboardActionStyle
	errors   cog.BuildErrors
}

func NewDashboardActionStyleBuilder() *DashboardActionStyleBuilder {
	resource := NewDashboardActionStyle()
	builder := &DashboardActionStyleBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *DashboardActionStyleBuilder) Build() (DashboardActionStyle, error) {
	if err := builder.internal.Validate(); err != nil {
		return DashboardActionStyle{}, err
	}

	if len(builder.errors) > 0 {
		return DashboardActionStyle{}, cog.MakeBuildErrors("dashboard.dashboardActionStyle", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *DashboardActionStyleBuilder) RecordError(path string, err error) *DashboardActionStyleBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *DashboardActionStyleBuilder) BackgroundColor(backgroundColor string) *DashboardActionStyleBuilder {
	builder.internal.BackgroundColor = &backgroundColor

	return builder
}
