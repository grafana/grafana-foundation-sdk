---
title: <span class="badge builder"></span> Dashboardv2beta1RegexMapOptionsBuilder
---
# <span class="badge builder"></span> Dashboardv2beta1RegexMapOptionsBuilder

## Constructor

```go
func NewDashboardv2beta1RegexMapOptionsBuilder() *Dashboardv2beta1RegexMapOptionsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *Dashboardv2beta1RegexMapOptionsBuilder) Build() (Dashboardv2beta1RegexMapOptions, error)
```

### <span class="badge object-method"></span> Pattern

Regular expression to match against

```go
func (builder *Dashboardv2beta1RegexMapOptionsBuilder) Pattern(pattern string) *Dashboardv2beta1RegexMapOptionsBuilder
```

### <span class="badge object-method"></span> Result

Config to apply when the value matches the regex

```go
func (builder *Dashboardv2beta1RegexMapOptionsBuilder) Result(result cog.Builder[dashboardv2beta1.ValueMappingResult]) *Dashboardv2beta1RegexMapOptionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Dashboardv2beta1RegexMapOptions](./object-Dashboardv2beta1RegexMapOptions.md)
