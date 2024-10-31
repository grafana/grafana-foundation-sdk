# preferences

## Objects

 * <span class="badge object-type-struct"></span> [CookiePreferences](./object-CookiePreferences.md)
 * <span class="badge object-type-struct"></span> [NavbarPreference](./object-NavbarPreference.md)
 * <span class="badge object-type-struct"></span> [Preferences](./object-Preferences.md)
 * <span class="badge object-type-struct"></span> [QueryHistoryPreference](./object-QueryHistoryPreference.md)
## Builders

 * <span class="badge builder"></span> [CookiePreferencesBuilder](./builder-CookiePreferencesBuilder.md)
 * <span class="badge builder"></span> [NavbarPreferenceBuilder](./builder-NavbarPreferenceBuilder.md)
 * <span class="badge builder"></span> [PreferencesBuilder](./builder-PreferencesBuilder.md)
 * <span class="badge builder"></span> [QueryHistoryPreferenceBuilder](./builder-QueryHistoryPreferenceBuilder.md)
## Functions

### <span class="badge function"></span> PreferencesConverter

PreferencesConverter accepts a `Preferences` object and generates the Go code to build this object using builders.

```go
func PreferencesConverter(input Preferences) string
```

### <span class="badge function"></span> QueryHistoryPreferenceConverter

QueryHistoryPreferenceConverter accepts a `QueryHistoryPreference` object and generates the Go code to build this object using builders.

```go
func QueryHistoryPreferenceConverter(input QueryHistoryPreference) string
```

### <span class="badge function"></span> CookiePreferencesConverter

CookiePreferencesConverter accepts a `CookiePreferences` object and generates the Go code to build this object using builders.

```go
func CookiePreferencesConverter(input CookiePreferences) string
```

### <span class="badge function"></span> NavbarPreferenceConverter

NavbarPreferenceConverter accepts a `NavbarPreference` object and generates the Go code to build this object using builders.

```go
func NavbarPreferenceConverter(input NavbarPreference) string
```

