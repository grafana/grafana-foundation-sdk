---
title: <span class="badge object-type-struct"></span> LineStyle
---
# <span class="badge object-type-struct"></span> LineStyle

TODO docs

## Definition

```go
type LineStyle struct {
    Fill *common.LineStyleFill `json:"fill,omitempty"`
    Dash []float64 `json:"dash,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `LineStyle` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (lineStyle *LineStyle) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `LineStyle` objects.

```go
func (lineStyle *LineStyle) Equals(other LineStyle) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `LineStyle` fields for violations and returns them.

```go
func (lineStyle *LineStyle) Validate() error
```

## See also

 * <span class="badge builder"></span> [LineStyleBuilder](./builder-LineStyleBuilder.md)
