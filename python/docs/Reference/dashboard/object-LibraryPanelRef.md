---
title: <span class="badge object-type-class"></span> LibraryPanelRef
---
# <span class="badge object-type-class"></span> LibraryPanelRef

A library panel is a reusable panel that you can use in any dashboard.

When you make a change to a library panel, that change propagates to all instances of where the panel is used.

Library panels streamline reuse of panels across multiple dashboards.

## Definition

```python
class LibraryPanelRef:
    """
    A library panel is a reusable panel that you can use in any dashboard.
    When you make a change to a library panel, that change propagates to all instances of where the panel is used.
    Library panels streamline reuse of panels across multiple dashboards.
    """

    # Library panel name
    name: str
    # Library panel uid
    uid: str
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

