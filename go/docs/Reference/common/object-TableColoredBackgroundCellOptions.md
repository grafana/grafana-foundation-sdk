---
title: <span class="badge object-type-struct"></span> TableColoredBackgroundCellOptions
---
# <span class="badge object-type-struct"></span> TableColoredBackgroundCellOptions

Colored background cell options

## Definition

```go
type TableColoredBackgroundCellOptions struct {
    Type string `json:"type"`
    Mode *common.TableCellBackgroundDisplayMode `json:"mode,omitempty"`
    ApplyToRow *bool `json:"applyToRow,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TableColoredBackgroundCellOptions` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (tableColoredBackgroundCellOptions *TableColoredBackgroundCellOptions) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `TableColoredBackgroundCellOptions` objects.

```go
func (tableColoredBackgroundCellOptions *TableColoredBackgroundCellOptions) Equals(other TableColoredBackgroundCellOptions) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `TableColoredBackgroundCellOptions` fields for violations and returns them.

```go
func (tableColoredBackgroundCellOptions *TableColoredBackgroundCellOptions) Validate() error
```

## See also

 * <span class="badge builder"></span> [TableColoredBackgroundCellOptionsBuilder](./builder-TableColoredBackgroundCellOptionsBuilder.md)
