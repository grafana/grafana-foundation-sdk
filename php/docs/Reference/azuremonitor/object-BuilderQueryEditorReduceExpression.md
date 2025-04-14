---
title: <span class="badge object-type-class"></span> BuilderQueryEditorReduceExpression
---
# <span class="badge object-type-class"></span> BuilderQueryEditorReduceExpression

## Definition

```php
class BuilderQueryEditorReduceExpression implements \JsonSerializable
{
    public ?\Grafana\Foundation\Azuremonitor\BuilderQueryEditorProperty $property;

    public ?\Grafana\Foundation\Azuremonitor\BuilderQueryEditorProperty $reduce;

    /**
     * @var array<\Grafana\Foundation\Azuremonitor\BuilderQueryEditorFunctionParameterExpression>|null
     */
    public ?array $parameters;

    public ?bool $focus;

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

 * <span class="badge builder"></span> [BuilderQueryEditorReduceExpressionBuilder](./builder-BuilderQueryEditorReduceExpressionBuilder.md)
