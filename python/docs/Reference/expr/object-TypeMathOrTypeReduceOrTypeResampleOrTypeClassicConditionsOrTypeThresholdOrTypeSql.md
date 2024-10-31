---
title: <span class="badge object-type-class"></span> TypeMathOrTypeReduceOrTypeResampleOrTypeClassicConditionsOrTypeThresholdOrTypeSql
---
# <span class="badge object-type-class"></span> TypeMathOrTypeReduceOrTypeResampleOrTypeClassicConditionsOrTypeThresholdOrTypeSql

## Definition

```python
class TypeMathOrTypeReduceOrTypeResampleOrTypeClassicConditionsOrTypeThresholdOrTypeSql:
    type_math: typing.Optional[expr.TypeMath]
    type_reduce: typing.Optional[expr.TypeReduce]
    type_resample: typing.Optional[expr.TypeResample]
    type_classic_conditions: typing.Optional[expr.TypeClassicConditions]
    type_threshold: typing.Optional[expr.TypeThreshold]
    type_sql: typing.Optional[expr.TypeSql]
```
## Methods

### <span class="badge object-method"></span> to_json

Converts this object into a representation that can easily be encoded to JSON.

```python
def to_json() -> dict[str, object]
```

### <span class="badge object-method"></span> from_json

Builds this object from a JSON-decoded dict.

```python
@classmethod
def from_json(data: dict[str, typing.Any]) -> typing.Self
```

