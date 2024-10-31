// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package preferences

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[CookiePreferences] = (*CookiePreferencesBuilder)(nil)

type CookiePreferencesBuilder struct {
	internal *CookiePreferences
	errors   map[string]cog.BuildErrors
}

func NewCookiePreferencesBuilder() *CookiePreferencesBuilder {
	resource := &CookiePreferences{}
	builder := &CookiePreferencesBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *CookiePreferencesBuilder) Build() (CookiePreferences, error) {
	if err := builder.internal.Validate(); err != nil {
		return CookiePreferences{}, err
	}

	return *builder.internal, nil
}

func (builder *CookiePreferencesBuilder) Analytics(analytics any) *CookiePreferencesBuilder {
	builder.internal.Analytics = &analytics

	return builder
}

func (builder *CookiePreferencesBuilder) Performance(performance any) *CookiePreferencesBuilder {
	builder.internal.Performance = &performance

	return builder
}

func (builder *CookiePreferencesBuilder) Functional(functional any) *CookiePreferencesBuilder {
	builder.internal.Functional = &functional

	return builder
}

func (builder *CookiePreferencesBuilder) applyDefaults() {
}
