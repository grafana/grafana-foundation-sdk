---
title: <span class="badge object-type-struct"></span> Preferences
---
# <span class="badge object-type-struct"></span> Preferences

## Definition

```go
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
    QueryHistory *preferences.QueryHistoryPreference `json:"queryHistory,omitempty"`
    // Cookie preferences
    CookiePreferences *preferences.CookiePreferences `json:"cookiePreferences,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Preferences` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (preferences *Preferences) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `Preferences` objects.

```go
func (preferences *Preferences) Equals(other Preferences) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `Preferences` fields for violations and returns them.

```go
func (preferences *Preferences) Validate() error
```

## See also

 * <span class="badge builder"></span> [PreferencesBuilder](./builder-PreferencesBuilder.md)
