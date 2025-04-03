---
title: <span class="badge object-type-class"></span> NodeOptions
---
# <span class="badge object-type-class"></span> NodeOptions

## Definition

```python
class NodeOptions:
    # Unit for the main stat to override what ever is set in the data frame.
    main_stat_unit: typing.Optional[str]
    # Unit for the secondary stat to override what ever is set in the data frame.
    secondary_stat_unit: typing.Optional[str]
    # Define which fields are shown as part of the node arc (colored circle around the node).
    arcs: typing.Optional[list[nodegraph.ArcOption]]
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

 * <span class="badge builder"></span> [NodeOptions](./builder-NodeOptions.md)
