---
title: <span class="badge object-type-struct"></span> LibraryPanel
---
# <span class="badge object-type-struct"></span> LibraryPanel

## Definition

```go
type LibraryPanel struct {
    // Folder UID
    FolderUid *string `json:"folderUid,omitempty"`
    // Library element UID
    Uid string `json:"uid"`
    // Panel name (also saved in the model)
    Name string `json:"name"`
    // Panel description
    Description *string `json:"description,omitempty"`
    // The panel type (from inside the model)
    Type string `json:"type"`
    // Dashboard version when this was saved (zero if unknown)
    SchemaVersion *uint16 `json:"schemaVersion,omitempty"`
    // panel version, incremented each time the dashboard is updated.
    Version int64 `json:"version"`
    // TODO: should be the same panel schema defined in dashboard
    // Typescript: Omit<Panel, 'gridPos' | 'id' | 'libraryPanel'>;
    Model librarypanel.LibrarypanelLibraryPanelModel `json:"model"`
    // Object storage metadata
    Meta *librarypanel.LibraryElementDTOMeta `json:"meta,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `LibraryPanel` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (libraryPanel *LibraryPanel) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `LibraryPanel` objects.

```go
func (libraryPanel *LibraryPanel) Equals(other LibraryPanel) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `LibraryPanel` fields for violations and returns them.

```go
func (libraryPanel *LibraryPanel) Validate() error
```

## See also

 * <span class="badge builder"></span> [LibraryPanelBuilder](./builder-LibraryPanelBuilder.md)
