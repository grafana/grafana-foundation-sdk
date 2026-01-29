---
title: <span class="badge builder"></span> DataLinkBuilder
---
# <span class="badge builder"></span> DataLinkBuilder

## Constructor

```go
func NewDataLinkBuilder() *DataLinkBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *DataLinkBuilder) Build() (DataLink, error)
```

### <span class="badge object-method"></span> TargetBlank

```go
func (builder *DataLinkBuilder) TargetBlank(targetBlank bool) *DataLinkBuilder
```

### <span class="badge object-method"></span> Title

```go
func (builder *DataLinkBuilder) Title(title string) *DataLinkBuilder
```

### <span class="badge object-method"></span> Url

```go
func (builder *DataLinkBuilder) Url(url string) *DataLinkBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [DataLink](./object-DataLink.md)
