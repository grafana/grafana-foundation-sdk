---
title: <span class="badge object-type-struct"></span> BaseBucketAggregation
---
# <span class="badge object-type-struct"></span> BaseBucketAggregation

## Definition

```go
type BaseBucketAggregation struct {
    Id string `json:"id"`
    Type elasticsearch.BucketAggregationType `json:"type"`
    Settings any `json:"settings,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `BaseBucketAggregation` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (baseBucketAggregation *BaseBucketAggregation) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `BaseBucketAggregation` objects.

```go
func (baseBucketAggregation *BaseBucketAggregation) Equals(other BaseBucketAggregation) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `BaseBucketAggregation` fields for violations and returns them.

```go
func (baseBucketAggregation *BaseBucketAggregation) Validate() error
```

## See also

 * <span class="badge builder"></span> [BaseBucketAggregationBuilder](./builder-BaseBucketAggregationBuilder.md)
