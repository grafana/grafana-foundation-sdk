---
title: <span class="badge object-type-class"></span> LibraryPanel
---
# <span class="badge object-type-class"></span> LibraryPanel

## Definition

```python
class LibraryPanel:
    # Folder UID
    folder_uid: typing.Optional[str]
    # Library element UID
    uid: str
    # Panel name (also saved in the model)
    name: str
    # Panel description
    description: typing.Optional[str]
    # The panel type (from inside the model)
    type_val: str
    # Dashboard version when this was saved (zero if unknown)
    schema_version: typing.Optional[int]
    # panel version, incremented each time the dashboard is updated.
    version: int
    # TODO: should be the same panel schema defined in dashboard
    # Typescript: Omit<Panel, 'gridPos' | 'id' | 'libraryPanel'>;
    model: librarypanel.PanelModel
    # Object storage metadata
    meta: typing.Optional[librarypanel.LibraryElementDTOMeta]
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

 * <span class="badge builder"></span> [LibraryPanel](./builder-LibraryPanel.md)
