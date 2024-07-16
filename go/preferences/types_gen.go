// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package preferences

// Spec defines user, team or org Grafana preferences
// swagger:model Preferences
type Preferences struct {
	// UID for the home dashboard
	HomeDashboardUID *string `json:"homeDashboardUID,omitempty"`
	// The timezone selection
	// TODO: this should use the timezone defined in common
	Timezone *string `json:"timezone,omitempty"`
	// day of the week (sunday, monday, etc)
	WeekStart *string `json:"weekStart,omitempty"`
	// light, dark, empty is default
	Theme *string `json:"theme,omitempty"`
	// Selected language (beta)
	Language *string `json:"language,omitempty"`
	// Explore query history preferences
	QueryHistory *QueryHistoryPreference `json:"queryHistory,omitempty"`
	// Cookie preferences
	CookiePreferences *CookiePreferences `json:"cookiePreferences,omitempty"`
	// Navigation preferences
	Navbar *NavbarPreference `json:"navbar,omitempty"`
}

type QueryHistoryPreference struct {
	// one of: '' | 'query' | 'starred';
	HomeTab *string `json:"homeTab,omitempty"`
}

type CookiePreferences struct {
	Analytics   any `json:"analytics,omitempty"`
	Performance any `json:"performance,omitempty"`
	Functional  any `json:"functional,omitempty"`
}

type NavbarPreference struct {
	SavedItemIds []string `json:"savedItemIds"`
}
