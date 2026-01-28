---
title: <span class="badge builder"></span> FieldConfigBuilder
---
# <span class="badge builder"></span> FieldConfigBuilder

## Constructor

```go
func NewFieldConfigBuilder() *FieldConfigBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *FieldConfigBuilder) Build() (FieldConfig, error)
```

### <span class="badge object-method"></span> FillOpacity

```go
func (builder *FieldConfigBuilder) FillOpacity(fillOpacity uint32) *FieldConfigBuilder
```

### <span class="badge object-method"></span> HideFrom

```go
func (builder *FieldConfigBuilder) HideFrom(hideFrom cog.Builder[common.HideSeriesConfig]) *FieldConfigBuilder
```

### <span class="badge object-method"></span> LineWidth

```go
func (builder *FieldConfigBuilder) LineWidth(lineWidth uint32) *FieldConfigBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [FieldConfig](./object-FieldConfig.md)
