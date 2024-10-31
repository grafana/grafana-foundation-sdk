---
title: <span class="badge object-type-class"></span> ValueMappingResult
---
# <span class="badge object-type-class"></span> ValueMappingResult

Result used as replacement with text and color when the value matches

## Definition

```python
class ValueMappingResult:
    """
    Result used as replacement with text and color when the value matches
    """

    # Text to display when the value matches
    text: typing.Optional[str]
    # Text to use when the value matches
    color: typing.Optional[str]
    # Icon to display when the value matches. Only specific visualizations.
    icon: typing.Optional[str]
    # Position in the mapping array. Only used internally.
    index: typing.Optional[int]
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

