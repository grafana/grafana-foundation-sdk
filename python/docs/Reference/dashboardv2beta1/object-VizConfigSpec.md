---
title: <span class="badge object-type-class"></span> VizConfigSpec
---
# <span class="badge object-type-class"></span> VizConfigSpec

--- Kinds ---

## Definition

```python
class VizConfigSpec:
    """
    --- Kinds ---
    """

    options: typing.Optional[object]
    field_config: dashboardv2beta1.FieldConfigSource
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

