---
title: <span class="badge object-type-class"></span> <span class="badge deprecated"></span> Preferences
---
# <span class="badge object-type-class"></span> <span class="badge deprecated"></span> Preferences

<span class="badge deprecated"></span>Prefer using preferencesv1alpha1.Preferences instead.

Spec defines user, team or org Grafana preferences

swagger:model Preferences

## Definition

```python
class Preferences:
    """
    Spec defines user, team or org Grafana preferences
    swagger:model Preferences
    """

    warnings.warn("Prefer using preferencesv1alpha1.Preferences instead.", DeprecationWarning)
    # UID for the home dashboard
    home_dashboard_uid: typing.Optional[str]
    # The timezone selection
    # TODO: this should use the timezone defined in common
    timezone: typing.Optional[str]
    # day of the week (sunday, monday, etc)
    week_start: typing.Optional[str]
    # light, dark, empty is default
    theme: typing.Optional[str]
    # Selected language (beta)
    language: typing.Optional[str]
    # Explore query history preferences
    query_history: typing.Optional[preferences.QueryHistoryPreference]
    # Cookie preferences
    cookie_preferences: typing.Optional[preferences.CookiePreferences]
    # Navigation preferences
    navbar: typing.Optional[preferences.NavbarPreference]
```
## Methods

### <span class="badge object-method"></span> to_json

Converts this object into a representation that can easily be encoded to JSON.

```python
def to_json() -> dict[str, object]
```

### <span class="badge object-method"></span> from_json

Builds this object from a JSON-decoded dict.

```python
@classmethod
def from_json(data: dict[str, typing.Any]) -> typing.Self
```

## See also

 * <span class="badge builder"></span> <span class="badge deprecated"></span> [Preferences](./builder-Preferences.md)
