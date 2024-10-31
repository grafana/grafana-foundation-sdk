---
title: <span class="badge object-type-struct"></span> HideSeriesConfig
---
# <span class="badge object-type-struct"></span> HideSeriesConfig

TODO docs

## Definition

```go
type HideSeriesConfig struct {
    Tooltip bool `json:"tooltip"`
    Legend bool `json:"legend"`
    Viz bool `json:"viz"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `HideSeriesConfig` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (hideSeriesConfig *HideSeriesConfig) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `HideSeriesConfig` objects.

```go
func (hideSeriesConfig *HideSeriesConfig) Equals(other HideSeriesConfig) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `HideSeriesConfig` fields for violations and returns them.

```go
func (hideSeriesConfig *HideSeriesConfig) Validate() error
```

## See also

 * <span class="badge builder"></span> [HideSeriesConfigBuilder](./builder-HideSeriesConfigBuilder.md)
