---
title: <span class="badge object-type-struct"></span> Percentiles
---
# <span class="badge object-type-struct"></span> Percentiles

## Definition

```go
type Percentiles struct {
    Type string `json:"type"`
    Field *string `json:"field,omitempty"`
    Id string `json:"id"`
    Settings *elasticsearch.ElasticsearchPercentilesSettings `json:"settings,omitempty"`
    Hide *bool `json:"hide,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Percentiles` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (percentiles *Percentiles) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `Percentiles` objects.

```go
func (percentiles *Percentiles) Equals(other Percentiles) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `Percentiles` fields for violations and returns them.

```go
func (percentiles *Percentiles) Validate() error
```

## See also

 * <span class="badge builder"></span> [PercentilesBuilder](./builder-PercentilesBuilder.md)
