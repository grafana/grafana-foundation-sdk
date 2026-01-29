// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Dashboardv2beta1SpecialValueMapOptions] = (*Dashboardv2beta1SpecialValueMapOptionsBuilder)(nil)

type Dashboardv2beta1SpecialValueMapOptionsBuilder struct {
	internal *Dashboardv2beta1SpecialValueMapOptions
	errors   cog.BuildErrors
}

func NewDashboardv2beta1SpecialValueMapOptionsBuilder() *Dashboardv2beta1SpecialValueMapOptionsBuilder {
	resource := NewDashboardv2beta1SpecialValueMapOptions()
	builder := &Dashboardv2beta1SpecialValueMapOptionsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *Dashboardv2beta1SpecialValueMapOptionsBuilder) Build() (Dashboardv2beta1SpecialValueMapOptions, error) {
	if err := builder.internal.Validate(); err != nil {
		return Dashboardv2beta1SpecialValueMapOptions{}, err
	}

	if len(builder.errors) > 0 {
		return Dashboardv2beta1SpecialValueMapOptions{}, cog.MakeBuildErrors("dashboardv2beta1.dashboardv2beta1SpecialValueMapOptions", builder.errors)
	}

	return *builder.internal, nil
}

// Special value to match against
func (builder *Dashboardv2beta1SpecialValueMapOptionsBuilder) Match(match SpecialValueMatch) *Dashboardv2beta1SpecialValueMapOptionsBuilder {
	builder.internal.Match = match

	return builder
}

// Config to apply when the value matches the special value
func (builder *Dashboardv2beta1SpecialValueMapOptionsBuilder) Result(result cog.Builder[ValueMappingResult]) *Dashboardv2beta1SpecialValueMapOptionsBuilder {
	resultResource, err := result.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Result = resultResource

	return builder
}
