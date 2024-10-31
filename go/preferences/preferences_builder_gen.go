// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package preferences

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Preferences] = (*PreferencesBuilder)(nil)

type PreferencesBuilder struct {
	internal *Preferences
	errors   map[string]cog.BuildErrors
}

func NewPreferencesBuilder() *PreferencesBuilder {
	resource := &Preferences{}
	builder := &PreferencesBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *PreferencesBuilder) Build() (Preferences, error) {
	if err := builder.internal.Validate(); err != nil {
		return Preferences{}, err
	}

	return *builder.internal, nil
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

// Explore query history preferences
func (builder *PreferencesBuilder) QueryHistory(queryHistory cog.Builder[QueryHistoryPreference]) *PreferencesBuilder {
	queryHistoryResource, err := queryHistory.Build()
	if err != nil {
		builder.errors["queryHistory"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.QueryHistory = &queryHistoryResource

	return builder
}

// Cookie preferences
func (builder *PreferencesBuilder) CookiePreferences(cookiePreferences cog.Builder[CookiePreferences]) *PreferencesBuilder {
	cookiePreferencesResource, err := cookiePreferences.Build()
	if err != nil {
		builder.errors["cookiePreferences"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.CookiePreferences = &cookiePreferencesResource

	return builder
}

func (builder *PreferencesBuilder) applyDefaults() {
}
