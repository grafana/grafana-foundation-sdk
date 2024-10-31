---
title: <span class="badge object-type-struct"></span> BaseDimensionConfig
---
# <span class="badge object-type-struct"></span> BaseDimensionConfig

## Definition

```go
type BaseDimensionConfig struct {
    // fixed: T -- will be added by each element
    Field *string `json:"field,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `BaseDimensionConfig` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (baseDimensionConfig *BaseDimensionConfig) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `BaseDimensionConfig` objects.

```go
func (baseDimensionConfig *BaseDimensionConfig) Equals(other BaseDimensionConfig) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `BaseDimensionConfig` fields for violations and returns them.

```go
func (baseDimensionConfig *BaseDimensionConfig) Validate() error
```

## See also

 * <span class="badge builder"></span> [BaseDimensionConfigBuilder](./builder-BaseDimensionConfigBuilder.md)
