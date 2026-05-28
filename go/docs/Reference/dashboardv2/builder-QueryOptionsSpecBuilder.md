---
title: <span class="badge builder"></span> QueryOptionsSpecBuilder
---
# <span class="badge builder"></span> QueryOptionsSpecBuilder

## Constructor

```go
func NewQueryOptionsSpecBuilder() *QueryOptionsSpecBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *QueryOptionsSpecBuilder) Build() (QueryOptionsSpec, error)
```

### <span class="badge object-method"></span> CacheTimeout

```go
func (builder *QueryOptionsSpecBuilder) CacheTimeout(cacheTimeout string) *QueryOptionsSpecBuilder
```

### <span class="badge object-method"></span> HideTimeOverride

```go
func (builder *QueryOptionsSpecBuilder) HideTimeOverride(hideTimeOverride bool) *QueryOptionsSpecBuilder
```

### <span class="badge object-method"></span> Interval

```go
func (builder *QueryOptionsSpecBuilder) Interval(interval string) *QueryOptionsSpecBuilder
```

### <span class="badge object-method"></span> MaxDataPoints

```go
func (builder *QueryOptionsSpecBuilder) MaxDataPoints(maxDataPoints int64) *QueryOptionsSpecBuilder
```

### <span class="badge object-method"></span> QueryCachingTTL

```go
func (builder *QueryOptionsSpecBuilder) QueryCachingTTL(queryCachingTTL int64) *QueryOptionsSpecBuilder
```

### <span class="badge object-method"></span> TimeCompare

```go
func (builder *QueryOptionsSpecBuilder) TimeCompare(timeCompare string) *QueryOptionsSpecBuilder
```

### <span class="badge object-method"></span> TimeFrom

```go
func (builder *QueryOptionsSpecBuilder) TimeFrom(timeFrom string) *QueryOptionsSpecBuilder
```

### <span class="badge object-method"></span> TimeShift

```go
func (builder *QueryOptionsSpecBuilder) TimeShift(timeShift string) *QueryOptionsSpecBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [QueryOptionsSpec](./object-QueryOptionsSpec.md)
