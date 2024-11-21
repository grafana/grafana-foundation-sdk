---
title: <span class="badge builder"></span> LibraryPanel
---
# <span class="badge builder"></span> LibraryPanel

## Constructor

```python
LibraryPanel()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> librarypanel.LibraryPanel
```

### <span class="badge object-method"></span> description

Panel description

```python
def description(description: str) -> typing.Self
```

### <span class="badge object-method"></span> folder_uid

Folder UID

```python
def folder_uid(folder_uid: str) -> typing.Self
```

### <span class="badge object-method"></span> meta

Object storage metadata

```python
def meta(meta: cogbuilder.Builder[librarypanel.LibraryElementDTOMeta]) -> typing.Self
```

### <span class="badge object-method"></span> model

TODO: should be the same panel schema defined in dashboard

Typescript: Omit<Panel, 'gridPos' | 'id' | 'libraryPanel'>;

```python
def model(model: cogbuilder.Builder[librarypanel.PanelModel]) -> typing.Self
```

### <span class="badge object-method"></span> name

Panel name (also saved in the model)

```python
def name(name: str) -> typing.Self
```

### <span class="badge object-method"></span> schema_version

Dashboard version when this was saved (zero if unknown)

```python
def schema_version(schema_version: int) -> typing.Self
```

### <span class="badge object-method"></span> type_val

The panel type (from inside the model)

```python
def type_val(type_val: str) -> typing.Self
```

### <span class="badge object-method"></span> uid

Library element UID

```python
def uid(uid: str) -> typing.Self
```

### <span class="badge object-method"></span> version

panel version, incremented each time the dashboard is updated.

```python
def version(version: int) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [LibraryPanel](./object-LibraryPanel.md)
