---
title: <span class="badge object-type-struct"></span> FieldConfig
---
# <span class="badge object-type-struct"></span> FieldConfig

## Definition

```go
type FieldConfig struct {
    ScaleDistribution *common.ScaleDistributionConfig `json:"scaleDistribution,omitempty"`
    HideFrom *common.HideSeriesConfig `json:"hideFrom,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `FieldConfig` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (fieldConfig *FieldConfig) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `FieldConfig` objects.

```go
func (fieldConfig *FieldConfig) Equals(other FieldConfig) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `FieldConfig` fields for violations and returns them.

```go
func (fieldConfig *FieldConfig) Validate() error
```

