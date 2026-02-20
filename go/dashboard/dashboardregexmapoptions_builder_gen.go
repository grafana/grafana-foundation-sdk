// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboard

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[DashboardRegexMapOptions] = (*DashboardRegexMapOptionsBuilder)(nil)

type DashboardRegexMapOptionsBuilder struct {
	internal *DashboardRegexMapOptions
	errors   cog.BuildErrors
}

func NewDashboardRegexMapOptionsBuilder() *DashboardRegexMapOptionsBuilder {
	resource := NewDashboardRegexMapOptions()
	builder := &DashboardRegexMapOptionsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *DashboardRegexMapOptionsBuilder) Build() (DashboardRegexMapOptions, error) {
	if err := builder.internal.Validate(); err != nil {
		return DashboardRegexMapOptions{}, err
	}

	if len(builder.errors) > 0 {
		return DashboardRegexMapOptions{}, cog.MakeBuildErrors("dashboard.dashboardRegexMapOptions", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *DashboardRegexMapOptionsBuilder) RecordError(path string, err error) *DashboardRegexMapOptionsBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

// Regular expression to match against
func (builder *DashboardRegexMapOptionsBuilder) Pattern(pattern string) *DashboardRegexMapOptionsBuilder {
	builder.internal.Pattern = pattern

	return builder
}

// Config to apply when the value matches the regex
func (builder *DashboardRegexMapOptionsBuilder) Result(result ValueMappingResult) *DashboardRegexMapOptionsBuilder {
	builder.internal.Result = result

	return builder
}
