---
title: <span class="badge builder"></span> TableSortByFieldStateBuilder
---
# <span class="badge builder"></span> TableSortByFieldStateBuilder

## Constructor

```go
func NewTableSortByFieldStateBuilder() *TableSortByFieldStateBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *TableSortByFieldStateBuilder) Build() (TableSortByFieldState, error)
```

### <span class="badge object-method"></span> Desc

Flag used to indicate descending sort order

```go
func (builder *TableSortByFieldStateBuilder) Desc(desc bool) *TableSortByFieldStateBuilder
```

### <span class="badge object-method"></span> DisplayName

Sets the display name of the field to sort by

```go
func (builder *TableSortByFieldStateBuilder) DisplayName(displayName string) *TableSortByFieldStateBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [TableSortByFieldState](./object-TableSortByFieldState.md)
