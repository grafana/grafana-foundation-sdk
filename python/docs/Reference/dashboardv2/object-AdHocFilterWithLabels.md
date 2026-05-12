---
title: <span class="badge object-type-class"></span> AdHocFilterWithLabels
---
# <span class="badge object-type-class"></span> AdHocFilterWithLabels

Define the AdHocFilterWithLabels type

## Definition

```python
class AdHocFilterWithLabels:
    """
    Define the AdHocFilterWithLabels type
    """

    key: str
    operator: str
    value: str
    values: typing.Optional[list[str]]
    key_label: typing.Optional[str]
    value_labels: typing.Optional[list[str]]
    force_edit: typing.Optional[bool]
    origin: str
    # @deprecated
    condition: typing.Optional[str]
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

 * <span class="badge builder"></span> [AdHocFilterWithLabels](./builder-AdHocFilterWithLabels.md)
