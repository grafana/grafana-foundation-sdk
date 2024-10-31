---
title: <span class="badge object-type-class"></span> TableImageCellOptions
---
# <span class="badge object-type-class"></span> TableImageCellOptions

Json view cell options

## Definition

```python
class TableImageCellOptions:
    """
    Json view cell options
    """

    type_val: typing.Literal["image"]
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

