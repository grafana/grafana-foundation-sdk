---
title: <span class="badge object-type-struct"></span> ReduceDataOptions
---
# <span class="badge object-type-struct"></span> ReduceDataOptions

TODO docs

## Definition

```go
type ReduceDataOptions struct {
    // If true show each row value
    Values *bool `json:"values,omitempty"`
    // if showing all values limit
    Limit *float64 `json:"limit,omitempty"`
    // When !values, pick one value for the whole field
    Calcs []string `json:"calcs"`
    // Which fields to show.  By default this is only numeric fields
    Fields *string `json:"fields,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ReduceDataOptions` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (reduceDataOptions *ReduceDataOptions) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ReduceDataOptions` objects.

```go
func (reduceDataOptions *ReduceDataOptions) Equals(other ReduceDataOptions) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ReduceDataOptions` fields for violations and returns them.

```go
func (reduceDataOptions *ReduceDataOptions) Validate() error
```

## See also

 * <span class="badge builder"></span> [ReduceDataOptionsBuilder](./builder-ReduceDataOptionsBuilder.md)
