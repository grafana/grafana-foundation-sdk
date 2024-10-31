---
title: <span class="badge object-type-struct"></span> RowsHeatmapOptions
---
# <span class="badge object-type-struct"></span> RowsHeatmapOptions

Controls frame rows options

## Definition

```go
type RowsHeatmapOptions struct {
    // Sets the name of the cell when not calculating from data
    Value *string `json:"value,omitempty"`
    // Controls tick alignment when not calculating from data
    Layout *common.HeatmapCellLayout `json:"layout,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `RowsHeatmapOptions` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (rowsHeatmapOptions *RowsHeatmapOptions) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `RowsHeatmapOptions` objects.

```go
func (rowsHeatmapOptions *RowsHeatmapOptions) Equals(other RowsHeatmapOptions) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `RowsHeatmapOptions` fields for violations and returns them.

```go
func (rowsHeatmapOptions *RowsHeatmapOptions) Validate() error
```

## See also

 * <span class="badge builder"></span> [RowsHeatmapOptionsBuilder](./builder-RowsHeatmapOptionsBuilder.md)
