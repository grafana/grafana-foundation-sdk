---
title: <span class="badge builder"></span> OptionsBuilder
---
# <span class="badge builder"></span> OptionsBuilder

## Constructor

```go
func NewOptionsBuilder() *OptionsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *OptionsBuilder) Build() (Options, error)
```

### <span class="badge object-method"></span> BucketOffset

Offset buckets by this amount

```go
func (builder *OptionsBuilder) BucketOffset(bucketOffset int32) *OptionsBuilder
```

### <span class="badge object-method"></span> BucketSize

Size of each bucket

```go
func (builder *OptionsBuilder) BucketSize(bucketSize int32) *OptionsBuilder
```

### <span class="badge object-method"></span> Combine

Combines multiple series into a single histogram

```go
func (builder *OptionsBuilder) Combine(combine bool) *OptionsBuilder
```

### <span class="badge object-method"></span> Legend

```go
func (builder *OptionsBuilder) Legend(legend cog.Builder[common.VizLegendOptions]) *OptionsBuilder
```

### <span class="badge object-method"></span> Tooltip

```go
func (builder *OptionsBuilder) Tooltip(tooltip cog.Builder[common.VizTooltipOptions]) *OptionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Options](./object-Options.md)
