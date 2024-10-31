---
title: <span class="badge object-type-struct"></span> TableColorTextCellOptions
---
# <span class="badge object-type-struct"></span> TableColorTextCellOptions

Colored text cell options

## Definition

```go
type TableColorTextCellOptions struct {
    Type string `json:"type"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TableColorTextCellOptions` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (tableColorTextCellOptions *TableColorTextCellOptions) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `TableColorTextCellOptions` objects.

```go
func (tableColorTextCellOptions *TableColorTextCellOptions) Equals(other TableColorTextCellOptions) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `TableColorTextCellOptions` fields for violations and returns them.

```go
func (tableColorTextCellOptions *TableColorTextCellOptions) Validate() error
```

