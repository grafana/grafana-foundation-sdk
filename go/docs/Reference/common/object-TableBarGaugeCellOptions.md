---
title: <span class="badge object-type-struct"></span> TableBarGaugeCellOptions
---
# <span class="badge object-type-struct"></span> TableBarGaugeCellOptions

Gauge cell options

## Definition

```go
type TableBarGaugeCellOptions struct {
    Type string `json:"type"`
    Mode *common.BarGaugeDisplayMode `json:"mode,omitempty"`
    ValueDisplayMode *common.BarGaugeValueMode `json:"valueDisplayMode,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TableBarGaugeCellOptions` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (tableBarGaugeCellOptions *TableBarGaugeCellOptions) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `TableBarGaugeCellOptions` objects.

```go
func (tableBarGaugeCellOptions *TableBarGaugeCellOptions) Equals(other TableBarGaugeCellOptions) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `TableBarGaugeCellOptions` fields for violations and returns them.

```go
func (tableBarGaugeCellOptions *TableBarGaugeCellOptions) Validate() error
```

## See also

 * <span class="badge builder"></span> [TableBarGaugeCellOptionsBuilder](./builder-TableBarGaugeCellOptionsBuilder.md)
