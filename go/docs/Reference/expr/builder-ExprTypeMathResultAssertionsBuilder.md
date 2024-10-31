---
title: <span class="badge builder"></span> ExprTypeMathResultAssertionsBuilder
---
# <span class="badge builder"></span> ExprTypeMathResultAssertionsBuilder

## Constructor

```go
func NewExprTypeMathResultAssertionsBuilder() *ExprTypeMathResultAssertionsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ExprTypeMathResultAssertionsBuilder) Build() (ExprTypeMathResultAssertions, error)
```

### <span class="badge object-method"></span> MaxFrames

Maximum frame count

```go
func (builder *ExprTypeMathResultAssertionsBuilder) MaxFrames(maxFrames int64) *ExprTypeMathResultAssertionsBuilder
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
func (builder *ExprTypeMathResultAssertionsBuilder) Type(typeArg expr.TypeMathType) *ExprTypeMathResultAssertionsBuilder
```

### <span class="badge object-method"></span> TypeVersion

TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane

contract documentation https://grafana.github.io/dataplane/contract/.

```go
func (builder *ExprTypeMathResultAssertionsBuilder) TypeVersion(typeVersion []int64) *ExprTypeMathResultAssertionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [ExprTypeMathResultAssertions](./object-ExprTypeMathResultAssertions.md)
