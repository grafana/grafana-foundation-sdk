---
title: <span class="badge builder"></span> DashboardRegexMapOptionsBuilder
---
# <span class="badge builder"></span> DashboardRegexMapOptionsBuilder

## Constructor

```go
func NewDashboardRegexMapOptionsBuilder() *DashboardRegexMapOptionsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *DashboardRegexMapOptionsBuilder) Build() (DashboardRegexMapOptions, error)
```

### <span class="badge object-method"></span> Pattern

Regular expression to match against

```go
func (builder *DashboardRegexMapOptionsBuilder) Pattern(pattern string) *DashboardRegexMapOptionsBuilder
```

### <span class="badge object-method"></span> Result

Config to apply when the value matches the regex

```go
func (builder *DashboardRegexMapOptionsBuilder) Result(result cog.Builder[dashboard.ValueMappingResult]) *DashboardRegexMapOptionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [DashboardRegexMapOptions](./object-DashboardRegexMapOptions.md)
