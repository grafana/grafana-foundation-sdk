---
title: <span class="badge object-type-class"></span> Preferences
---
# <span class="badge object-type-class"></span> Preferences

Dashboard specific preferences (applied per dashboard = all users using the dashboard)

## Definition

```php
class Preferences implements \JsonSerializable
{
    /**
     * default layout template to be used when new containers are created
     * @var \Grafana\Foundation\Dashboardv2\AutoGridLayoutKind|\Grafana\Foundation\Dashboardv2\GridLayoutKind|null
     */
    public $layout;

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
