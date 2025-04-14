---
title: <span class="badge object-type-class"></span> BuilderQueryEditorGroupByExpression
---
# <span class="badge object-type-class"></span> BuilderQueryEditorGroupByExpression

## Definition

```php
class BuilderQueryEditorGroupByExpression implements \JsonSerializable
{
    public ?\Grafana\Foundation\Azuremonitor\BuilderQueryEditorProperty $property;

    public ?\Grafana\Foundation\Azuremonitor\BuilderQueryEditorProperty $interval;

    public ?bool $focus;

    public ?\Grafana\Foundation\Azuremonitor\BuilderQueryEditorExpressionType $type;

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

 * <span class="badge builder"></span> [BuilderQueryEditorGroupByExpressionBuilder](./builder-BuilderQueryEditorGroupByExpressionBuilder.md)
