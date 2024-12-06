---
title: <span class="badge object-type-struct"></span> ExtendedStats
---
# <span class="badge object-type-struct"></span> ExtendedStats

## Definition

```go
type ExtendedStats struct {
    Type string `json:"type"`
    Settings *elasticsearch.ElasticsearchExtendedStatsSettings `json:"settings,omitempty"`
    Field *string `json:"field,omitempty"`
    Id string `json:"id"`
    Meta any `json:"meta,omitempty"`
    Hide *bool `json:"hide,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ExtendedStats` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (extendedStats *ExtendedStats) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ExtendedStats` objects.

```go
func (extendedStats *ExtendedStats) Equals(other ExtendedStats) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ExtendedStats` fields for violations and returns them.

```go
func (extendedStats *ExtendedStats) Validate() error
```

## See also

 * <span class="badge builder"></span> [ExtendedStatsBuilder](./builder-ExtendedStatsBuilder.md)
