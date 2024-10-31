---
title: <span class="badge object-type-struct"></span> TableAutoCellOptions
---
# <span class="badge object-type-struct"></span> TableAutoCellOptions

Auto mode table cell options

## Definition

```go
type TableAutoCellOptions struct {
    Type string `json:"type"`
    WrapText *bool `json:"wrapText,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TableAutoCellOptions` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (tableAutoCellOptions *TableAutoCellOptions) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `TableAutoCellOptions` objects.

```go
func (tableAutoCellOptions *TableAutoCellOptions) Equals(other TableAutoCellOptions) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `TableAutoCellOptions` fields for violations and returns them.

```go
func (tableAutoCellOptions *TableAutoCellOptions) Validate() error
```

## See also

 * <span class="badge builder"></span> [TableAutoCellOptionsBuilder](./builder-TableAutoCellOptionsBuilder.md)
