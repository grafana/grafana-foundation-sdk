---
title: <span class="badge object-type-struct"></span> LineConfig
---
# <span class="badge object-type-struct"></span> LineConfig

TODO docs

## Definition

```go
type LineConfig struct {
    LineColor *string `json:"lineColor,omitempty"`
    LineWidth *float64 `json:"lineWidth,omitempty"`
    LineInterpolation *common.LineInterpolation `json:"lineInterpolation,omitempty"`
    LineStyle *common.LineStyle `json:"lineStyle,omitempty"`
    // Indicate if null values should be treated as gaps or connected.
    // When the value is a number, it represents the maximum delta in the
    // X axis that should be considered connected.  For timeseries, this is milliseconds
    SpanNulls *common.BoolOrFloat64 `json:"spanNulls,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `LineConfig` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (lineConfig *LineConfig) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `LineConfig` objects.

```go
func (lineConfig *LineConfig) Equals(other LineConfig) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `LineConfig` fields for violations and returns them.

```go
func (lineConfig *LineConfig) Validate() error
```

## See also

 * <span class="badge builder"></span> [LineConfigBuilder](./builder-LineConfigBuilder.md)
