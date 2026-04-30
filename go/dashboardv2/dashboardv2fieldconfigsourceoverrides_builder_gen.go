// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Dashboardv2FieldConfigSourceOverrides] = (*Dashboardv2FieldConfigSourceOverridesBuilder)(nil)

type Dashboardv2FieldConfigSourceOverridesBuilder struct {
	internal *Dashboardv2FieldConfigSourceOverrides
	errors   cog.BuildErrors
}

func NewDashboardv2FieldConfigSourceOverridesBuilder() *Dashboardv2FieldConfigSourceOverridesBuilder {
	resource := NewDashboardv2FieldConfigSourceOverrides()
	builder := &Dashboardv2FieldConfigSourceOverridesBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *Dashboardv2FieldConfigSourceOverridesBuilder) Build() (Dashboardv2FieldConfigSourceOverrides, error) {
	if err := builder.internal.Validate(); err != nil {
		return Dashboardv2FieldConfigSourceOverrides{}, err
	}

	if len(builder.errors) > 0 {
		return Dashboardv2FieldConfigSourceOverrides{}, cog.MakeBuildErrors("dashboardv2.dashboardv2FieldConfigSourceOverrides", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *Dashboardv2FieldConfigSourceOverridesBuilder) RecordError(path string, err error) *Dashboardv2FieldConfigSourceOverridesBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *Dashboardv2FieldConfigSourceOverridesBuilder) SystemRef(systemRef string) *Dashboardv2FieldConfigSourceOverridesBuilder {
	builder.internal.SystemRef = &systemRef

	return builder
}

func (builder *Dashboardv2FieldConfigSourceOverridesBuilder) Matcher(matcher MatcherConfig) *Dashboardv2FieldConfigSourceOverridesBuilder {
	builder.internal.Matcher = matcher

	return builder
}

func (builder *Dashboardv2FieldConfigSourceOverridesBuilder) Properties(properties []DynamicConfigValue) *Dashboardv2FieldConfigSourceOverridesBuilder {
	builder.internal.Properties = properties

	return builder
}
