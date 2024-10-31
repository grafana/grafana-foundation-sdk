---
title: <span class="badge builder"></span> ExprTypeReduceResultAssertionsBuilder
---
# <span class="badge builder"></span> ExprTypeReduceResultAssertionsBuilder

## Constructor

```go
func NewExprTypeReduceResultAssertionsBuilder() *ExprTypeReduceResultAssertionsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ExprTypeReduceResultAssertionsBuilder) Build() (ExprTypeReduceResultAssertions, error)
```

### <span class="badge object-method"></span> MaxFrames

Maximum frame count

```go
func (builder *ExprTypeReduceResultAssertionsBuilder) MaxFrames(maxFrames int64) *ExprTypeReduceResultAssertionsBuilder
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
func (builder *ExprTypeReduceResultAssertionsBuilder) Type(typeArg expr.TypeReduceType) *ExprTypeReduceResultAssertionsBuilder
```

### <span class="badge object-method"></span> TypeVersion

TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane

contract documentation https://grafana.github.io/dataplane/contract/.

```go
func (builder *ExprTypeReduceResultAssertionsBuilder) TypeVersion(typeVersion []int64) *ExprTypeReduceResultAssertionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [ExprTypeReduceResultAssertions](./object-ExprTypeReduceResultAssertions.md)
