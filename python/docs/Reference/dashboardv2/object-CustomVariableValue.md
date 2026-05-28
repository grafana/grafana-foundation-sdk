---
title: <span class="badge object-type-class"></span> CustomVariableValue
---
# <span class="badge object-type-class"></span> CustomVariableValue

Custom variable value

## Definition

```python
class CustomVariableValue:
    """
    Custom variable value
    """

    # The format name or function used in the expression
    formatter: typing.Optional[str]
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

