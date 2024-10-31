---
title: <span class="badge builder"></span> Preferences
---
# <span class="badge builder"></span> Preferences

## Constructor

```python
Preferences()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> preferences.Preferences
```

### <span class="badge object-method"></span> cookie_preferences

Cookie preferences

```python
def cookie_preferences(cookie_preferences: cogbuilder.Builder[preferences.CookiePreferences]) -> typing.Self
```

### <span class="badge object-method"></span> home_dashboard_uid

UID for the home dashboard

```python
def home_dashboard_uid(home_dashboard_uid: str) -> typing.Self
```

### <span class="badge object-method"></span> language

Selected language (beta)

```python
def language(language: str) -> typing.Self
```

### <span class="badge object-method"></span> navbar

Navigation preferences

```python
def navbar(navbar: cogbuilder.Builder[preferences.NavbarPreference]) -> typing.Self
```

### <span class="badge object-method"></span> query_history

Explore query history preferences

```python
def query_history(query_history: cogbuilder.Builder[preferences.QueryHistoryPreference]) -> typing.Self
```

### <span class="badge object-method"></span> theme

light, dark, empty is default

```python
def theme(theme: str) -> typing.Self
```

### <span class="badge object-method"></span> timezone

The timezone selection

TODO: this should use the timezone defined in common

```python
def timezone(timezone: str) -> typing.Self
```

### <span class="badge object-method"></span> week_start

day of the week (sunday, monday, etc)

```python
def week_start(week_start: str) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [Preferences](./object-Preferences.md)
