---
title: <span class="badge object-type-class"></span> Threshold
---
# <span class="badge object-type-class"></span> Threshold

## Definition

```python
class Threshold:
    # Value null means -Infinity
    value: typing.Optional[float]
    # Optional dashboard-variable expression (e.g. `$myVar`) resolved at render time; `value` is the numeric fallback when the expression cannot be resolved to a single finite number.
    value_expr: typing.Optional[str]
    color: str
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

