---
title: <span class="badge object-type-struct"></span> GraphThresholdsStyleConfig
---
# <span class="badge object-type-struct"></span> GraphThresholdsStyleConfig

TODO docs

## Definition

```go
type GraphThresholdsStyleConfig struct {
    Mode common.GraphThresholdsStyleMode `json:"mode"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `GraphThresholdsStyleConfig` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (graphThresholdsStyleConfig *GraphThresholdsStyleConfig) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `GraphThresholdsStyleConfig` objects.

```go
func (graphThresholdsStyleConfig *GraphThresholdsStyleConfig) Equals(other GraphThresholdsStyleConfig) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `GraphThresholdsStyleConfig` fields for violations and returns them.

```go
func (graphThresholdsStyleConfig *GraphThresholdsStyleConfig) Validate() error
```

## See also

 * <span class="badge builder"></span> [GraphThresholdsStyleConfigBuilder](./builder-GraphThresholdsStyleConfigBuilder.md)
