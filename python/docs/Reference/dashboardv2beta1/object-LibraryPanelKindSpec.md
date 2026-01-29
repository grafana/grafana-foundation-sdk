---
title: <span class="badge object-type-class"></span> LibraryPanelKindSpec
---
# <span class="badge object-type-class"></span> LibraryPanelKindSpec

## Definition

```python
class LibraryPanelKindSpec:
    # Panel ID for the library panel in the dashboard
    id_val: float
    # Title for the library panel in the dashboard
    title: str
    library_panel: dashboardv2beta1.LibraryPanelRef
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

 * <span class="badge builder"></span> [LibraryPanelKindSpec](./builder-LibraryPanelKindSpec.md)
