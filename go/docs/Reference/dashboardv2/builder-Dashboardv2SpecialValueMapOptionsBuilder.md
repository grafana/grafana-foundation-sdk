---
title: <span class="badge builder"></span> Dashboardv2SpecialValueMapOptionsBuilder
---
# <span class="badge builder"></span> Dashboardv2SpecialValueMapOptionsBuilder

## Constructor

```go
func NewDashboardv2SpecialValueMapOptionsBuilder() *Dashboardv2SpecialValueMapOptionsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *Dashboardv2SpecialValueMapOptionsBuilder) Build() (Dashboardv2SpecialValueMapOptions, error)
```

### <span class="badge object-method"></span> Match

Special value to match against

```go
func (builder *Dashboardv2SpecialValueMapOptionsBuilder) Match(match dashboardv2.SpecialValueMatch) *Dashboardv2SpecialValueMapOptionsBuilder
```

### <span class="badge object-method"></span> Result

Config to apply when the value matches the special value

```go
func (builder *Dashboardv2SpecialValueMapOptionsBuilder) Result(result cog.Builder[dashboardv2.ValueMappingResult]) *Dashboardv2SpecialValueMapOptionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Dashboardv2SpecialValueMapOptions](./object-Dashboardv2SpecialValueMapOptions.md)
