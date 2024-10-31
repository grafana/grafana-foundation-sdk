---
title: <span class="badge builder"></span> DateHistogramSettingsBuilder
---
# <span class="badge builder"></span> DateHistogramSettingsBuilder

## Constructor

```go
func NewDateHistogramSettingsBuilder() *DateHistogramSettingsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *DateHistogramSettingsBuilder) Build() (DateHistogramSettings, error)
```

### <span class="badge object-method"></span> Interval

```go
func (builder *DateHistogramSettingsBuilder) Interval(interval string) *DateHistogramSettingsBuilder
```

### <span class="badge object-method"></span> MinDocCount

```go
func (builder *DateHistogramSettingsBuilder) MinDocCount(minDocCount string) *DateHistogramSettingsBuilder
```

### <span class="badge object-method"></span> Offset

```go
func (builder *DateHistogramSettingsBuilder) Offset(offset string) *DateHistogramSettingsBuilder
```

### <span class="badge object-method"></span> TimeZone

```go
func (builder *DateHistogramSettingsBuilder) TimeZone(timeZone string) *DateHistogramSettingsBuilder
```

### <span class="badge object-method"></span> TrimEdges

```go
func (builder *DateHistogramSettingsBuilder) TrimEdges(trimEdges string) *DateHistogramSettingsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [DateHistogramSettings](./object-DateHistogramSettings.md)
