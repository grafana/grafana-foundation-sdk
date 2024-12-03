// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboard

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[DashboardSpecialValueMapOptions] = (*DashboardSpecialValueMapOptionsBuilder)(nil)

type DashboardSpecialValueMapOptionsBuilder struct {
	internal *DashboardSpecialValueMapOptions
	errors   map[string]cog.BuildErrors
}

func NewDashboardSpecialValueMapOptionsBuilder() *DashboardSpecialValueMapOptionsBuilder {
	resource := NewDashboardSpecialValueMapOptions()
	builder := &DashboardSpecialValueMapOptionsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *DashboardSpecialValueMapOptionsBuilder) Build() (DashboardSpecialValueMapOptions, error) {
	if err := builder.internal.Validate(); err != nil {
		return DashboardSpecialValueMapOptions{}, err
	}

	return *builder.internal, nil
}

// Special value to match against
func (builder *DashboardSpecialValueMapOptionsBuilder) Match(match SpecialValueMatch) *DashboardSpecialValueMapOptionsBuilder {
	builder.internal.Match = match

	return builder
}

// Config to apply when the value matches the special value
func (builder *DashboardSpecialValueMapOptionsBuilder) Result(result cog.Builder[ValueMappingResult]) *DashboardSpecialValueMapOptionsBuilder {
	resultResource, err := result.Build()
	if err != nil {
		builder.errors["result"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Result = resultResource

	return builder
}