---
title: <span class="badge object-type-struct"></span> TableDataLinksCellOptions
---
# <span class="badge object-type-struct"></span> TableDataLinksCellOptions

Show data links in the cell

## Definition

```go
type TableDataLinksCellOptions struct {
    Type string `json:"type"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TableDataLinksCellOptions` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (tableDataLinksCellOptions *TableDataLinksCellOptions) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `TableDataLinksCellOptions` objects.

```go
func (tableDataLinksCellOptions *TableDataLinksCellOptions) Equals(other TableDataLinksCellOptions) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `TableDataLinksCellOptions` fields for violations and returns them.

```go
func (tableDataLinksCellOptions *TableDataLinksCellOptions) Validate() error
```

