---
title: <span class="badge object-type-class"></span> TypeMathOrTypeReduceOrTypeResampleOrTypeClassicConditionsOrTypeThresholdOrTypeSql
---
# <span class="badge object-type-class"></span> TypeMathOrTypeReduceOrTypeResampleOrTypeClassicConditionsOrTypeThresholdOrTypeSql

## Definition

```php
class TypeMathOrTypeReduceOrTypeResampleOrTypeClassicConditionsOrTypeThresholdOrTypeSql implements \JsonSerializable
{
    public ?\Grafana\Foundation\Expr\TypeMath $typeMath;

    public ?\Grafana\Foundation\Expr\TypeReduce $typeReduce;

    public ?\Grafana\Foundation\Expr\TypeResample $typeResample;

    public ?\Grafana\Foundation\Expr\TypeClassicConditions $typeClassicConditions;

    public ?\Grafana\Foundation\Expr\TypeThreshold $typeThreshold;

    public ?\Grafana\Foundation\Expr\TypeSql $typeSql;

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

