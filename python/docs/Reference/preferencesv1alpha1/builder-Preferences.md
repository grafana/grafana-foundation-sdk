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
def build() -> preferencesv1alpha1.Preferences
```

### <span class="badge object-method"></span> home_dashboard_uid

UID for the home dashboard

```python
def home_dashboard_uid(home_dashboard_uid: str) -> typing.Self
```

### <span class="badge object-method"></span> home_url

Explicit home URL (NOTE: this can only be modified in the system settings)

```python
def home_url(home_url: str) -> typing.Self
```

### <span class="badge object-method"></span> language

Selected language

```python
def language(language: str) -> typing.Self
```

### <span class="badge object-method"></span> navbar

Navigation preferences

```python
def navbar(navbar: cogbuilder.Builder[preferencesv1alpha1.NavbarPreference]) -> typing.Self
```

### <span class="badge object-method"></span> query_history

Explore query history preferences

```python
def query_history(query_history: cogbuilder.Builder[preferencesv1alpha1.QueryHistoryPreference]) -> typing.Self
```

### <span class="badge object-method"></span> theme

user interface theme

```python
def theme(theme: str) -> typing.Self
```

### <span class="badge object-method"></span> timezone

The timezone selection

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
