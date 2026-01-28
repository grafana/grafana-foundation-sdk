---
title: <span class="badge builder"></span> OptionsBuilder
---
# <span class="badge builder"></span> OptionsBuilder

## Constructor

```typescript
new OptionsBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> bucketCount

Bucket count (approx)

```typescript
bucketCount(bucketCount: number)
```

### <span class="badge object-method"></span> bucketOffset

Offset buckets by this amount

```typescript
bucketOffset(bucketOffset: number)
```

### <span class="badge object-method"></span> bucketSize

Size of each bucket

```typescript
bucketSize(bucketSize: number)
```

### <span class="badge object-method"></span> combine

Combines multiple series into a single histogram

```typescript
combine(combine: boolean)
```

### <span class="badge object-method"></span> legend

```typescript
legend(legend: cog.Builder<common.VizLegendOptions>)
```

### <span class="badge object-method"></span> tooltip

```typescript
tooltip(tooltip: cog.Builder<common.VizTooltipOptions>)
```

## See also

 * <span class="badge object-type-interface"></span> [Options](./object-Options.md)
