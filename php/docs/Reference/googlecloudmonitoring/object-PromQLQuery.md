---
title: <span class="badge object-type-class"></span> PromQLQuery
---
# <span class="badge object-type-class"></span> PromQLQuery

PromQL sub-query properties.

## Definition

```php
class PromQLQuery implements \JsonSerializable
{
    /**
     * GCP project to execute the query against.
     */
    public string $projectName;

    /**
     * PromQL expression/query to be executed.
     */
    public string $expr;

    /**
     * PromQL min step
     */
    public string $step;

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

 * <span class="badge builder"></span> [PromQLQueryBuilder](./builder-PromQLQueryBuilder.md)
