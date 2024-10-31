---
title: <span class="badge object-type-class"></span> USAQuery
---
# <span class="badge object-type-class"></span> USAQuery

## Definition

```php
class USAQuery implements \JsonSerializable
{
    public ?string $mode;

    public ?string $period;

    /**
     * @var array<string>|null
     */
    public ?array $fields;

    /**
     * @var array<string>|null
     */
    public ?array $states;

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

 * <span class="badge builder"></span> [USAQueryBuilder](./builder-USAQueryBuilder.md)
