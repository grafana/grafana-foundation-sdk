---
title: <span class="badge builder"></span> PreferencesBuilder
---
# <span class="badge builder"></span> PreferencesBuilder

## Constructor

```go
func NewPreferencesBuilder() *PreferencesBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *PreferencesBuilder) Build() (Preferences, error)
```

### <span class="badge object-method"></span> CookiePreferences

Cookie preferences

```go
func (builder *PreferencesBuilder) CookiePreferences(cookiePreferences cog.Builder[preferences.CookiePreferences]) *PreferencesBuilder
```

### <span class="badge object-method"></span> HomeDashboardUID

UID for the home dashboard

```go
func (builder *PreferencesBuilder) HomeDashboardUID(homeDashboardUID string) *PreferencesBuilder
```

### <span class="badge object-method"></span> Language

Selected language (beta)

```go
func (builder *PreferencesBuilder) Language(language string) *PreferencesBuilder
```

### <span class="badge object-method"></span> Navbar

Navigation preferences

```go
func (builder *PreferencesBuilder) Navbar(navbar cog.Builder[preferences.NavbarPreference]) *PreferencesBuilder
```

### <span class="badge object-method"></span> QueryHistory

Explore query history preferences

```go
func (builder *PreferencesBuilder) QueryHistory(queryHistory cog.Builder[preferences.QueryHistoryPreference]) *PreferencesBuilder
```

### <span class="badge object-method"></span> Theme

light, dark, empty is default

```go
func (builder *PreferencesBuilder) Theme(theme string) *PreferencesBuilder
```

### <span class="badge object-method"></span> Timezone

The timezone selection

TODO: this should use the timezone defined in common

```go
func (builder *PreferencesBuilder) Timezone(timezone string) *PreferencesBuilder
```

### <span class="badge object-method"></span> WeekStart

day of the week (sunday, monday, etc)

```go
func (builder *PreferencesBuilder) WeekStart(weekStart string) *PreferencesBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Preferences](./object-Preferences.md)
