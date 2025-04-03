---
title: <span class="badge object-type-struct"></span> TableSortByFieldState
---
# <span class="badge object-type-struct"></span> TableSortByFieldState

Sort by field state

## Definition

```go
type TableSortByFieldState struct {
    // Sets the display name of the field to sort by
    DisplayName string `json:"displayName"`
    // Flag used to indicate descending sort order
    Desc *bool `json:"desc,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TableSortByFieldState` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (tableSortByFieldState *TableSortByFieldState) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `TableSortByFieldState` objects.

```go
func (tableSortByFieldState *TableSortByFieldState) Equals(other TableSortByFieldState) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `TableSortByFieldState` fields for violations and returns them.

```go
func (tableSortByFieldState *TableSortByFieldState) Validate() error
```

## See also

 * <span class="badge builder"></span> [TableSortByFieldStateBuilder](./builder-TableSortByFieldStateBuilder.md)
