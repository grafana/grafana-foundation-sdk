---
title: <span class="badge object-type-struct"></span> Folder
---
# <span class="badge object-type-struct"></span> Folder

## Definition

```go
type Folder struct {
    Title string `json:"title"`
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
