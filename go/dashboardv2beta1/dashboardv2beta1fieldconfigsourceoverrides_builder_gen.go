// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Dashboardv2beta1FieldConfigSourceOverrides] = (*Dashboardv2beta1FieldConfigSourceOverridesBuilder)(nil)

type Dashboardv2beta1FieldConfigSourceOverridesBuilder struct {
	internal *Dashboardv2beta1FieldConfigSourceOverrides
	errors   cog.BuildErrors
}

func NewDashboardv2beta1FieldConfigSourceOverridesBuilder() *Dashboardv2beta1FieldConfigSourceOverridesBuilder {
	resource := NewDashboardv2beta1FieldConfigSourceOverrides()
	builder := &Dashboardv2beta1FieldConfigSourceOverridesBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *Dashboardv2beta1FieldConfigSourceOverridesBuilder) Build() (Dashboardv2beta1FieldConfigSourceOverrides, error) {
	if err := builder.internal.Validate(); err != nil {
		return Dashboardv2beta1FieldConfigSourceOverrides{}, err
	}

	if len(builder.errors) > 0 {
		return Dashboardv2beta1FieldConfigSourceOverrides{}, cog.MakeBuildErrors("dashboardv2beta1.dashboardv2beta1FieldConfigSourceOverrides", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *Dashboardv2beta1FieldConfigSourceOverridesBuilder) RecordError(path string, err error) *Dashboardv2beta1FieldConfigSourceOverridesBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *Dashboardv2beta1FieldConfigSourceOverridesBuilder) SystemRef(systemRef string) *Dashboardv2beta1FieldConfigSourceOverridesBuilder {
	builder.internal.SystemRef = &systemRef

	return builder
}

func (builder *Dashboardv2beta1FieldConfigSourceOverridesBuilder) Matcher(matcher MatcherConfig) *Dashboardv2beta1FieldConfigSourceOverridesBuilder {
	builder.internal.Matcher = matcher

	return builder
}

func (builder *Dashboardv2beta1FieldConfigSourceOverridesBuilder) Properties(properties []DynamicConfigValue) *Dashboardv2beta1FieldConfigSourceOverridesBuilder {
	builder.internal.Properties = properties

	return builder
}
