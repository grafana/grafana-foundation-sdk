---
title: <span class="badge builder"></span> DashboardSpecialValueMapOptionsBuilder
---
# <span class="badge builder"></span> DashboardSpecialValueMapOptionsBuilder

## Constructor

```go
func NewDashboardSpecialValueMapOptionsBuilder() *DashboardSpecialValueMapOptionsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *DashboardSpecialValueMapOptionsBuilder) Build() (DashboardSpecialValueMapOptions, error)
```

### <span class="badge object-method"></span> Match

Special value to match against

```go
func (builder *DashboardSpecialValueMapOptionsBuilder) Match(match dashboard.SpecialValueMatch) *DashboardSpecialValueMapOptionsBuilder
```

### <span class="badge object-method"></span> Result

Config to apply when the value matches the special value

```go
func (builder *DashboardSpecialValueMapOptionsBuilder) Result(result cog.Builder[dashboard.ValueMappingResult]) *DashboardSpecialValueMapOptionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [DashboardSpecialValueMapOptions](./object-DashboardSpecialValueMapOptions.md)
