---
title: <span class="badge object-type-struct"></span> UpdateConfig
---
# <span class="badge object-type-struct"></span> UpdateConfig

## Definition

```go
type UpdateConfig struct {
    Render bool `json:"render"`
    DataChanged bool `json:"dataChanged"`
    SchemaChanged bool `json:"schemaChanged"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `UpdateConfig` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (updateConfig *UpdateConfig) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `UpdateConfig` objects.

```go
func (updateConfig *UpdateConfig) Equals(other UpdateConfig) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `UpdateConfig` fields for violations and returns them.

```go
func (updateConfig *UpdateConfig) Validate() error
```

## See also

 * <span class="badge builder"></span> [UpdateConfigBuilder](./builder-UpdateConfigBuilder.md)
