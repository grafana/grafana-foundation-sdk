// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Preferences] = (*PreferencesBuilder)(nil)

// Dashboard specific preferences (applied per dashboard = all users using the dashboard)
type PreferencesBuilder struct {
	internal *Preferences
	errors   cog.BuildErrors
}

func NewPreferencesBuilder() *PreferencesBuilder {
	resource := NewPreferences()
	builder := &PreferencesBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *PreferencesBuilder) Build() (Preferences, error) {
	if err := builder.internal.Validate(); err != nil {
		return Preferences{}, err
	}

	if len(builder.errors) > 0 {
		return Preferences{}, cog.MakeBuildErrors("dashboardv2.preferences", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *PreferencesBuilder) RecordError(path string, err error) *PreferencesBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

// default layout template to be used when new containers are created
func (builder *PreferencesBuilder) Layout(layout AutoGridLayoutKindOrGridLayoutKind) *PreferencesBuilder {
	builder.internal.Layout = &layout

	return builder
}
