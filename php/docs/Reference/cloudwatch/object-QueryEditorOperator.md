---
title: <span class="badge object-type-class"></span> QueryEditorOperator
---
# <span class="badge object-type-class"></span> QueryEditorOperator

TS type is QueryEditorOperator<T extends QueryEditorOperatorValueType>, extended in veneer

## Definition

```php
class QueryEditorOperator implements \JsonSerializable
{
    public ?string $name;

    /**
     * @var string|bool|int|array<string|bool|int>|null
     */
    public $value;

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

 * <span class="badge builder"></span> [QueryEditorOperatorBuilder](./builder-QueryEditorOperatorBuilder.md)
