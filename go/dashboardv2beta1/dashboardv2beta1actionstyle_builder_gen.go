// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Dashboardv2beta1ActionStyle] = (*Dashboardv2beta1ActionStyleBuilder)(nil)

type Dashboardv2beta1ActionStyleBuilder struct {
	internal *Dashboardv2beta1ActionStyle
	errors   cog.BuildErrors
}

func NewDashboardv2beta1ActionStyleBuilder() *Dashboardv2beta1ActionStyleBuilder {
	resource := NewDashboardv2beta1ActionStyle()
	builder := &Dashboardv2beta1ActionStyleBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *Dashboardv2beta1ActionStyleBuilder) Build() (Dashboardv2beta1ActionStyle, error) {
	if err := builder.internal.Validate(); err != nil {
		return Dashboardv2beta1ActionStyle{}, err
	}

	if len(builder.errors) > 0 {
		return Dashboardv2beta1ActionStyle{}, cog.MakeBuildErrors("dashboardv2beta1.dashboardv2beta1ActionStyle", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *Dashboardv2beta1ActionStyleBuilder) RecordError(path string, err error) *Dashboardv2beta1ActionStyleBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *Dashboardv2beta1ActionStyleBuilder) BackgroundColor(backgroundColor string) *Dashboardv2beta1ActionStyleBuilder {
	builder.internal.BackgroundColor = &backgroundColor

	return builder
}
