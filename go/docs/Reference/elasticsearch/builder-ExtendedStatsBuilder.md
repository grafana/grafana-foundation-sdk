---
title: <span class="badge builder"></span> ExtendedStatsBuilder
---
# <span class="badge builder"></span> ExtendedStatsBuilder

## Constructor

```go
func NewExtendedStatsBuilder() *ExtendedStatsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ExtendedStatsBuilder) Build() (ExtendedStats, error)
```

### <span class="badge object-method"></span> Field

```go
func (builder *ExtendedStatsBuilder) Field(field string) *ExtendedStatsBuilder
```

### <span class="badge object-method"></span> Hide

```go
func (builder *ExtendedStatsBuilder) Hide(hide bool) *ExtendedStatsBuilder
```

### <span class="badge object-method"></span> Id

```go
func (builder *ExtendedStatsBuilder) Id(id string) *ExtendedStatsBuilder
```

### <span class="badge object-method"></span> Meta

```go
func (builder *ExtendedStatsBuilder) Meta(meta any) *ExtendedStatsBuilder
```

### <span class="badge object-method"></span> Settings

```go
func (builder *ExtendedStatsBuilder) Settings(settings cog.Builder[elasticsearch.ElasticsearchExtendedStatsSettings]) *ExtendedStatsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [ExtendedStats](./object-ExtendedStats.md)
