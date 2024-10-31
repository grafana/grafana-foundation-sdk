---
title: <span class="badge builder"></span> ExprTypeClassicConditionsResultAssertionsBuilder
---
# <span class="badge builder"></span> ExprTypeClassicConditionsResultAssertionsBuilder

## Constructor

```go
func NewExprTypeClassicConditionsResultAssertionsBuilder() *ExprTypeClassicConditionsResultAssertionsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ExprTypeClassicConditionsResultAssertionsBuilder) Build() (ExprTypeClassicConditionsResultAssertions, error)
```

### <span class="badge object-method"></span> MaxFrames

Maximum frame count

```go
func (builder *ExprTypeClassicConditionsResultAssertionsBuilder) MaxFrames(maxFrames int64) *ExprTypeClassicConditionsResultAssertionsBuilder
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
func (builder *ExprTypeClassicConditionsResultAssertionsBuilder) Type(typeArg expr.TypeClassicConditionsType) *ExprTypeClassicConditionsResultAssertionsBuilder
```

### <span class="badge object-method"></span> TypeVersion

TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane

contract documentation https://grafana.github.io/dataplane/contract/.

```go
func (builder *ExprTypeClassicConditionsResultAssertionsBuilder) TypeVersion(typeVersion []int64) *ExprTypeClassicConditionsResultAssertionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [ExprTypeClassicConditionsResultAssertions](./object-ExprTypeClassicConditionsResultAssertions.md)
