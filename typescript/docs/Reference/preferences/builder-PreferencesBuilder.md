---
title: <span class="badge builder"></span> PreferencesBuilder
---
# <span class="badge builder"></span> PreferencesBuilder

## Constructor

```typescript
new PreferencesBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> cookiePreferences

Cookie preferences

```typescript
cookiePreferences(cookiePreferences: cog.Builder<preferences.CookiePreferences>)
```

### <span class="badge object-method"></span> homeDashboardUID

UID for the home dashboard

```typescript
homeDashboardUID(homeDashboardUID: string)
```

### <span class="badge object-method"></span> language

Selected language (beta)

```typescript
language(language: string)
```

### <span class="badge object-method"></span> navbar

Navigation preferences

```typescript
navbar(navbar: cog.Builder<preferences.NavbarPreference>)
```

### <span class="badge object-method"></span> queryHistory

Explore query history preferences

```typescript
queryHistory(queryHistory: cog.Builder<preferences.QueryHistoryPreference>)
```

### <span class="badge object-method"></span> theme

light, dark, empty is default

```typescript
theme(theme: string)
```

### <span class="badge object-method"></span> timezone

The timezone selection

TODO: this should use the timezone defined in common

```typescript
timezone(timezone: string)
```

### <span class="badge object-method"></span> weekStart

day of the week (sunday, monday, etc)

```typescript
weekStart(weekStart: string)
```

## See also

 * <span class="badge object-type-interface"></span> [Preferences](./object-Preferences.md)
