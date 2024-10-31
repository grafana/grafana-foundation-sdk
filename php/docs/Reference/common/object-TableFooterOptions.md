---
title: <span class="badge object-type-class"></span> TableFooterOptions
---
# <span class="badge object-type-class"></span> TableFooterOptions

Footer options

## Definition

```php
class TableFooterOptions implements \JsonSerializable
{
    public bool $show;

    /**
     * actually 1 value
     * @var array<string>
     */
    public array $reducer;

    /**
     * @var array<string>|null
     */
    public ?array $fields;

    public ?bool $enablePagination;

    public ?bool $countRows;

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

 * <span class="badge builder"></span> [TableFooterOptionsBuilder](./builder-TableFooterOptionsBuilder.md)
