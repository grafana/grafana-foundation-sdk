---
title: <span class="badge object-type-struct"></span> LibraryElementDTOMetaUser
---
# <span class="badge object-type-struct"></span> LibraryElementDTOMetaUser

## Definition

```go
type LibraryElementDTOMetaUser struct {
    Id int64 `json:"id"`
    Name string `json:"name"`
    AvatarUrl string `json:"avatarUrl"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `LibraryElementDTOMetaUser` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (libraryElementDTOMetaUser *LibraryElementDTOMetaUser) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `LibraryElementDTOMetaUser` objects.

```go
func (libraryElementDTOMetaUser *LibraryElementDTOMetaUser) Equals(other LibraryElementDTOMetaUser) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `LibraryElementDTOMetaUser` fields for violations and returns them.

```go
func (libraryElementDTOMetaUser *LibraryElementDTOMetaUser) Validate() error
```

## See also

 * <span class="badge builder"></span> [LibraryElementDTOMetaUserBuilder](./builder-LibraryElementDTOMetaUserBuilder.md)
