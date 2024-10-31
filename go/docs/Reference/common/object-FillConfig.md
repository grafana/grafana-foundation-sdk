---
title: <span class="badge object-type-struct"></span> FillConfig
---
# <span class="badge object-type-struct"></span> FillConfig

TODO docs

## Definition

```go
type FillConfig struct {
    FillColor *string `json:"fillColor,omitempty"`
    FillOpacity *float64 `json:"fillOpacity,omitempty"`
    FillBelowTo *string `json:"fillBelowTo,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `FillConfig` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (fillConfig *FillConfig) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `FillConfig` objects.

```go
func (fillConfig *FillConfig) Equals(other FillConfig) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `FillConfig` fields for violations and returns them.

```go
func (fillConfig *FillConfig) Validate() error
```

## See also

 * <span class="badge builder"></span> [FillConfigBuilder](./builder-FillConfigBuilder.md)
