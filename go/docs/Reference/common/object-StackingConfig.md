---
title: <span class="badge object-type-struct"></span> StackingConfig
---
# <span class="badge object-type-struct"></span> StackingConfig

TODO docs

## Definition

```go
type StackingConfig struct {
    Mode *common.StackingMode `json:"mode,omitempty"`
    Group *string `json:"group,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `StackingConfig` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (stackingConfig *StackingConfig) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `StackingConfig` objects.

```go
func (stackingConfig *StackingConfig) Equals(other StackingConfig) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `StackingConfig` fields for violations and returns them.

```go
func (stackingConfig *StackingConfig) Validate() error
```

## See also

 * <span class="badge builder"></span> [StackingConfigBuilder](./builder-StackingConfigBuilder.md)
