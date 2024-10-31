---
title: <span class="badge object-type-class"></span> PublicDashboard
---
# <span class="badge object-type-class"></span> PublicDashboard

## Definition

```php
class PublicDashboard implements \JsonSerializable
{
    /**
     * Unique public dashboard identifier
     */
    public string $uid;

    /**
     * Dashboard unique identifier referenced by this public dashboard
     */
    public string $dashboardUid;

    /**
     * Unique public access token
     */
    public ?string $accessToken;

    /**
     * Flag that indicates if the public dashboard is enabled
     */
    public bool $isEnabled;

    /**
     * Flag that indicates if annotations are enabled
     */
    public bool $annotationsEnabled;

    /**
     * Flag that indicates if the time range picker is enabled
     */
    public bool $timeSelectionEnabled;

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

 * <span class="badge builder"></span> [PublicDashboardBuilder](./builder-PublicDashboardBuilder.md)
