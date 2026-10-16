---
title: <span class="badge builder"></span> CodeCellContentSpecBuilder
---
# <span class="badge builder"></span> CodeCellContentSpecBuilder

## Constructor

```go
func NewCodeCellContentSpecBuilder() *CodeCellContentSpecBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *CodeCellContentSpecBuilder) Build() (CodeCellContentSpec, error)
```

### <span class="badge object-method"></span> Annotation

```go
func (builder *CodeCellContentSpecBuilder) Annotation(annotation string) *CodeCellContentSpecBuilder
```

### <span class="badge object-method"></span> Code

```go
func (builder *CodeCellContentSpecBuilder) Code(code string) *CodeCellContentSpecBuilder
```

### <span class="badge object-method"></span> Highlight

```go
func (builder *CodeCellContentSpecBuilder) Highlight(highlight []int64) *CodeCellContentSpecBuilder
```

### <span class="badge object-method"></span> Language

```go
func (builder *CodeCellContentSpecBuilder) Language(language string) *CodeCellContentSpecBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [CodeCellContentSpec](./object-CodeCellContentSpec.md)
