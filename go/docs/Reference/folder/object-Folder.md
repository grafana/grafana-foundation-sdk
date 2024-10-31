---
title: <span class="badge object-type-struct"></span> Folder
---
# <span class="badge object-type-struct"></span> Folder

TODO:

common metadata will soon support setting the parent folder in the metadata

## Definition

```go
type Folder struct {
    // Unique folder id. (will be k8s name)
    Uid string `json:"uid"`
    // Folder title
    Title string `json:"title"`
    // Description of the folder.
    Description *string `json:"description,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Folder` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (folder *Folder) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `Folder` objects.

```go
func (folder *Folder) Equals(other Folder) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `Folder` fields for violations and returns them.

```go
func (folder *Folder) Validate() error
```

## See also

 * <span class="badge builder"></span> [FolderBuilder](./builder-FolderBuilder.md)
