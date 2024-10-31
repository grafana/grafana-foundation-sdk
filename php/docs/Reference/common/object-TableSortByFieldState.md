---
title: <span class="badge object-type-class"></span> TableSortByFieldState
---
# <span class="badge object-type-class"></span> TableSortByFieldState

Sort by field state

## Definition

```php
class TableSortByFieldState implements \JsonSerializable
{
    /**
     * Sets the display name of the field to sort by
     */
    public string $displayName;

    /**
     * Flag used to indicate descending sort order
     */
    public ?bool $desc;

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

 * <span class="badge builder"></span> [TableSortByFieldStateBuilder](./builder-TableSortByFieldStateBuilder.md)
