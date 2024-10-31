---
title: <span class="badge object-type-struct"></span> HideableFieldConfig
---
# <span class="badge object-type-struct"></span> HideableFieldConfig

TODO docs

## Definition

```go
type HideableFieldConfig struct {
    HideFrom *common.HideSeriesConfig `json:"hideFrom,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `HideableFieldConfig` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (hideableFieldConfig *HideableFieldConfig) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `HideableFieldConfig` objects.

```go
func (hideableFieldConfig *HideableFieldConfig) Equals(other HideableFieldConfig) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `HideableFieldConfig` fields for violations and returns them.

```go
func (hideableFieldConfig *HideableFieldConfig) Validate() error
```

## See also

 * <span class="badge builder"></span> [HideableFieldConfigBuilder](./builder-HideableFieldConfigBuilder.md)
