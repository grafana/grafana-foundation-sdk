---
title: <span class="badge object-type-class"></span> VariableOption
---
# <span class="badge object-type-class"></span> VariableOption

Option to be selected in a variable.

## Definition

```python
class VariableOption:
    """
    Option to be selected in a variable.
    """

    # Whether the option is selected or not
    selected: typing.Optional[bool]
    # Text to be displayed for the option
    text: typing.Union[str, list[str]]
    # Value of the option
    value: typing.Union[str, list[str]]
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

