---
title: <span class="badge object-type-class"></span> ExprTypeThresholdConditionsUnloadEvaluator
---
# <span class="badge object-type-class"></span> ExprTypeThresholdConditionsUnloadEvaluator

## Definition

```python
class ExprTypeThresholdConditionsUnloadEvaluator:
    params: list[float]
    # e.g. "gt"
    type_val: typing.Literal["gt", "lt", "within_range", "outside_range"]
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

 * <span class="badge builder"></span> [ExprTypeThresholdConditionsUnloadEvaluator](./builder-ExprTypeThresholdConditionsUnloadEvaluator.md)
