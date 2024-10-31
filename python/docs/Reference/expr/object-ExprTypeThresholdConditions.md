---
title: <span class="badge object-type-class"></span> ExprTypeThresholdConditions
---
# <span class="badge object-type-class"></span> ExprTypeThresholdConditions

## Definition

```python
class ExprTypeThresholdConditions:
    evaluator: expr.ExprTypeThresholdConditionsEvaluator
    loaded_dimensions: typing.Optional[object]
    unload_evaluator: typing.Optional[expr.ExprTypeThresholdConditionsUnloadEvaluator]
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

## See also

 * <span class="badge builder"></span> [ExprTypeThresholdConditions](./builder-ExprTypeThresholdConditions.md)
