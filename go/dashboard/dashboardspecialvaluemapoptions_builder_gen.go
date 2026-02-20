// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboard

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[DashboardSpecialValueMapOptions] = (*DashboardSpecialValueMapOptionsBuilder)(nil)

type DashboardSpecialValueMapOptionsBuilder struct {
	internal *DashboardSpecialValueMapOptions
	errors   cog.BuildErrors
}

func NewDashboardSpecialValueMapOptionsBuilder() *DashboardSpecialValueMapOptionsBuilder {
	resource := NewDashboardSpecialValueMapOptions()
	builder := &DashboardSpecialValueMapOptionsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *DashboardSpecialValueMapOptionsBuilder) Build() (DashboardSpecialValueMapOptions, error) {
	if err := builder.internal.Validate(); err != nil {
		return DashboardSpecialValueMapOptions{}, err
	}

	if len(builder.errors) > 0 {
		return DashboardSpecialValueMapOptions{}, cog.MakeBuildErrors("dashboard.dashboardSpecialValueMapOptions", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *DashboardSpecialValueMapOptionsBuilder) RecordError(path string, err error) *DashboardSpecialValueMapOptionsBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

// Special value to match against
func (builder *DashboardSpecialValueMapOptionsBuilder) Match(match SpecialValueMatch) *DashboardSpecialValueMapOptionsBuilder {
	builder.internal.Match = match

	return builder
}

// Config to apply when the value matches the special value
func (builder *DashboardSpecialValueMapOptionsBuilder) Result(result ValueMappingResult) *DashboardSpecialValueMapOptionsBuilder {
	builder.internal.Result = result

	return builder
}
