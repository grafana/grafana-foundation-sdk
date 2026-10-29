---
title: <span class="badge object-type-struct"></span> FacetValue
---
# <span class="badge object-type-struct"></span> FacetValue

#FacetValue is a single value/count pair in a facet breakdown.

## Definition

```go
type FacetValue struct {
    Value string `json:"value"`
    Count int64 `json:"count"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `FacetValue` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (facetValue *FacetValue) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `FacetValue` objects.

```go
func (facetValue *FacetValue) Equals(other FacetValue) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `FacetValue` fields for violations and returns them.

```go
func (facetValue *FacetValue) Validate() error
```

## See also

 * <span class="badge builder"></span> [FacetValueBuilder](./builder-FacetValueBuilder.md)
