---
title: <span class="badge object-type-class"></span> Preferences
---
# <span class="badge object-type-class"></span> Preferences

## Definition

```php
class Preferences implements \JsonSerializable
{
    /**
     * UID for the home dashboard
     */
    public ?string $homeDashboardUID;

    /**
     * The timezone selection
     * TODO: this should use the timezone defined in common
     */
    public ?string $timezone;

    /**
     * day of the week (sunday, monday, etc)
     */
    public ?string $weekStart;

    /**
     * light, dark, empty is default
     */
    public ?string $theme;

    /**
     * Selected language (beta)
     */
    public ?string $language;

    /**
     * Explore query history preferences
     */
    public ?\Grafana\Foundation\Preferences\QueryHistoryPreference $queryHistory;

    /**
     * Cookie preferences
     */
    public ?\Grafana\Foundation\Preferences\CookiePreferences $cookiePreferences;

}
```
## Methods

### <span class="badge object-method"></span> fromArray

Builds this object from an array.

This function is meant to be used with the return value of `json_decode($json, true)`.

```php
static fromArray(array $inputData)
```

### <span class="badge object-method"></span> jsonSerialize

Returns the data representing this object, preparing it for JSON serialization with `json_encode()`.

```php
jsonSerialize()
```

## See also

 * <span class="badge builder"></span> [PreferencesBuilder](./builder-PreferencesBuilder.md)
