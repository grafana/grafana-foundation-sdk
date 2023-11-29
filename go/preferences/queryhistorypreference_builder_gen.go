// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package preferences

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[QueryHistoryPreference] = (*QueryHistoryPreferenceBuilder)(nil)

type QueryHistoryPreferenceBuilder struct {
	internal *QueryHistoryPreference
	errors   map[string]cog.BuildErrors
}

func NewQueryHistoryPreferenceBuilder() *QueryHistoryPreferenceBuilder {
	resource := &QueryHistoryPreference{}
	builder := &QueryHistoryPreferenceBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *QueryHistoryPreferenceBuilder) Build() (QueryHistoryPreference, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("QueryHistoryPreference", err)...)
	}

	if len(errs) != 0 {
		return QueryHistoryPreference{}, errs
	}

	return *builder.internal, nil
}

// one of: ” | 'query' | 'starred';
func (builder *QueryHistoryPreferenceBuilder) HomeTab(homeTab string) *QueryHistoryPreferenceBuilder {
	builder.internal.HomeTab = &homeTab

	return builder
}

func (builder *QueryHistoryPreferenceBuilder) applyDefaults() {
}
