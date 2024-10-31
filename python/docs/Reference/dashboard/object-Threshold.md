---
title: <span class="badge object-type-class"></span> Threshold
---
# <span class="badge object-type-class"></span> Threshold

User-defined value for a metric that triggers visual changes in a panel when this value is met or exceeded

They are used to conditionally style and color visualizations based on query results , and can be applied to most visualizations.

## Definition

```python
class Threshold:
    """
    User-defined value for a metric that triggers visual changes in a panel when this value is met or exceeded
    They are used to conditionally style and color visualizations based on query results , and can be applied to most visualizations.
    """

    # Value represents a specified metric for the threshold, which triggers a visual change in the dashboard when this value is met or exceeded.
    # Nulls currently appear here when serializing -Infinity to JSON.
    value: typing.Optional[float]
    # Color represents the color of the visual change that will occur in the dashboard when the threshold value is met or exceeded.
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

