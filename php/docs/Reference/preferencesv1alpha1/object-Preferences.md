---
title: <span class="badge object-type-class"></span> Preferences
---
# <span class="badge object-type-class"></span> Preferences

## Definition

```php
class Preferences implements \JsonSerializable
{
    /**
     * Explicit home URL (NOTE: this can only be modified in the system settings)
     */
    public ?string $homeURL;

    /**
     * UID for the home dashboard
     */
    public ?string $homeDashboardUID;

    /**
     * The timezone selection
     */
    public ?string $timezone;

    /**
     * day of the week (sunday, monday, etc)
     */
    public ?string $weekStart;

    /**
     * user interface theme
     */
    public ?string $theme;

    /**
     * Selected language
     */
    public ?string $language;

    /**
     * Explore query history preferences
     */
    public ?\Grafana\Foundation\Preferencesv1alpha1\QueryHistoryPreference $queryHistory;

    /**
     * Navigation preferences
     */
    public ?\Grafana\Foundation\Preferencesv1alpha1\NavbarPreference $navbar;

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
