---
title: <span class="badge object-type-class"></span> TimeSeriesQuery
---
# <span class="badge object-type-class"></span> TimeSeriesQuery

Time Series sub-query properties.

## Definition

```php
class TimeSeriesQuery implements \JsonSerializable
{
    /**
     * GCP project to execute the query against.
     */
    public string $projectName;

    /**
     * MQL query to be executed.
     */
    public string $query;

    /**
     * To disable the graphPeriod, it should explictly be set to 'disabled'.
     */
    public ?string $graphPeriod;

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

 * <span class="badge builder"></span> [TimeSeriesQueryBuilder](./builder-TimeSeriesQueryBuilder.md)
