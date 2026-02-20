// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package preferencesv1alpha1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[CookiePreferences] = (*CookiePreferencesBuilder)(nil)

type CookiePreferencesBuilder struct {
	internal *CookiePreferences
	errors   cog.BuildErrors
}

func NewCookiePreferencesBuilder() *CookiePreferencesBuilder {
	resource := NewCookiePreferences()
	builder := &CookiePreferencesBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *CookiePreferencesBuilder) Build() (CookiePreferences, error) {
	if err := builder.internal.Validate(); err != nil {
		return CookiePreferences{}, err
	}

	if len(builder.errors) > 0 {
		return CookiePreferences{}, cog.MakeBuildErrors("preferencesv1alpha1.cookiePreferences", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *CookiePreferencesBuilder) RecordError(path string, err error) *CookiePreferencesBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
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
