---
title: <span class="badge builder"></span> LibraryPanelBuilder
---
# <span class="badge builder"></span> LibraryPanelBuilder

## Constructor

```go
func NewLibraryPanelBuilder() *LibraryPanelBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *LibraryPanelBuilder) Build() (LibraryPanel, error)
```

### <span class="badge object-method"></span> Description

Panel description

```go
func (builder *LibraryPanelBuilder) Description(description string) *LibraryPanelBuilder
```

### <span class="badge object-method"></span> FolderUid

Folder UID

```go
func (builder *LibraryPanelBuilder) FolderUid(folderUid string) *LibraryPanelBuilder
```

### <span class="badge object-method"></span> Meta

Object storage metadata

```go
func (builder *LibraryPanelBuilder) Meta(meta cog.Builder[librarypanel.LibraryElementDTOMeta]) *LibraryPanelBuilder
```

### <span class="badge object-method"></span> Model

TODO: should be the same panel schema defined in dashboard

Typescript: Omit<Panel, 'gridPos' | 'id' | 'libraryPanel'>;

```go
func (builder *LibraryPanelBuilder) Model(model cog.Builder[librarypanel.LibrarypanelLibraryPanelModel]) *LibraryPanelBuilder
```

### <span class="badge object-method"></span> Name

Panel name (also saved in the model)

```go
func (builder *LibraryPanelBuilder) Name(name string) *LibraryPanelBuilder
```

### <span class="badge object-method"></span> SchemaVersion

Dashboard version when this was saved (zero if unknown)

```go
func (builder *LibraryPanelBuilder) SchemaVersion(schemaVersion uint16) *LibraryPanelBuilder
```

### <span class="badge object-method"></span> Type

The panel type (from inside the model)

```go
func (builder *LibraryPanelBuilder) Type(typeArg string) *LibraryPanelBuilder
```

### <span class="badge object-method"></span> Uid

Library element UID

```go
func (builder *LibraryPanelBuilder) Uid(uid string) *LibraryPanelBuilder
```

### <span class="badge object-method"></span> Version

panel version, incremented each time the dashboard is updated.

```go
func (builder *LibraryPanelBuilder) Version(version int64) *LibraryPanelBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [LibraryPanel](./object-LibraryPanel.md)
