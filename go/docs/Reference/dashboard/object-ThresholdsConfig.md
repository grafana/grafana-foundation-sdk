---
title: <span class="badge object-type-struct"></span> ThresholdsConfig
---
# <span class="badge object-type-struct"></span> ThresholdsConfig

Thresholds configuration for the panel

## Definition

```go
type ThresholdsConfig struct {
    // Thresholds mode.
    Mode dashboard.ThresholdsMode `json:"mode"`
    // Must be sorted by 'value', first value is always -Infinity
    Steps []dashboard.Threshold `json:"steps"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ThresholdsConfig` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (thresholdsConfig *ThresholdsConfig) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ThresholdsConfig` objects.

```go
func (thresholdsConfig *ThresholdsConfig) Equals(other ThresholdsConfig) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ThresholdsConfig` fields for violations and returns them.

```go
func (thresholdsConfig *ThresholdsConfig) Validate() error
```

## See also

 * <span class="badge builder"></span> [ThresholdsConfigBuilder](./builder-ThresholdsConfigBuilder.md)
