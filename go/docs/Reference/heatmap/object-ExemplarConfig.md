---
title: <span class="badge object-type-struct"></span> ExemplarConfig
---
# <span class="badge object-type-struct"></span> ExemplarConfig

Controls exemplar options

## Definition

```go
type ExemplarConfig struct {
    // Sets the color of the exemplar markers
    Color string `json:"color"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ExemplarConfig` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (exemplarConfig *ExemplarConfig) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ExemplarConfig` objects.

```go
func (exemplarConfig *ExemplarConfig) Equals(other ExemplarConfig) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ExemplarConfig` fields for violations and returns them.

```go
func (exemplarConfig *ExemplarConfig) Validate() error
```

## See also

 * <span class="badge builder"></span> [ExemplarConfigBuilder](./builder-ExemplarConfigBuilder.md)
