---
title: <span class="badge builder"></span> Dashboardv2RegexMapOptionsBuilder
---
# <span class="badge builder"></span> Dashboardv2RegexMapOptionsBuilder

## Constructor

```go
func NewDashboardv2RegexMapOptionsBuilder() *Dashboardv2RegexMapOptionsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *Dashboardv2RegexMapOptionsBuilder) Build() (Dashboardv2RegexMapOptions, error)
```

### <span class="badge object-method"></span> Pattern

Regular expression to match against

```go
func (builder *Dashboardv2RegexMapOptionsBuilder) Pattern(pattern string) *Dashboardv2RegexMapOptionsBuilder
```

### <span class="badge object-method"></span> Result

Config to apply when the value matches the regex

```go
func (builder *Dashboardv2RegexMapOptionsBuilder) Result(result cog.Builder[dashboardv2.ValueMappingResult]) *Dashboardv2RegexMapOptionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Dashboardv2RegexMapOptions](./object-Dashboardv2RegexMapOptions.md)
