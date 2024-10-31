---
title: <span class="badge object-type-struct"></span> XYDimensionConfig
---
# <span class="badge object-type-struct"></span> XYDimensionConfig

Configuration for the Table/Auto mode

## Definition

```go
type XYDimensionConfig struct {
    Frame int32 `json:"frame"`
    X *string `json:"x,omitempty"`
    Exclude []string `json:"exclude,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `XYDimensionConfig` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (xYDimensionConfig *XYDimensionConfig) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `XYDimensionConfig` objects.

```go
func (xYDimensionConfig *XYDimensionConfig) Equals(other XYDimensionConfig) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `XYDimensionConfig` fields for violations and returns them.

```go
func (xYDimensionConfig *XYDimensionConfig) Validate() error
```

## See also

 * <span class="badge builder"></span> [XYDimensionConfigBuilder](./builder-XYDimensionConfigBuilder.md)
