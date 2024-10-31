---
title: <span class="badge object-type-struct"></span> ResourceDimensionConfig
---
# <span class="badge object-type-struct"></span> ResourceDimensionConfig

Links to a resource (image/svg path)

## Definition

```go
type ResourceDimensionConfig struct {
    Mode common.ResourceDimensionMode `json:"mode"`
    // fixed: T -- will be added by each element
    Field *string `json:"field,omitempty"`
    Fixed *string `json:"fixed,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ResourceDimensionConfig` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (resourceDimensionConfig *ResourceDimensionConfig) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ResourceDimensionConfig` objects.

```go
func (resourceDimensionConfig *ResourceDimensionConfig) Equals(other ResourceDimensionConfig) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ResourceDimensionConfig` fields for violations and returns them.

```go
func (resourceDimensionConfig *ResourceDimensionConfig) Validate() error
```

## See also

 * <span class="badge builder"></span> [ResourceDimensionConfigBuilder](./builder-ResourceDimensionConfigBuilder.md)
