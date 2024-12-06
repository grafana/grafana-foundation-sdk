---
title: <span class="badge object-type-class"></span> FieldConfigSource
---
# <span class="badge object-type-class"></span> FieldConfigSource

The data model used in Grafana, namely the data frame, is a columnar-oriented table structure that unifies both time series and table query results.

Each column within this structure is called a field. A field can represent a single time series or table column.

Field options allow you to change how the data is displayed in your visualizations.

## Definition

```php
class FieldConfigSource implements \JsonSerializable
{
    /**
     * Defaults are the options applied to all fields.
     */
    public \Grafana\Foundation\Dashboard\FieldConfig $defaults;

    /**
     * Overrides are the options applied to specific fields overriding the defaults.
     * @var array<\Grafana\Foundation\Dashboard\DashboardFieldConfigSourceOverrides>
     */
    public array $overrides;

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

