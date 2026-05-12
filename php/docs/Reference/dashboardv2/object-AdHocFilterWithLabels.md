---
title: <span class="badge object-type-class"></span> AdHocFilterWithLabels
---
# <span class="badge object-type-class"></span> AdHocFilterWithLabels

Define the AdHocFilterWithLabels type

## Definition

```php
class AdHocFilterWithLabels implements \JsonSerializable
{
    public string $key;

    public string $operator;

    public string $value;

    /**
     * @var array<string>|null
     */
    public ?array $values;

    public ?string $keyLabel;

    /**
     * @var array<string>|null
     */
    public ?array $valueLabels;

    public ?bool $forceEdit;

    public ?string $origin;

    /**
     * @deprecated
     */
    public ?string $condition;

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

 * <span class="badge builder"></span> [AdHocFilterWithLabelsBuilder](./builder-AdHocFilterWithLabelsBuilder.md)
