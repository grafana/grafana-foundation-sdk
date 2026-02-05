---
title: <span class="badge builder"></span> Dashboardv2beta1SpecialValueMapOptionsBuilder
---
# <span class="badge builder"></span> Dashboardv2beta1SpecialValueMapOptionsBuilder

## Constructor

```go
func NewDashboardv2beta1SpecialValueMapOptionsBuilder() *Dashboardv2beta1SpecialValueMapOptionsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *Dashboardv2beta1SpecialValueMapOptionsBuilder) Build() (Dashboardv2beta1SpecialValueMapOptions, error)
```

### <span class="badge object-method"></span> Match

Special value to match against

```go
func (builder *Dashboardv2beta1SpecialValueMapOptionsBuilder) Match(match dashboardv2beta1.SpecialValueMatch) *Dashboardv2beta1SpecialValueMapOptionsBuilder
```

### <span class="badge object-method"></span> Result

Config to apply when the value matches the special value

```go
func (builder *Dashboardv2beta1SpecialValueMapOptionsBuilder) Result(result cog.Builder[dashboardv2beta1.ValueMappingResult]) *Dashboardv2beta1SpecialValueMapOptionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Dashboardv2beta1SpecialValueMapOptions](./object-Dashboardv2beta1SpecialValueMapOptions.md)
