---
title: <span class="badge object-type-struct"></span> BarConfig
---
# <span class="badge object-type-struct"></span> BarConfig

TODO docs

## Definition

```go
type BarConfig struct {
    BarAlignment *common.BarAlignment `json:"barAlignment,omitempty"`
    BarWidthFactor *float64 `json:"barWidthFactor,omitempty"`
    BarMaxWidth *float64 `json:"barMaxWidth,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `BarConfig` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (barConfig *BarConfig) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `BarConfig` objects.

```go
func (barConfig *BarConfig) Equals(other BarConfig) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `BarConfig` fields for violations and returns them.

```go
func (barConfig *BarConfig) Validate() error
```

## See also

 * <span class="badge builder"></span> [BarConfigBuilder](./builder-BarConfigBuilder.md)
