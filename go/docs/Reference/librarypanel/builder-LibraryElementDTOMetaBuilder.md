---
title: <span class="badge builder"></span> LibraryElementDTOMetaBuilder
---
# <span class="badge builder"></span> LibraryElementDTOMetaBuilder

## Constructor

```go
func NewLibraryElementDTOMetaBuilder() *LibraryElementDTOMetaBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *LibraryElementDTOMetaBuilder) Build() (LibraryElementDTOMeta, error)
```

### <span class="badge object-method"></span> ConnectedDashboards

```go
func (builder *LibraryElementDTOMetaBuilder) ConnectedDashboards(connectedDashboards int64) *LibraryElementDTOMetaBuilder
```

### <span class="badge object-method"></span> Created

```go
func (builder *LibraryElementDTOMetaBuilder) Created(created time.Time) *LibraryElementDTOMetaBuilder
```

### <span class="badge object-method"></span> CreatedBy

```go
func (builder *LibraryElementDTOMetaBuilder) CreatedBy(createdBy cog.Builder[librarypanel.LibraryElementDTOMetaUser]) *LibraryElementDTOMetaBuilder
```

### <span class="badge object-method"></span> FolderName

```go
func (builder *LibraryElementDTOMetaBuilder) FolderName(folderName string) *LibraryElementDTOMetaBuilder
```

### <span class="badge object-method"></span> FolderUid

```go
func (builder *LibraryElementDTOMetaBuilder) FolderUid(folderUid string) *LibraryElementDTOMetaBuilder
```

### <span class="badge object-method"></span> Updated

```go
func (builder *LibraryElementDTOMetaBuilder) Updated(updated time.Time) *LibraryElementDTOMetaBuilder
```

### <span class="badge object-method"></span> UpdatedBy

```go
func (builder *LibraryElementDTOMetaBuilder) UpdatedBy(updatedBy cog.Builder[librarypanel.LibraryElementDTOMetaUser]) *LibraryElementDTOMetaBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [LibraryElementDTOMeta](./object-LibraryElementDTOMeta.md)
