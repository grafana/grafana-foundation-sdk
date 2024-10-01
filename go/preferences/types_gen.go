// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package preferences

import "reflect"

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

func (resource Preferences) Equals(other Preferences) bool {
	if resource.HomeDashboardUID == nil && other.HomeDashboardUID != nil || resource.HomeDashboardUID != nil && other.HomeDashboardUID == nil {
		return false
	}

	if resource.HomeDashboardUID != nil {
		if *resource.HomeDashboardUID != *other.HomeDashboardUID {
			return false
		}
	}
	if resource.Timezone == nil && other.Timezone != nil || resource.Timezone != nil && other.Timezone == nil {
		return false
	}

	if resource.Timezone != nil {
		if *resource.Timezone != *other.Timezone {
			return false
		}
	}
	if resource.WeekStart == nil && other.WeekStart != nil || resource.WeekStart != nil && other.WeekStart == nil {
		return false
	}

	if resource.WeekStart != nil {
		if *resource.WeekStart != *other.WeekStart {
			return false
		}
	}
	if resource.Theme == nil && other.Theme != nil || resource.Theme != nil && other.Theme == nil {
		return false
	}

	if resource.Theme != nil {
		if *resource.Theme != *other.Theme {
			return false
		}
	}
	if resource.Language == nil && other.Language != nil || resource.Language != nil && other.Language == nil {
		return false
	}

	if resource.Language != nil {
		if *resource.Language != *other.Language {
			return false
		}
	}
	if resource.QueryHistory == nil && other.QueryHistory != nil || resource.QueryHistory != nil && other.QueryHistory == nil {
		return false
	}

	if resource.QueryHistory != nil {
		if !resource.QueryHistory.Equals(*other.QueryHistory) {
			return false
		}
	}
	if resource.CookiePreferences == nil && other.CookiePreferences != nil || resource.CookiePreferences != nil && other.CookiePreferences == nil {
		return false
	}

	if resource.CookiePreferences != nil {
		if !resource.CookiePreferences.Equals(*other.CookiePreferences) {
			return false
		}
	}
	if resource.Navbar == nil && other.Navbar != nil || resource.Navbar != nil && other.Navbar == nil {
		return false
	}

	if resource.Navbar != nil {
		if !resource.Navbar.Equals(*other.Navbar) {
			return false
		}
	}

	return true
}

type QueryHistoryPreference struct {
	// one of: '' | 'query' | 'starred';
	HomeTab *string `json:"homeTab,omitempty"`
}

func (resource QueryHistoryPreference) Equals(other QueryHistoryPreference) bool {
	if resource.HomeTab == nil && other.HomeTab != nil || resource.HomeTab != nil && other.HomeTab == nil {
		return false
	}

	if resource.HomeTab != nil {
		if *resource.HomeTab != *other.HomeTab {
			return false
		}
	}

	return true
}

type CookiePreferences struct {
	Analytics   any `json:"analytics,omitempty"`
	Performance any `json:"performance,omitempty"`
	Functional  any `json:"functional,omitempty"`
}

func (resource CookiePreferences) Equals(other CookiePreferences) bool {
	// is DeepEqual good enough here?
	if !reflect.DeepEqual(resource.Analytics, other.Analytics) {
		return false
	}
	// is DeepEqual good enough here?
	if !reflect.DeepEqual(resource.Performance, other.Performance) {
		return false
	}
	// is DeepEqual good enough here?
	if !reflect.DeepEqual(resource.Functional, other.Functional) {
		return false
	}

	return true
}

type NavbarPreference struct {
	BookmarkUrls []string `json:"bookmarkUrls"`
}

func (resource NavbarPreference) Equals(other NavbarPreference) bool {

	if len(resource.BookmarkUrls) != len(other.BookmarkUrls) {
		return false
	}

	for i1 := range resource.BookmarkUrls {
		if resource.BookmarkUrls[i1] != other.BookmarkUrls[i1] {
			return false
		}
	}

	return true
}
