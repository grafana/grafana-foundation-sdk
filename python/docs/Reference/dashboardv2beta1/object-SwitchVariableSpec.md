---
title: <span class="badge object-type-class"></span> SwitchVariableSpec
---
# <span class="badge object-type-class"></span> SwitchVariableSpec

## Definition

```python
class SwitchVariableSpec:
    name: str
    current: str
    enabled_value: str
    disabled_value: str
    label: typing.Optional[str]
    hide: dashboardv2beta1.VariableHide
    skip_url_sync: bool
    description: typing.Optional[str]
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

 * <span class="badge builder"></span> [SwitchVariableSpec](./builder-SwitchVariableSpec.md)
