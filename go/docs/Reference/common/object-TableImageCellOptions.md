---
title: <span class="badge object-type-struct"></span> TableImageCellOptions
---
# <span class="badge object-type-struct"></span> TableImageCellOptions

Json view cell options

## Definition

```go
type TableImageCellOptions struct {
    Type string `json:"type"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TableImageCellOptions` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (tableImageCellOptions *TableImageCellOptions) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `TableImageCellOptions` objects.

```go
func (tableImageCellOptions *TableImageCellOptions) Equals(other TableImageCellOptions) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `TableImageCellOptions` fields for violations and returns them.

```go
func (tableImageCellOptions *TableImageCellOptions) Validate() error
```

