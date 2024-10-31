---
title: <span class="badge builder"></span> RawDocumentBuilder
---
# <span class="badge builder"></span> RawDocumentBuilder

## Constructor

```go
func NewRawDocumentBuilder() *RawDocumentBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *RawDocumentBuilder) Build() (RawDocument, error)
```

### <span class="badge object-method"></span> Hide

```go
func (builder *RawDocumentBuilder) Hide(hide bool) *RawDocumentBuilder
```

### <span class="badge object-method"></span> Id

```go
func (builder *RawDocumentBuilder) Id(id string) *RawDocumentBuilder
```

### <span class="badge object-method"></span> Settings

```go
func (builder *RawDocumentBuilder) Settings(settings cog.Builder[elasticsearch.ElasticsearchRawDocumentSettings]) *RawDocumentBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [RawDocument](./object-RawDocument.md)
