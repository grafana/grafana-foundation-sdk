---
title: <span class="badge object-type-class"></span> Scenario
---
# <span class="badge object-type-class"></span> Scenario

TODO: Should this live here given it's not used in the dataquery?

## Definition

```python
class Scenario:
    """
    TODO: Should this live here given it's not used in the dataquery?
    """

    id_val: str
    name: str
    string_input: str
    description: typing.Optional[str]
    hide_alias_field: typing.Optional[bool]
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

 * <span class="badge builder"></span> [Scenario](./builder-Scenario.md)
