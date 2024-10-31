---
title: <span class="badge object-type-class"></span> ExprTypeThresholdConditionsUnloadEvaluator
---
# <span class="badge object-type-class"></span> ExprTypeThresholdConditionsUnloadEvaluator

## Definition

```php
class ExprTypeThresholdConditionsUnloadEvaluator implements \JsonSerializable
{
    /**
     * @var array<float>
     */
    public array $params;

    /**
     * e.g. "gt"
     */
    public \Grafana\Foundation\Expr\TypeThresholdType $type;

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

 * <span class="badge builder"></span> [ExprTypeThresholdConditionsUnloadEvaluatorBuilder](./builder-ExprTypeThresholdConditionsUnloadEvaluatorBuilder.md)
