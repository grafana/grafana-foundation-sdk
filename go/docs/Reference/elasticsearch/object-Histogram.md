---
title: <span class="badge object-type-struct"></span> Histogram
---
# <span class="badge object-type-struct"></span> Histogram

## Definition

```go
type Histogram struct {
    Field *string `json:"field,omitempty"`
    Id string `json:"id"`
    Type string `json:"type"`
    Settings *elasticsearch.ElasticsearchHistogramSettings `json:"settings,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Histogram` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (histogram *Histogram) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `Histogram` objects.

```go
func (histogram *Histogram) Equals(other Histogram) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `Histogram` fields for violations and returns them.

```go
func (histogram *Histogram) Validate() error
```

## See also

 * <span class="badge builder"></span> [HistogramBuilder](./builder-HistogramBuilder.md)
