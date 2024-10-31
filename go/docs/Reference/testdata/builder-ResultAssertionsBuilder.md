---
title: <span class="badge builder"></span> ResultAssertionsBuilder
---
# <span class="badge builder"></span> ResultAssertionsBuilder

## Constructor

```go
func NewResultAssertionsBuilder() *ResultAssertionsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ResultAssertionsBuilder) Build() (ResultAssertions, error)
```

### <span class="badge object-method"></span> MaxFrames

Maximum frame count

```go
func (builder *ResultAssertionsBuilder) MaxFrames(maxFrames int64) *ResultAssertionsBuilder
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
func (builder *ResultAssertionsBuilder) Type(typeArg testdata.ResultAssertionsType) *ResultAssertionsBuilder
```

### <span class="badge object-method"></span> TypeVersion

TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane

contract documentation https://grafana.github.io/dataplane/contract/.

```go
func (builder *ResultAssertionsBuilder) TypeVersion(typeVersion []int64) *ResultAssertionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [ResultAssertions](./object-ResultAssertions.md)
