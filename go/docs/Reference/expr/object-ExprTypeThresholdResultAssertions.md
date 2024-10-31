---
title: <span class="badge object-type-struct"></span> ExprTypeThresholdResultAssertions
---
# <span class="badge object-type-struct"></span> ExprTypeThresholdResultAssertions

## Definition

```go
type ExprTypeThresholdResultAssertions struct {
    // Maximum frame count
    MaxFrames *int64 `json:"maxFrames,omitempty"`
    // Type asserts that the frame matches a known type structure.
    // Possible enum values:
    //  - `""` 
    //  - `"timeseries-wide"` 
    //  - `"timeseries-long"` 
    //  - `"timeseries-many"` 
    //  - `"timeseries-multi"` 
    //  - `"directory-listing"` 
    //  - `"table"` 
    //  - `"numeric-wide"` 
    //  - `"numeric-multi"` 
    //  - `"numeric-long"` 
    //  - `"log-lines"` 
    Type *expr.TypeThresholdType `json:"type,omitempty"`
    // TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane
    // contract documentation https://grafana.github.io/dataplane/contract/.
    TypeVersion []int64 `json:"typeVersion"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ExprTypeThresholdResultAssertions` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (exprTypeThresholdResultAssertions *ExprTypeThresholdResultAssertions) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ExprTypeThresholdResultAssertions` objects.

```go
func (exprTypeThresholdResultAssertions *ExprTypeThresholdResultAssertions) Equals(other ExprTypeThresholdResultAssertions) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ExprTypeThresholdResultAssertions` fields for violations and returns them.

```go
func (exprTypeThresholdResultAssertions *ExprTypeThresholdResultAssertions) Validate() error
```

## See also

 * <span class="badge builder"></span> [ExprTypeThresholdResultAssertionsBuilder](./builder-ExprTypeThresholdResultAssertionsBuilder.md)
