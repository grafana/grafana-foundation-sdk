---
title: <span class="badge object-type-struct"></span> TableActionsCellOptions
---
# <span class="badge object-type-struct"></span> TableActionsCellOptions

Show actions in the cell

## Definition

```go
type TableActionsCellOptions struct {
    Type string `json:"type"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TableActionsCellOptions` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (tableActionsCellOptions *TableActionsCellOptions) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `TableActionsCellOptions` objects.

```go
func (tableActionsCellOptions *TableActionsCellOptions) Equals(other TableActionsCellOptions) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `TableActionsCellOptions` fields for violations and returns them.

```go
func (tableActionsCellOptions *TableActionsCellOptions) Validate() error
```

