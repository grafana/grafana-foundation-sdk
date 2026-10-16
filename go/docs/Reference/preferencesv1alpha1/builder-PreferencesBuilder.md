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

### <span class="badge object-method"></span> HomeDashboardUID

UID for the home dashboard

```go
func (builder *PreferencesBuilder) HomeDashboardUID(homeDashboardUID string) *PreferencesBuilder
```

### <span class="badge object-method"></span> HomeURL

Explicit home URL (NOTE: this can only be modified in the system settings)

```go
func (builder *PreferencesBuilder) HomeURL(homeURL string) *PreferencesBuilder
```

### <span class="badge object-method"></span> Language

Selected language

```go
func (builder *PreferencesBuilder) Language(language string) *PreferencesBuilder
```

### <span class="badge object-method"></span> Navbar

Navigation preferences

```go
func (builder *PreferencesBuilder) Navbar(navbar cog.Builder[preferencesv1alpha1.NavbarPreference]) *PreferencesBuilder
```

### <span class="badge object-method"></span> QueryHistory

Explore query history preferences

```go
func (builder *PreferencesBuilder) QueryHistory(queryHistory cog.Builder[preferencesv1alpha1.QueryHistoryPreference]) *PreferencesBuilder
```

### <span class="badge object-method"></span> Theme

user interface theme

```go
func (builder *PreferencesBuilder) Theme(theme string) *PreferencesBuilder
```

### <span class="badge object-method"></span> Timezone

The timezone selection

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
