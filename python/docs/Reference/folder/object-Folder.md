---
title: <span class="badge object-type-class"></span> Folder
---
# <span class="badge object-type-class"></span> Folder

TODO:

common metadata will soon support setting the parent folder in the metadata

## Definition

```python
class Folder:
    """
    TODO:
    common metadata will soon support setting the parent folder in the metadata
    """

    # Unique folder id. (will be k8s name)
    uid: str
    # Folder title
    title: str
    # Description of the folder.
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

 * <span class="badge builder"></span> [Folder](./builder-Folder.md)
