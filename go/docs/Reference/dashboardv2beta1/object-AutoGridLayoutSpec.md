---
title: <span class="badge object-type-struct"></span> AutoGridLayoutSpec
---
# <span class="badge object-type-struct"></span> AutoGridLayoutSpec

## Definition

```go
type AutoGridLayoutSpec struct {
    MaxColumnCount *float64 `json:"maxColumnCount,omitempty"`
    ColumnWidthMode dashboardv2beta1.AutoGridLayoutSpecColumnWidthMode `json:"columnWidthMode"`
    ColumnWidth *float64 `json:"columnWidth,omitempty"`
    RowHeightMode dashboardv2beta1.AutoGridLayoutSpecRowHeightMode `json:"rowHeightMode"`
    RowHeight *float64 `json:"rowHeight,omitempty"`
    FillScreen *bool `json:"fillScreen,omitempty"`
    Items []dashboardv2beta1.AutoGridLayoutItemKind `json:"items"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AutoGridLayoutSpec` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (autoGridLayoutSpec *AutoGridLayoutSpec) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `AutoGridLayoutSpec` objects.

```go
func (autoGridLayoutSpec *AutoGridLayoutSpec) Equals(other AutoGridLayoutSpec) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `AutoGridLayoutSpec` fields for violations and returns them.

```go
func (autoGridLayoutSpec *AutoGridLayoutSpec) Validate() error
```

