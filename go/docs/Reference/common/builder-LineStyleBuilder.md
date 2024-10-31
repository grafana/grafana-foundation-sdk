---
title: <span class="badge builder"></span> LineStyleBuilder
---
# <span class="badge builder"></span> LineStyleBuilder

## Constructor

```go
func NewLineStyleBuilder() *LineStyleBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *LineStyleBuilder) Build() (LineStyle, error)
```

### <span class="badge object-method"></span> Dash

```go
func (builder *LineStyleBuilder) Dash(dash []float64) *LineStyleBuilder
```

### <span class="badge object-method"></span> Fill

```go
func (builder *LineStyleBuilder) Fill(fill common.LineStyleFill) *LineStyleBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [LineStyle](./object-LineStyle.md)
