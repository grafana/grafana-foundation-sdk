---
title: <span class="badge builder"></span> PreferencesBuilder
---
# <span class="badge builder"></span> PreferencesBuilder

## Constructor

```php
new PreferencesBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> cookiePreferences

Cookie preferences

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Preferences\CookiePreferences> $cookiePreferences

```php
cookiePreferences(\Grafana\Foundation\Cog\Builder $cookiePreferences)
```

### <span class="badge object-method"></span> homeDashboardUID

UID for the home dashboard

```php
homeDashboardUID(string $homeDashboardUID)
```

### <span class="badge object-method"></span> language

Selected language (beta)

```php
language(string $language)
```

### <span class="badge object-method"></span> navbar

Navigation preferences

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Preferences\NavbarPreference> $navbar

```php
navbar(\Grafana\Foundation\Cog\Builder $navbar)
```

### <span class="badge object-method"></span> queryHistory

Explore query history preferences

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Preferences\QueryHistoryPreference> $queryHistory

```php
queryHistory(\Grafana\Foundation\Cog\Builder $queryHistory)
```

### <span class="badge object-method"></span> theme

light, dark, empty is default

```php
theme(string $theme)
```

### <span class="badge object-method"></span> timezone

The timezone selection

TODO: this should use the timezone defined in common

```php
timezone(string $timezone)
```

### <span class="badge object-method"></span> weekStart

day of the week (sunday, monday, etc)

```php
weekStart(string $weekStart)
```

## See also

 * <span class="badge object-type-class"></span> [Preferences](./object-Preferences.md)
