---
title: <span class="badge builder"></span> ExprTypeSqlResultAssertionsBuilder
---
# <span class="badge builder"></span> ExprTypeSqlResultAssertionsBuilder

## Constructor

```go
func NewExprTypeSqlResultAssertionsBuilder() *ExprTypeSqlResultAssertionsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ExprTypeSqlResultAssertionsBuilder) Build() (ExprTypeSqlResultAssertions, error)
```

### <span class="badge object-method"></span> MaxFrames

Maximum frame count

```go
func (builder *ExprTypeSqlResultAssertionsBuilder) MaxFrames(maxFrames int64) *ExprTypeSqlResultAssertionsBuilder
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
func (builder *ExprTypeSqlResultAssertionsBuilder) Type(typeArg expr.TypeSqlType) *ExprTypeSqlResultAssertionsBuilder
```

### <span class="badge object-method"></span> TypeVersion

TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane

contract documentation https://grafana.github.io/dataplane/contract/.

```go
func (builder *ExprTypeSqlResultAssertionsBuilder) TypeVersion(typeVersion []int64) *ExprTypeSqlResultAssertionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [ExprTypeSqlResultAssertions](./object-ExprTypeSqlResultAssertions.md)
