// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboard

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[DashboardFieldConfigSourceOverrides] = (*DashboardFieldConfigSourceOverridesBuilder)(nil)

type DashboardFieldConfigSourceOverridesBuilder struct {
	internal *DashboardFieldConfigSourceOverrides
	errors   map[string]cog.BuildErrors
}

func NewDashboardFieldConfigSourceOverridesBuilder() *DashboardFieldConfigSourceOverridesBuilder {
	resource := &DashboardFieldConfigSourceOverrides{}
	builder := &DashboardFieldConfigSourceOverridesBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *DashboardFieldConfigSourceOverridesBuilder) Build() (DashboardFieldConfigSourceOverrides, error) {
	if err := builder.internal.Validate(); err != nil {
		return DashboardFieldConfigSourceOverrides{}, err
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

func (builder *DashboardFieldConfigSourceOverridesBuilder) applyDefaults() {
}
