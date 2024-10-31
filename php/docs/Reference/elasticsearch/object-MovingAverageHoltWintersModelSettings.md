---
title: <span class="badge object-type-class"></span> MovingAverageHoltWintersModelSettings
---
# <span class="badge object-type-class"></span> MovingAverageHoltWintersModelSettings

## Definition

```php
class MovingAverageHoltWintersModelSettings implements \JsonSerializable
{
    public string $model;

    public \Grafana\Foundation\Elasticsearch\ElasticsearchMovingAverageHoltWintersModelSettingsSettings $settings;

    public string $window;

    public bool $minimize;

    public string $predict;

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

 * <span class="badge builder"></span> [MovingAverageHoltWintersModelSettingsBuilder](./builder-MovingAverageHoltWintersModelSettingsBuilder.md)
