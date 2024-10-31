---
title: <span class="badge object-type-class"></span> AzureTracesFilter
---
# <span class="badge object-type-class"></span> AzureTracesFilter

## Definition

```php
class AzureTracesFilter implements \JsonSerializable
{
    /**
     * Property name, auto-populated based on available traces.
     */
    public string $property;

    /**
     * Comparison operator to use. Either equals or not equals.
     */
    public string $operation;

    /**
     * Values to filter by.
     * @var array<string>
     */
    public array $filters;

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

 * <span class="badge builder"></span> [AzureTracesFilterBuilder](./builder-AzureTracesFilterBuilder.md)
