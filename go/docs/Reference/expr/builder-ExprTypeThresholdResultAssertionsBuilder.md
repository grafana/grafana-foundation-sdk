---
title: <span class="badge builder"></span> ExprTypeThresholdResultAssertionsBuilder
---
# <span class="badge builder"></span> ExprTypeThresholdResultAssertionsBuilder

## Constructor

```go
func NewExprTypeThresholdResultAssertionsBuilder() *ExprTypeThresholdResultAssertionsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ExprTypeThresholdResultAssertionsBuilder) Build() (ExprTypeThresholdResultAssertions, error)
```

### <span class="badge object-method"></span> MaxFrames

Maximum frame count

```go
func (builder *ExprTypeThresholdResultAssertionsBuilder) MaxFrames(maxFrames int64) *ExprTypeThresholdResultAssertionsBuilder
```

### <span class="badge object-method"></span> Type

Type asserts that the frame matches a known type structure.

Possible enum values:

 - `""` 

 - `"timeseries-wide"` 

 - `"timeseries-long"` 

 - `"timeseries-many"` 

 - `"timeseries-multi"` 

 - `"directory-listing"` 

 - `"table"` 

 - `"numeric-wide"` 

 - `"numeric-multi"` 

 - `"numeric-long"` 

 - `"log-lines"` 

```go
func (builder *ExprTypeThresholdResultAssertionsBuilder) Type(typeArg expr.TypeThresholdType) *ExprTypeThresholdResultAssertionsBuilder
```

### <span class="badge object-method"></span> TypeVersion

TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane

contract documentation https://grafana.github.io/dataplane/contract/.

```go
func (builder *ExprTypeThresholdResultAssertionsBuilder) TypeVersion(typeVersion []int64) *ExprTypeThresholdResultAssertionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [ExprTypeThresholdResultAssertions](./object-ExprTypeThresholdResultAssertions.md)
