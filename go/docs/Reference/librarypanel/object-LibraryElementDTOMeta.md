---
title: <span class="badge object-type-struct"></span> LibraryElementDTOMeta
---
# <span class="badge object-type-struct"></span> LibraryElementDTOMeta

## Definition

```go
type LibraryElementDTOMeta struct {
    FolderName string `json:"folderName"`
    FolderUid string `json:"folderUid"`
    ConnectedDashboards int64 `json:"connectedDashboards"`
    Created time.Time `json:"created"`
    Updated time.Time `json:"updated"`
    CreatedBy librarypanel.LibraryElementDTOMetaUser `json:"createdBy"`
    UpdatedBy librarypanel.LibraryElementDTOMetaUser `json:"updatedBy"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `LibraryElementDTOMeta` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (libraryElementDTOMeta *LibraryElementDTOMeta) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `LibraryElementDTOMeta` objects.

```go
func (libraryElementDTOMeta *LibraryElementDTOMeta) Equals(other LibraryElementDTOMeta) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `LibraryElementDTOMeta` fields for violations and returns them.

```go
func (libraryElementDTOMeta *LibraryElementDTOMeta) Validate() error
```

## See also

 * <span class="badge builder"></span> [LibraryElementDTOMetaBuilder](./builder-LibraryElementDTOMetaBuilder.md)
