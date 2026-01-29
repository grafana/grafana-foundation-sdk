---
title: <span class="badge builder"></span> KindBuilder
---
# <span class="badge builder"></span> KindBuilder

## Constructor

```go
func NewKindBuilder() *KindBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *KindBuilder) Build() (Kind, error)
```

### <span class="badge object-method"></span> Kind

```go
func (builder *KindBuilder) Kind(kind string) *KindBuilder
```

### <span class="badge object-method"></span> Metadata

```go
func (builder *KindBuilder) Metadata(metadata any) *KindBuilder
```

### <span class="badge object-method"></span> Spec

```go
func (builder *KindBuilder) Spec(spec any) *KindBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Kind](./object-Kind.md)
