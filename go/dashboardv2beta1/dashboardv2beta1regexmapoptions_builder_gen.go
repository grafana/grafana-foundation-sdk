// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Dashboardv2beta1RegexMapOptions] = (*Dashboardv2beta1RegexMapOptionsBuilder)(nil)

type Dashboardv2beta1RegexMapOptionsBuilder struct {
	internal *Dashboardv2beta1RegexMapOptions
	errors   cog.BuildErrors
}

func NewDashboardv2beta1RegexMapOptionsBuilder() *Dashboardv2beta1RegexMapOptionsBuilder {
	resource := NewDashboardv2beta1RegexMapOptions()
	builder := &Dashboardv2beta1RegexMapOptionsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *Dashboardv2beta1RegexMapOptionsBuilder) Build() (Dashboardv2beta1RegexMapOptions, error) {
	if err := builder.internal.Validate(); err != nil {
		return Dashboardv2beta1RegexMapOptions{}, err
	}

	if len(builder.errors) > 0 {
		return Dashboardv2beta1RegexMapOptions{}, cog.MakeBuildErrors("dashboardv2beta1.dashboardv2beta1RegexMapOptions", builder.errors)
	}

	return *builder.internal, nil
}

// Regular expression to match against
func (builder *Dashboardv2beta1RegexMapOptionsBuilder) Pattern(pattern string) *Dashboardv2beta1RegexMapOptionsBuilder {
	builder.internal.Pattern = pattern

	return builder
}

// Config to apply when the value matches the regex
func (builder *Dashboardv2beta1RegexMapOptionsBuilder) Result(result cog.Builder[ValueMappingResult]) *Dashboardv2beta1RegexMapOptionsBuilder {
	resultResource, err := result.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Result = resultResource

	return builder
}
