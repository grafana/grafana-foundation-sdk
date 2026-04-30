// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Dashboardv2ActionStyle] = (*Dashboardv2ActionStyleBuilder)(nil)

type Dashboardv2ActionStyleBuilder struct {
	internal *Dashboardv2ActionStyle
	errors   cog.BuildErrors
}

func NewDashboardv2ActionStyleBuilder() *Dashboardv2ActionStyleBuilder {
	resource := NewDashboardv2ActionStyle()
	builder := &Dashboardv2ActionStyleBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *Dashboardv2ActionStyleBuilder) Build() (Dashboardv2ActionStyle, error) {
	if err := builder.internal.Validate(); err != nil {
		return Dashboardv2ActionStyle{}, err
	}

	if len(builder.errors) > 0 {
		return Dashboardv2ActionStyle{}, cog.MakeBuildErrors("dashboardv2.dashboardv2ActionStyle", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *Dashboardv2ActionStyleBuilder) RecordError(path string, err error) *Dashboardv2ActionStyleBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *Dashboardv2ActionStyleBuilder) BackgroundColor(backgroundColor string) *Dashboardv2ActionStyleBuilder {
	builder.internal.BackgroundColor = &backgroundColor

	return builder
}
