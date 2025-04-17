// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package preferences

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

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
	// Selected locale (beta)
	Locale *string `json:"locale,omitempty"`
	// Explore query history preferences
	QueryHistory *QueryHistoryPreference `json:"queryHistory,omitempty"`
	// Cookie preferences
	CookiePreferences *CookiePreferences `json:"cookiePreferences,omitempty"`
	// Navigation preferences
	Navbar *NavbarPreference `json:"navbar,omitempty"`
}

// NewPreferences creates a new Preferences object.
func NewPreferences() *Preferences {
	return &Preferences{}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Preferences` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *Preferences) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "homeDashboardUID"
	if fields["homeDashboardUID"] != nil {
		if string(fields["homeDashboardUID"]) != "null" {
			if err := json.Unmarshal(fields["homeDashboardUID"], &resource.HomeDashboardUID); err != nil {
				errs = append(errs, cog.MakeBuildErrors("homeDashboardUID", err)...)
			}

		}
		delete(fields, "homeDashboardUID")

	}
	// Field "timezone"
	if fields["timezone"] != nil {
		if string(fields["timezone"]) != "null" {
			if err := json.Unmarshal(fields["timezone"], &resource.Timezone); err != nil {
				errs = append(errs, cog.MakeBuildErrors("timezone", err)...)
			}

		}
		delete(fields, "timezone")

	}
	// Field "weekStart"
	if fields["weekStart"] != nil {
		if string(fields["weekStart"]) != "null" {
			if err := json.Unmarshal(fields["weekStart"], &resource.WeekStart); err != nil {
				errs = append(errs, cog.MakeBuildErrors("weekStart", err)...)
			}

		}
		delete(fields, "weekStart")

	}
	// Field "theme"
	if fields["theme"] != nil {
		if string(fields["theme"]) != "null" {
			if err := json.Unmarshal(fields["theme"], &resource.Theme); err != nil {
				errs = append(errs, cog.MakeBuildErrors("theme", err)...)
			}

		}
		delete(fields, "theme")

	}
	// Field "language"
	if fields["language"] != nil {
		if string(fields["language"]) != "null" {
			if err := json.Unmarshal(fields["language"], &resource.Language); err != nil {
				errs = append(errs, cog.MakeBuildErrors("language", err)...)
			}

		}
		delete(fields, "language")

	}
	// Field "locale"
	if fields["locale"] != nil {
		if string(fields["locale"]) != "null" {
			if err := json.Unmarshal(fields["locale"], &resource.Locale); err != nil {
				errs = append(errs, cog.MakeBuildErrors("locale", err)...)
			}

		}
		delete(fields, "locale")

	}
	// Field "queryHistory"
	if fields["queryHistory"] != nil {
		if string(fields["queryHistory"]) != "null" {

			resource.QueryHistory = &QueryHistoryPreference{}
			if err := resource.QueryHistory.UnmarshalJSONStrict(fields["queryHistory"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("queryHistory", err)...)
			}

		}
		delete(fields, "queryHistory")

	}
	// Field "cookiePreferences"
	if fields["cookiePreferences"] != nil {
		if string(fields["cookiePreferences"]) != "null" {

			resource.CookiePreferences = &CookiePreferences{}
			if err := resource.CookiePreferences.UnmarshalJSONStrict(fields["cookiePreferences"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("cookiePreferences", err)...)
			}

		}
		delete(fields, "cookiePreferences")

	}
	// Field "navbar"
	if fields["navbar"] != nil {
		if string(fields["navbar"]) != "null" {

			resource.Navbar = &NavbarPreference{}
			if err := resource.Navbar.UnmarshalJSONStrict(fields["navbar"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("navbar", err)...)
			}

		}
		delete(fields, "navbar")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("Preferences", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `Preferences` objects.
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
	if resource.Locale == nil && other.Locale != nil || resource.Locale != nil && other.Locale == nil {
		return false
	}

	if resource.Locale != nil {
		if *resource.Locale != *other.Locale {
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

// Validate checks all the validation constraints that may be defined on `Preferences` fields for violations and returns them.
func (resource Preferences) Validate() error {
	var errs cog.BuildErrors
	if resource.QueryHistory != nil {
		if err := resource.QueryHistory.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("queryHistory", err)...)
		}
	}
	if resource.CookiePreferences != nil {
		if err := resource.CookiePreferences.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("cookiePreferences", err)...)
		}
	}
	if resource.Navbar != nil {
		if err := resource.Navbar.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("navbar", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type QueryHistoryPreference struct {
	// one of: '' | 'query' | 'starred';
	HomeTab *string `json:"homeTab,omitempty"`
}

// NewQueryHistoryPreference creates a new QueryHistoryPreference object.
func NewQueryHistoryPreference() *QueryHistoryPreference {
	return &QueryHistoryPreference{}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `QueryHistoryPreference` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *QueryHistoryPreference) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "homeTab"
	if fields["homeTab"] != nil {
		if string(fields["homeTab"]) != "null" {
			if err := json.Unmarshal(fields["homeTab"], &resource.HomeTab); err != nil {
				errs = append(errs, cog.MakeBuildErrors("homeTab", err)...)
			}

		}
		delete(fields, "homeTab")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("QueryHistoryPreference", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `QueryHistoryPreference` objects.
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

// Validate checks all the validation constraints that may be defined on `QueryHistoryPreference` fields for violations and returns them.
func (resource QueryHistoryPreference) Validate() error {
	return nil
}

type CookiePreferences struct {
	Analytics   any `json:"analytics,omitempty"`
	Performance any `json:"performance,omitempty"`
	Functional  any `json:"functional,omitempty"`
}

// NewCookiePreferences creates a new CookiePreferences object.
func NewCookiePreferences() *CookiePreferences {
	return &CookiePreferences{}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `CookiePreferences` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *CookiePreferences) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "analytics"
	if fields["analytics"] != nil {
		if string(fields["analytics"]) != "null" {
			if err := json.Unmarshal(fields["analytics"], &resource.Analytics); err != nil {
				errs = append(errs, cog.MakeBuildErrors("analytics", err)...)
			}

		}
		delete(fields, "analytics")

	}
	// Field "performance"
	if fields["performance"] != nil {
		if string(fields["performance"]) != "null" {
			if err := json.Unmarshal(fields["performance"], &resource.Performance); err != nil {
				errs = append(errs, cog.MakeBuildErrors("performance", err)...)
			}

		}
		delete(fields, "performance")

	}
	// Field "functional"
	if fields["functional"] != nil {
		if string(fields["functional"]) != "null" {
			if err := json.Unmarshal(fields["functional"], &resource.Functional); err != nil {
				errs = append(errs, cog.MakeBuildErrors("functional", err)...)
			}

		}
		delete(fields, "functional")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("CookiePreferences", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `CookiePreferences` objects.
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

// Validate checks all the validation constraints that may be defined on `CookiePreferences` fields for violations and returns them.
func (resource CookiePreferences) Validate() error {
	return nil
}

type NavbarPreference struct {
	BookmarkUrls []string `json:"bookmarkUrls"`
}

// NewNavbarPreference creates a new NavbarPreference object.
func NewNavbarPreference() *NavbarPreference {
	return &NavbarPreference{
		BookmarkUrls: []string{},
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `NavbarPreference` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *NavbarPreference) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "bookmarkUrls"
	if fields["bookmarkUrls"] != nil {
		if string(fields["bookmarkUrls"]) != "null" {

			if err := json.Unmarshal(fields["bookmarkUrls"], &resource.BookmarkUrls); err != nil {
				errs = append(errs, cog.MakeBuildErrors("bookmarkUrls", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("bookmarkUrls", errors.New("required field is null"))...)

		}
		delete(fields, "bookmarkUrls")
	} else {
		errs = append(errs, cog.MakeBuildErrors("bookmarkUrls", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("NavbarPreference", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `NavbarPreference` objects.
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

// Validate checks all the validation constraints that may be defined on `NavbarPreference` fields for violations and returns them.
func (resource NavbarPreference) Validate() error {
	return nil
}
