---
title: <span class="badge object-type-class"></span> BuilderQueryEditorWhereExpression
---
# <span class="badge object-type-class"></span> BuilderQueryEditorWhereExpression

## Definition

```php
class BuilderQueryEditorWhereExpression implements \JsonSerializable
{
    public \Grafana\Foundation\Azuremonitor\BuilderQueryEditorExpressionType $type;

    /**
     * @var array<\Grafana\Foundation\Azuremonitor\BuilderQueryEditorWhereExpressionItems>
     */
    public array $expressions;

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

 * <span class="badge builder"></span> [BuilderQueryEditorWhereExpressionBuilder](./builder-BuilderQueryEditorWhereExpressionBuilder.md)
