---
title: <span class="badge builder"></span> RegexMapBuilder
---
# <span class="badge builder"></span> RegexMapBuilder

## Constructor

```go
func NewRegexMapBuilder() *RegexMapBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *RegexMapBuilder) Build() (RegexMap, error)
```

### <span class="badge object-method"></span> Options

Regular expression to match against and the result to apply when the value matches the regex

```go
func (builder *RegexMapBuilder) Options(options cog.Builder[dashboard.DashboardRegexMapOptions]) *RegexMapBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [RegexMap](./object-RegexMap.md)
