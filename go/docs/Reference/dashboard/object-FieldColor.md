---
title: <span class="badge object-type-struct"></span> FieldColor
---
# <span class="badge object-type-struct"></span> FieldColor

Map a field to a color.

## Definition

```go
type FieldColor struct {
    // The main color scheme mode.
    Mode dashboard.FieldColorModeId `json:"mode"`
    // The fixed color value for fixed or shades color modes.
    FixedColor *string `json:"fixedColor,omitempty"`
    // Some visualizations need to know how to assign a series color from by value color schemes.
    SeriesBy *dashboard.FieldColorSeriesByMode `json:"seriesBy,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `FieldColor` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (fieldColor *FieldColor) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `FieldColor` objects.

```go
func (fieldColor *FieldColor) Equals(other FieldColor) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `FieldColor` fields for violations and returns them.

```go
func (fieldColor *FieldColor) Validate() error
```

## See also

 * <span class="badge builder"></span> [FieldColorBuilder](./builder-FieldColorBuilder.md)
