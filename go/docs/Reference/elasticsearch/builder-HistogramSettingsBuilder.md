---
title: <span class="badge builder"></span> HistogramSettingsBuilder
---
# <span class="badge builder"></span> HistogramSettingsBuilder

## Constructor

```go
func NewHistogramSettingsBuilder() *HistogramSettingsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *HistogramSettingsBuilder) Build() (HistogramSettings, error)
```

### <span class="badge object-method"></span> Interval

```go
func (builder *HistogramSettingsBuilder) Interval(interval string) *HistogramSettingsBuilder
```

### <span class="badge object-method"></span> MinDocCount

```go
func (builder *HistogramSettingsBuilder) MinDocCount(minDocCount string) *HistogramSettingsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [HistogramSettings](./object-HistogramSettings.md)
