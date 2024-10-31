---
title: <span class="badge object-type-struct"></span> Rate
---
# <span class="badge object-type-struct"></span> Rate

## Definition

```go
type Rate struct {
    Type string `json:"type"`
    Field *string `json:"field,omitempty"`
    Id string `json:"id"`
    Settings *elasticsearch.ElasticsearchRateSettings `json:"settings,omitempty"`
    Hide *bool `json:"hide,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Rate` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (rate *Rate) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `Rate` objects.

```go
func (rate *Rate) Equals(other Rate) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `Rate` fields for violations and returns them.

```go
func (rate *Rate) Validate() error
```

## See also

 * <span class="badge builder"></span> [RateBuilder](./builder-RateBuilder.md)
