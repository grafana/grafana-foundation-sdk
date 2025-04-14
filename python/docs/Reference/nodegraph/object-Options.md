---
title: <span class="badge object-type-class"></span> Options
---
# <span class="badge object-type-class"></span> Options

## Definition

```python
class Options:
    nodes: typing.Optional[nodegraph.NodeOptions]
    edges: typing.Optional[nodegraph.EdgeOptions]
    # How to handle zoom/scroll events in the node graph
    zoom_mode: typing.Optional[nodegraph.ZoomMode]
    # How to layout the nodes in the node graph
    layout_algorithm: typing.Optional[nodegraph.LayoutAlgorithm]
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

