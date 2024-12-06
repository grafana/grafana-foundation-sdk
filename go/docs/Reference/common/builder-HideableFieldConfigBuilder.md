---
title: <span class="badge builder"></span> HideableFieldConfigBuilder
---
# <span class="badge builder"></span> HideableFieldConfigBuilder

## Constructor

```go
func NewHideableFieldConfigBuilder() *HideableFieldConfigBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *HideableFieldConfigBuilder) Build() (HideableFieldConfig, error)
```

### <span class="badge object-method"></span> HideFrom

```go
func (builder *HideableFieldConfigBuilder) HideFrom(hideFrom cog.Builder[common.HideSeriesConfig]) *HideableFieldConfigBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [HideableFieldConfig](./object-HideableFieldConfig.md)
