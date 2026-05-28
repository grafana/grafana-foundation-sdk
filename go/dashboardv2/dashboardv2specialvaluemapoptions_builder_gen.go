// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Dashboardv2SpecialValueMapOptions] = (*Dashboardv2SpecialValueMapOptionsBuilder)(nil)

type Dashboardv2SpecialValueMapOptionsBuilder struct {
	internal *Dashboardv2SpecialValueMapOptions
	errors   cog.BuildErrors
}

func NewDashboardv2SpecialValueMapOptionsBuilder() *Dashboardv2SpecialValueMapOptionsBuilder {
	resource := NewDashboardv2SpecialValueMapOptions()
	builder := &Dashboardv2SpecialValueMapOptionsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *Dashboardv2SpecialValueMapOptionsBuilder) Build() (Dashboardv2SpecialValueMapOptions, error) {
	if err := builder.internal.Validate(); err != nil {
		return Dashboardv2SpecialValueMapOptions{}, err
	}

	if len(builder.errors) > 0 {
		return Dashboardv2SpecialValueMapOptions{}, cog.MakeBuildErrors("dashboardv2.dashboardv2SpecialValueMapOptions", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *Dashboardv2SpecialValueMapOptionsBuilder) RecordError(path string, err error) *Dashboardv2SpecialValueMapOptionsBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

// Special value to match against
func (builder *Dashboardv2SpecialValueMapOptionsBuilder) Match(match SpecialValueMatch) *Dashboardv2SpecialValueMapOptionsBuilder {
	builder.internal.Match = match

	return builder
}

// Config to apply when the value matches the special value
func (builder *Dashboardv2SpecialValueMapOptionsBuilder) Result(result cog.Builder[ValueMappingResult]) *Dashboardv2SpecialValueMapOptionsBuilder {
	resultResource, err := result.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Result = resultResource

	return builder
}
