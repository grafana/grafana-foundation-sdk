---
title: <span class="badge builder"></span> FiltersSettingsBuilder
---
# <span class="badge builder"></span> FiltersSettingsBuilder

## Constructor

```go
func NewFiltersSettingsBuilder() *FiltersSettingsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *FiltersSettingsBuilder) Build() (FiltersSettings, error)
```

### <span class="badge object-method"></span> Filters

```go
func (builder *FiltersSettingsBuilder) Filters(filters []cog.Builder[elasticsearch.Filter]) *FiltersSettingsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [FiltersSettings](./object-FiltersSettings.md)
