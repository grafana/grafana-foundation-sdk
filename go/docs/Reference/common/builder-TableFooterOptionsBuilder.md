---
title: <span class="badge builder"></span> TableFooterOptionsBuilder
---
# <span class="badge builder"></span> TableFooterOptionsBuilder

## Constructor

```go
func NewTableFooterOptionsBuilder() *TableFooterOptionsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *TableFooterOptionsBuilder) Build() (TableFooterOptions, error)
```

### <span class="badge object-method"></span> CountRows

```go
func (builder *TableFooterOptionsBuilder) CountRows(countRows bool) *TableFooterOptionsBuilder
```

### <span class="badge object-method"></span> EnablePagination

```go
func (builder *TableFooterOptionsBuilder) EnablePagination(enablePagination bool) *TableFooterOptionsBuilder
```

### <span class="badge object-method"></span> Fields

```go
func (builder *TableFooterOptionsBuilder) Fields(fields []string) *TableFooterOptionsBuilder
```

### <span class="badge object-method"></span> Reducer

actually 1 value

```go
func (builder *TableFooterOptionsBuilder) Reducer(reducer []string) *TableFooterOptionsBuilder
```

### <span class="badge object-method"></span> Show

```go
func (builder *TableFooterOptionsBuilder) Show(show bool) *TableFooterOptionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [TableFooterOptions](./object-TableFooterOptions.md)
