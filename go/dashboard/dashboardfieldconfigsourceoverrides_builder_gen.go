// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboard

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[DashboardFieldConfigSourceOverrides] = (*DashboardFieldConfigSourceOverridesBuilder)(nil)

type DashboardFieldConfigSourceOverridesBuilder struct {
	internal *DashboardFieldConfigSourceOverrides
	errors   cog.BuildErrors
}

func NewDashboardFieldConfigSourceOverridesBuilder() *DashboardFieldConfigSourceOverridesBuilder {
	resource := NewDashboardFieldConfigSourceOverrides()
	builder := &DashboardFieldConfigSourceOverridesBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *DashboardFieldConfigSourceOverridesBuilder) Build() (DashboardFieldConfigSourceOverrides, error) {
	if err := builder.internal.Validate(); err != nil {
		return DashboardFieldConfigSourceOverrides{}, err
	}

	if len(builder.errors) > 0 {
		return DashboardFieldConfigSourceOverrides{}, cog.MakeBuildErrors("dashboard.dashboardFieldConfigSourceOverrides", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *DashboardFieldConfigSourceOverridesBuilder) Matcher(matcher MatcherConfig) *DashboardFieldConfigSourceOverridesBuilder {
	builder.internal.Matcher = matcher

	return builder
}

func (builder *DashboardFieldConfigSourceOverridesBuilder) Properties(properties []DynamicConfigValue) *DashboardFieldConfigSourceOverridesBuilder {
	builder.internal.Properties = properties

	return builder
}
