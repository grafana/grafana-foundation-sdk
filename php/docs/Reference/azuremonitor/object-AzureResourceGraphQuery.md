---
title: <span class="badge object-type-class"></span> AzureResourceGraphQuery
---
# <span class="badge object-type-class"></span> AzureResourceGraphQuery

## Definition

```php
class AzureResourceGraphQuery implements \JsonSerializable
{
    /**
     * Azure Resource Graph KQL query to be executed.
     */
    public ?string $query;

    /**
     * Specifies the format results should be returned as. Defaults to table.
     */
    public ?string $resultFormat;

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

 * <span class="badge builder"></span> [AzureResourceGraphQueryBuilder](./builder-AzureResourceGraphQueryBuilder.md)
