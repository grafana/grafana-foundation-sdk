// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboard

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[DashboardRegexMapOptions] = (*DashboardRegexMapOptionsBuilder)(nil)

type DashboardRegexMapOptionsBuilder struct {
	internal *DashboardRegexMapOptions
	errors   map[string]cog.BuildErrors
}

func NewDashboardRegexMapOptionsBuilder() *DashboardRegexMapOptionsBuilder {
	resource := NewDashboardRegexMapOptions()
	builder := &DashboardRegexMapOptionsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *DashboardRegexMapOptionsBuilder) Build() (DashboardRegexMapOptions, error) {
	if err := builder.internal.Validate(); err != nil {
		return DashboardRegexMapOptions{}, err
	}

	return *builder.internal, nil
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
