// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Dashboardv2RegexMapOptions] = (*Dashboardv2RegexMapOptionsBuilder)(nil)

type Dashboardv2RegexMapOptionsBuilder struct {
	internal *Dashboardv2RegexMapOptions
	errors   cog.BuildErrors
}

func NewDashboardv2RegexMapOptionsBuilder() *Dashboardv2RegexMapOptionsBuilder {
	resource := NewDashboardv2RegexMapOptions()
	builder := &Dashboardv2RegexMapOptionsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *Dashboardv2RegexMapOptionsBuilder) Build() (Dashboardv2RegexMapOptions, error) {
	if err := builder.internal.Validate(); err != nil {
		return Dashboardv2RegexMapOptions{}, err
	}

	if len(builder.errors) > 0 {
		return Dashboardv2RegexMapOptions{}, cog.MakeBuildErrors("dashboardv2.dashboardv2RegexMapOptions", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *Dashboardv2RegexMapOptionsBuilder) RecordError(path string, err error) *Dashboardv2RegexMapOptionsBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

// Regular expression to match against
func (builder *Dashboardv2RegexMapOptionsBuilder) Pattern(pattern string) *Dashboardv2RegexMapOptionsBuilder {
	builder.internal.Pattern = pattern

	return builder
}

// Config to apply when the value matches the regex
func (builder *Dashboardv2RegexMapOptionsBuilder) Result(result cog.Builder[ValueMappingResult]) *Dashboardv2RegexMapOptionsBuilder {
	resultResource, err := result.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Result = resultResource

	return builder
}
