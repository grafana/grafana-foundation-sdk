// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package preferences

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[QueryHistoryPreference] = (*QueryHistoryPreferenceBuilder)(nil)

type QueryHistoryPreferenceBuilder struct {
	internal *QueryHistoryPreference
	errors   cog.BuildErrors
}

func NewQueryHistoryPreferenceBuilder() *QueryHistoryPreferenceBuilder {
	resource := NewQueryHistoryPreference()
	builder := &QueryHistoryPreferenceBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *QueryHistoryPreferenceBuilder) Build() (QueryHistoryPreference, error) {
	if err := builder.internal.Validate(); err != nil {
		return QueryHistoryPreference{}, err
	}

	if len(builder.errors) > 0 {
		return QueryHistoryPreference{}, cog.MakeBuildErrors("preferences.queryHistoryPreference", builder.errors)
	}

	return *builder.internal, nil
}

// one of: ” | 'query' | 'starred';
func (builder *QueryHistoryPreferenceBuilder) HomeTab(homeTab string) *QueryHistoryPreferenceBuilder {
	builder.internal.HomeTab = &homeTab

	return builder
}
