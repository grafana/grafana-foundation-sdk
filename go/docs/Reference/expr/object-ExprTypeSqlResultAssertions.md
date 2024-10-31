---
title: <span class="badge object-type-struct"></span> ExprTypeSqlResultAssertions
---
# <span class="badge object-type-struct"></span> ExprTypeSqlResultAssertions

## Definition

```go
type ExprTypeSqlResultAssertions struct {
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
    Type *expr.TypeSqlType `json:"type,omitempty"`
    // TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane
    // contract documentation https://grafana.github.io/dataplane/contract/.
    TypeVersion []int64 `json:"typeVersion"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ExprTypeSqlResultAssertions` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (exprTypeSqlResultAssertions *ExprTypeSqlResultAssertions) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ExprTypeSqlResultAssertions` objects.

```go
func (exprTypeSqlResultAssertions *ExprTypeSqlResultAssertions) Equals(other ExprTypeSqlResultAssertions) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ExprTypeSqlResultAssertions` fields for violations and returns them.

```go
func (exprTypeSqlResultAssertions *ExprTypeSqlResultAssertions) Validate() error
```

## See also

 * <span class="badge builder"></span> [ExprTypeSqlResultAssertionsBuilder](./builder-ExprTypeSqlResultAssertionsBuilder.md)
