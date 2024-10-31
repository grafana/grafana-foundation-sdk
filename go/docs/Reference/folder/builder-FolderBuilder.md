---
title: <span class="badge builder"></span> FolderBuilder
---
# <span class="badge builder"></span> FolderBuilder

## Constructor

```go
func NewFolderBuilder() *FolderBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *FolderBuilder) Build() (Folder, error)
```

### <span class="badge object-method"></span> Description

Description of the folder.

```go
func (builder *FolderBuilder) Description(description string) *FolderBuilder
```

### <span class="badge object-method"></span> Title

Folder title

```go
func (builder *FolderBuilder) Title(title string) *FolderBuilder
```

### <span class="badge object-method"></span> Uid

Unique folder id. (will be k8s name)

```go
func (builder *FolderBuilder) Uid(uid string) *FolderBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Folder](./object-Folder.md)
