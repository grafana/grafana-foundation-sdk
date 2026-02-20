// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package preferencesv1alpha1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	resource "github.com/grafana/grafana-foundation-sdk/go/resource"
)

var _ cog.Builder[Preferences] = (*PreferencesBuilder)(nil)

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

// Creates a resource manifest from Preferences.
func Manifest(name string, preferences cog.Builder[Preferences]) *resource.ManifestBuilder {
	builder := resource.NewManifestBuilder()

	builder.ApiVersion("preferences.grafana.app/v1alpha1")

	builder.Kind("Preferences")

	builder.Metadata(resource.Named(name))
	preferencesResource, err := preferences.Build()
	if err != nil {
		builder.RecordError("Manifest(preferences)", err)
		return builder
	}
	builder.Spec(preferencesResource)

	return builder
}

func (builder *PreferencesBuilder) Build() (Preferences, error) {
	if err := builder.internal.Validate(); err != nil {
		return Preferences{}, err
	}

	if len(builder.errors) > 0 {
		return Preferences{}, cog.MakeBuildErrors("preferencesv1alpha1.preferences", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *PreferencesBuilder) RecordError(path string, err error) *PreferencesBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

// UID for the home dashboard
func (builder *PreferencesBuilder) HomeDashboardUID(homeDashboardUID string) *PreferencesBuilder {
	builder.internal.HomeDashboardUID = &homeDashboardUID

	return builder
}

// The timezone selection
// TODO: this should use the timezone defined in common
func (builder *PreferencesBuilder) Timezone(timezone string) *PreferencesBuilder {
	builder.internal.Timezone = &timezone

	return builder
}

// day of the week (sunday, monday, etc)
func (builder *PreferencesBuilder) WeekStart(weekStart string) *PreferencesBuilder {
	builder.internal.WeekStart = &weekStart

	return builder
}

// light, dark, empty is default
func (builder *PreferencesBuilder) Theme(theme string) *PreferencesBuilder {
	builder.internal.Theme = &theme

	return builder
}

// Selected language (beta)
func (builder *PreferencesBuilder) Language(language string) *PreferencesBuilder {
	builder.internal.Language = &language

	return builder
}

// Selected locale (beta)
func (builder *PreferencesBuilder) RegionalFormat(regionalFormat string) *PreferencesBuilder {
	builder.internal.RegionalFormat = &regionalFormat

	return builder
}

// Explore query history preferences
func (builder *PreferencesBuilder) QueryHistory(queryHistory cog.Builder[QueryHistoryPreference]) *PreferencesBuilder {
	queryHistoryResource, err := queryHistory.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.QueryHistory = &queryHistoryResource

	return builder
}

// Cookie preferences
func (builder *PreferencesBuilder) CookiePreferences(cookiePreferences cog.Builder[CookiePreferences]) *PreferencesBuilder {
	cookiePreferencesResource, err := cookiePreferences.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.CookiePreferences = &cookiePreferencesResource

	return builder
}

// Navigation preferences
func (builder *PreferencesBuilder) Navbar(navbar cog.Builder[NavbarPreference]) *PreferencesBuilder {
	navbarResource, err := navbar.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Navbar = &navbarResource

	return builder
}
