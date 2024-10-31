---
title: <span class="badge object-type-struct"></span> PointsConfig
---
# <span class="badge object-type-struct"></span> PointsConfig

TODO docs

## Definition

```go
type PointsConfig struct {
    ShowPoints *common.VisibilityMode `json:"showPoints,omitempty"`
    PointSize *float64 `json:"pointSize,omitempty"`
    PointColor *string `json:"pointColor,omitempty"`
    PointSymbol *string `json:"pointSymbol,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `PointsConfig` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (pointsConfig *PointsConfig) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `PointsConfig` objects.

```go
func (pointsConfig *PointsConfig) Equals(other PointsConfig) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `PointsConfig` fields for violations and returns them.

```go
func (pointsConfig *PointsConfig) Validate() error
```

## See also

 * <span class="badge builder"></span> [PointsConfigBuilder](./builder-PointsConfigBuilder.md)
