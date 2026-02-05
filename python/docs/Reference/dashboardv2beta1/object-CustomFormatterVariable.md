---
title: <span class="badge object-type-class"></span> CustomFormatterVariable
---
# <span class="badge object-type-class"></span> CustomFormatterVariable

Custom formatter variable

## Definition

```python
class CustomFormatterVariable:
    """
    Custom formatter variable
    """

    name: str
    type_val: dashboardv2beta1.VariableType
    multi: bool
    include_all: bool
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

 * <span class="badge builder"></span> [CustomFormatterVariable](./builder-CustomFormatterVariable.md)
