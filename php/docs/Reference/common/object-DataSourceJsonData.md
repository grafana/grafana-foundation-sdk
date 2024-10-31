---
title: <span class="badge object-type-class"></span> DataSourceJsonData
---
# <span class="badge object-type-class"></span> DataSourceJsonData

TODO docs

## Definition

```php
class DataSourceJsonData implements \JsonSerializable
{
    public ?string $authType;

    public ?string $defaultRegion;

    public ?string $profile;

    public ?bool $manageAlerts;

    public ?string $alertmanagerUid;

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

 * <span class="badge builder"></span> [DataSourceJsonDataBuilder](./builder-DataSourceJsonDataBuilder.md)
